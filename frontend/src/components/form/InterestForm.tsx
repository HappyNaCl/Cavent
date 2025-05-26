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
import { CategoryType } from "@/interface/CategoryType";
import { Checkbox } from "../ui/checkbox";
import { Button } from "../ui/button";
import { cn } from "@/lib/utils";
import api from "@/lib/axios";
import { Category } from "@/interface/Category";
import axios from "axios";
import { useNavigate } from "react-router";

export default function InterestForm() {
  const [categoryTypes, setCategoryTypes] = useState<CategoryType[]>([]);
  const nav = useNavigate();

  const form = useForm({
    resolver: zodResolver(InterestSchema),
    defaultValues: {
      interests: [],
    },
  });

  useEffect(() => {
    async function fetchTagTypes() {
      try {
        const res = await api.get("/category");
        console.log("Tag types fetched:", res.data.data);
        if (res.status === 200 && res.data.data) {
          setCategoryTypes(res.data.data);
        }
      } catch (error) {
        toast.error(`Error: ${error}`);
      }
    }

    async function fetchCurrentTags() {
      try {
        const res = await api.get("/user/interest");
        console.log("Current tags fetched:", res.data.data);
        if (res.status === 200 && res.data.data) {
          const { data } = res.data;
          form.reset({
            interests: data.map((category: Category) => category.id),
          });
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(`Error: ${error.response?.data.error || error.message}`);
        } else {
          toast.error(`Error: ${error}`);
        }
      }
    }

    fetchTagTypes();
    fetchCurrentTags();
  }, [form]);

  const onSubmit = async (data: z.infer<typeof InterestSchema>) => {
    const formData = new FormData();
    data.interests.forEach((interest) => {
      formData.append("categoryIds", interest);
    });
    console.log(data.interests);
    try {
      const res = await api.put("/user/interest", formData);
      if (res.status === 200) {
        toast.success("Interests updated successfully!");
        nav("/");
      } else {
        toast.error(`Error: ${res.data.error}`);
      }
      console.log("Submitting interests:", data.interests);
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
                {categoryTypes.map((categoryType) => (
                  <div key={categoryType.id} className="flex flex-col gap-4">
                    <span className="text-xl font-bold">
                      {categoryType.name}
                    </span>
                    <div className="flex flex-row gap-4">
                      {categoryType.categories.map((category) => {
                        const isChecked = field.value?.includes(category.id);
                        return (
                          <FormItem
                            key={category.id}
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
                                      ? [...field.value, category.id]
                                      : field.value.filter(
                                          (val) => val !== category.id
                                        );
                                    field.onChange(updated);
                                  }}
                                  className="hidden"
                                />
                                {category.name}
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
