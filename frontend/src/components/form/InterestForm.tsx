import { useAuthGuard } from "@/lib/hook/useAuthGuard";
import { InterestSchema } from "@/lib/schema/InterestSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
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
import { cn } from "@/lib/utils";
import { Tag } from "@/interface/Tag";

export default function InterestForm() {
  const user = useAuthGuard();
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

  async function fetchCurrentTags() {
    try {
      const res = await axios.get(`${env.VITE_BACKEND_URL}/api/user/tag`, {
        withCredentials: true,
      });

      if (res.status === 200) {
        const { data } = res.data;
        form.reset({
          interests: data.map((tag: Tag) => tag.id),
        });
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
    fetchCurrentTags();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [user]);

  const onSubmit = async (data: z.infer<typeof InterestSchema>) => {
    const formData = new FormData();
    if (user?.id) {
      formData.append("userId", user.id);
    } else {
      console.error("User ID is missing.");
      return;
    }

    formData.append("preferences", JSON.stringify(data.interests));

    try {
      const res = await axios.put(
        `${env.VITE_BACKEND_URL}/api/user/preference`,
        formData,
        {
          withCredentials: true,
        }
      );

      if (res.status === 200) {
        toast.success("Interests updated successfully!");
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
              <div className="mb-16 flex flex-col gap-8">
                <FormLabel className="text-5xl font-semibold">
                  Share your interests with us
                </FormLabel>
                <FormDescription className="text-2xl">
                  Choose your interests below to get personalized event
                  suggestions.
                </FormDescription>
              </div>
              <div className="flex flex-col gap-4">
                {tagTypes.map((tagType) => (
                  <div key={tagType.id} className="flex flex-col gap-4">
                    <span className="text-xl font-bold">{tagType.name}</span>
                    <div className="flex flex-row gap-4">
                      {tagType.tags.map((tag) => {
                        const isChecked = field.value?.includes(tag.id);
                        return (
                          <FormItem
                            key={tag.id}
                            className="flex flex-row items-start space-x-3 space-y-0"
                          >
                            <FormControl>
                              <label
                                className={cn(
                                  "inline-flex items-center px-4 py-2 rounded-full border text-sm font-medium cursor-pointer transition-all ",
                                  isChecked
                                    ? "bg-yellow-300 text-black border-black"
                                    : "bg-white text-gray-700 border-gray-300 hover:bg-gray-100"
                                )}
                              >
                                <Checkbox
                                  checked={isChecked}
                                  onCheckedChange={(checked) => {
                                    const updated = checked
                                      ? [...field.value, tag.id]
                                      : field.value.filter(
                                          (val) => val !== tag.id
                                        );
                                    field.onChange(updated);
                                  }}
                                  className="hidden"
                                />
                                {tag.name}
                                <span
                                  className={cn(
                                    "transition-all duration-300 ease-in-out transform",
                                    isChecked
                                      ? "pl-8 opacity-100 translate-x-0 max-w-fit"
                                      : "opacity-0 -translate-x-2 pointer-events-none max-w-0"
                                  )}
                                >
                                  X
                                </span>
                              </label>
                            </FormControl>
                          </FormItem>
                        );
                      })}
                    </div>
                    <div className="flex items-center justify-center my-6">
                      <div className="border-t border-gray-300 flex-grow"></div>
                    </div>
                  </div>
                ))}
              </div>
              <FormMessage />
            </FormItem>
          )}
        />

        <div className="flex w-full justify-end">
          <Button className="px-8 py-6 text-xl rounded-xl" type="submit">
            Save my interests
          </Button>
        </div>
      </form>
    </Form>
  );
}
