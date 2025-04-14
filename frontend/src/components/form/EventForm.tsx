"use client";

import { useState } from "react";
import { useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
} from "@/components/ui/form";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { toast } from "sonner";
import { cn } from "@/lib/utils";
import { Input } from "@/components/ui/input";
import { z } from "zod";
import { EventSchema } from "@/lib/schema/EventSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import TextInput from "../input/TextInput";

export const EventForm = () => {
  const [step, setStep] = useState(0);
  const totalSteps = 4;

  const crumbs = ["Edit", "Banner", "Ticketing", "Review"];
  const titles = [
    { title: "Event Details", description: "Enter your event details" },
    { title: "Event Banner", description: "Upload your event banner" },
    { title: "Ticketing", description: "Set up your ticketing options" },
    { title: "Review", description: "Review your event details" },
  ];

  const form = useForm<z.infer<typeof EventSchema>>({
    // resolver: zodResolver(EventSchema),
    defaultValues: {
      title: "",
      category: "",
      eventType: "single",
      startDate: new Date(),
      startTime: "",
      location: "",
      description: "",
    },
  });

  const onSubmit = async (formData: unknown) => {
    if (step < totalSteps - 1) {
      setStep(step + 1);
    } else {
      console.log(formData);
      setStep(0);
      form.reset();

      toast.success("Form successfully submitted");
    }
  };

  const handleBack = () => {
    if (step > 0) {
      setStep(step - 1);
    }
  };

  return (
    <div className="space-y-12 w-full">
      <div className="flex items-center justify-center">
        {Array.from({ length: totalSteps }).map((_, index) => (
          <div key={index} className="flex items-center">
            <div className="flex flex-col items-center justify-center relative">
              <div
                className={cn(
                  "w-4 h-4 rounded-full transition-all duration-300 ease-in-out",
                  index <= step ? "bg-primary" : "bg-primary/30",
                  index < step && "bg-primary"
                )}
              />
              <span className="absolute top-5">{crumbs[index]}</span>
            </div>
            {index < totalSteps - 1 && (
              <div
                className={cn(
                  "w-64 h-0.5",
                  index < step ? "bg-primary" : "bg-primary/30"
                )}
              />
            )}
          </div>
        ))}
      </div>
      <Card className="shadow-sm">
        <CardHeader>
          <CardTitle className="text-lg">{titles[step].title}</CardTitle>
          <CardDescription>{titles[step].description}</CardDescription>
        </CardHeader>
        <CardContent>
          {step === 0 && (
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="grid gap-y-4"
              >
                <FormField
                  key="title"
                  control={form.control}
                  name="title"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Event Title</FormLabel>
                      <FormControl>
                        <TextInput
                          type="text"
                          id="title"
                          placeholder="Title"
                          onChange={field.onChange}
                          value={field.value}
                          className="border border-gray-300 rounded-lg px-4 py-6"
                        />
                      </FormControl>
                      <FormDescription>
                        Enter your event's title
                      </FormDescription>
                    </FormItem>
                  )}
                />

                <FormField
                  key="category"
                  control={form.control}
                  name="category"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Event Category</FormLabel>
                      <FormControl>
                        <TextInput
                          type="text"
                          id="category"
                          placeholder="category"
                          onChange={field.onChange}
                          value={field.value}
                          className="border border-gray-300 rounded-lg px-4 py-6"
                        />
                      </FormControl>
                      <FormDescription></FormDescription>
                    </FormItem>
                  )}
                />

                <div className="flex justify-between">
                  <Button
                    type="button"
                    className="font-medium"
                    size="sm"
                    onClick={handleBack}
                    disabled={step === 0}
                  >
                    Back
                  </Button>
                  <Button type="submit" size="sm" className="font-medium">
                    {step === totalSteps - 1 ? "Submit" : "Next"}
                  </Button>
                </div>
              </form>
            </Form>
          )}

          {step === 1 && (
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="grid gap-y-4"
              >
                <div className="border border-dashed rounded-md">
                  <div className="flex flex-col items-center justify-center h-[8rem]">
                    <h3 className="text-base font-semibold text-center">
                      No Inputs Added Yet!
                    </h3>
                    <p className="text-xs text-muted-foreground text-center">
                      Start building your form by adding input fields.
                    </p>
                  </div>
                </div>

                <div className="flex justify-between">
                  <Button
                    type="button"
                    className="font-medium"
                    size="sm"
                    onClick={handleBack}
                    disabled={step <= 0}
                  >
                    Back
                  </Button>
                  <Button type="submit" size="sm" className="font-medium">
                    {step === 1 ? "Submit" : "Next"}
                  </Button>
                </div>
              </form>
            </Form>
          )}
        </CardContent>
      </Card>
    </div>
  );
};
