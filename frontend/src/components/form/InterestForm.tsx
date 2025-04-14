import { useAuthGuard } from "@/lib/hook/useAuthGuard";
import { InterestSchema } from "@/lib/schema/InterestSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router";
import { toast } from "sonner";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { useEffect, useState } from "react";
import { TagType } from "@/interface/TagType";
import { Checkbox } from "../ui/checkbox";
import { Button } from "../ui/button";
import axios from "axios";
import { env } from "@/lib/schema/EnvSchema";

export default function InterestForm() {
  const user = useAuthGuard();
  const nav = useNavigate();
  const [tagTypes, setTagTypes] = useState<TagType[]>([]);

  async function fetchTagTypes() {
    try {
      const res = await axios.get(`${env.VITE_BACKEND_URL}/api/tags`, {
        withCredentials: true,
      });
      if (res.status === 200) {
        const { data } = res.data;
        setTagTypes(data);
      }
    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  }

  const form = useForm({
    resolver: zodResolver(InterestSchema),
    defaultValues: {
      interests: [],
    },
  });

  useEffect(() => {
    fetchTagTypes();
  }, [user]);

  const onSubmit = async (data: z.infer<typeof InterestSchema>) => {
    const formData = new FormData();
    if (user?.id) {
      formData.append("userId", user.id);
    } else {
      toast.error("User ID is missing.");
      return;
    }
    formData.append("preferences", JSON.stringify(data.interests));
    console.log(formData);
    try {
      const res = await axios.put(
        `${env.VITE_BACKEND_URL}/api/user/preference`,
        formData,
        {
          withCredentials: true,
        }
      );

      if (res.status === 200) {
        alert("Preference updated successfully");
        nav("/");
      }
    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="interests"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Sidebar</FormLabel>
              <FormDescription>
                Select the items you want to display in the sidebar.
              </FormDescription>

              {tagTypes.map((tagType) => (
                <div key={tagType.id} className="flex flex-col gap-4">
                  <span className="text-xl font-bold">{tagType.name}</span>
                  {tagType.tags.map((tag) => {
                    const isChecked = field.value?.includes(tag.id);
                    return (
                      <FormItem
                        key={tag.id}
                        className="flex flex-row items-start space-x-3 space-y-0"
                      >
                        <FormControl>
                          <Checkbox
                            checked={isChecked}
                            onCheckedChange={(checked) => {
                              const updated = checked
                                ? [...field.value, tag.id]
                                : field.value.filter((val) => val !== tag.id);
                              field.onChange(updated);
                            }}
                          />
                        </FormControl>
                        <FormLabel className="text-sm font-normal">
                          {tag.name}
                        </FormLabel>
                      </FormItem>
                    );
                  })}
                </div>
              ))}

              <FormMessage />
            </FormItem>
          )}
        />

        <Button type="submit">Submit</Button>
      </form>
    </Form>
  );
}
