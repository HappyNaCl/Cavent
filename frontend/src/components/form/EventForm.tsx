"use client";

import { useEffect, useState } from "react";
import { useFieldArray, useForm } from "react-hook-form";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
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
import { z, ZodError } from "zod";
import {
  EventBannerSchema,
  EventDetailsSchema,
  EventSchema,
  EventTicketingSchema,
} from "@/lib/schema/EventSchema";
import TextInput from "../input/TextInput";
import FreeStamp from "@/assets/FreeStamp.png";
import TicketStamp from "@/assets/TicketStamp.png";
import { Input } from "../ui/input";
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group";
import {
  Popover,
  PopoverTrigger,
  PopoverContent,
} from "@/components/ui/popover";
import { format } from "date-fns";
import { Calendar } from "@/components/ui/calendar";
import { Textarea } from "../ui/textarea";
import { LucideCalendarDays, LucideClock, LucideMapPin } from "lucide-react";

export const EventForm = () => {
  const [step, setStep] = useState(0);
  const totalSteps = 4;

  const crumbs = ["Edit", "Banner", "Ticketing", "Review"];
  const titles = [
    { title: "Create a New Event", description: "Enter your event details" },
    { title: "Event Banner", description: "Upload your event banner" },
    { title: "Ticketing", description: "Set up your ticketing options" },
    { title: "Review", description: "Review your event details" },
  ];

  const now = new Date();

  const form = useForm<z.infer<typeof EventSchema>>({
    // resolver: zodResolver(EventSchema),
    defaultValues: {
      title: "",
      category: "",
      startDate: now,
      startTime: format(now, "HH:mm"),
      location: "",
      description: "",
      ticketType: "ticketed",
      tickets: [{ name: "", price: 0, quantity: 0 }],
    },
  });

  const { control, getValues, watch } = form;

  const { fields, append, remove } = useFieldArray({
    control,
    name: "tickets",
  });

  useEffect(() => {
    if (form.getValues("ticketType") === "ticketed" && fields.length === 0) {
      append({ name: "", price: 0, quantity: 0 });
    } else if (form.getValues("ticketType") === "free") {
      form.setValue("tickets", undefined);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [watch("ticketType")]);

  const onSubmit = async (formData: unknown) => {
    const values = form.getValues();
    const stepSchema = [
      EventDetailsSchema,
      EventBannerSchema,
      EventTicketingSchema,
    ];
    console.log(values);

    const currentSchema = stepSchema[step];

    const result = currentSchema.safeParse(values);
    if (!result.success) {
      const error: ZodError = result.error;
      toast.error(error.errors[0].message);
      return;
    }

    console.log(step);
    fields.forEach((field) => {
      console.log(field);
    });

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

  const formValuesForReview = watch();

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
      <Card className="shadow-sm px-8">
        <CardHeader>
          <CardTitle className="text-3xl font-semibold">
            {titles[step].title}
          </CardTitle>
          <CardDescription>{titles[step].description}</CardDescription>
        </CardHeader>
        <CardContent>
          {step === 0 && (
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="grid gap-y-8"
              >
                <div className="flex flex-col gap-4">
                  <span className="text-2xl font-bold">Event Details</span>
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
                            placeholder="Category"
                            onChange={field.onChange}
                            value={field.value}
                            className="border border-gray-300 rounded-lg px-4 py-6"
                          />
                        </FormControl>
                        <FormDescription>
                          Enter the event's category
                        </FormDescription>
                      </FormItem>
                    )}
                  />
                </div>

                <div className="flex flex-col gap-8">
                  <span className="text-2xl font-bold">Date & Time</span>
                  <FormField
                    key="eventType"
                    control={form.control}
                    name="eventType"
                    render={({ field }) => (
                      <FormItem>
                        <span className="font-semibold font-xl mb-2">
                          Event Type
                        </span>
                        <RadioGroup
                          onValueChange={field.onChange}
                          defaultValue={field.value}
                          className="flex flex-row space-x-8"
                        >
                          {[
                            { label: "Single", value: "single" },
                            { label: "Recurring", value: "recurring" },
                          ].map((option) => (
                            <FormItem
                              key={option.value}
                              className="flex items-center space-x-3 space-y-0"
                            >
                              <FormControl>
                                <RadioGroupItem
                                  value={option.value}
                                  className="border-black"
                                />
                              </FormControl>
                              <FormLabel className="font-normal">
                                {option.label}
                              </FormLabel>
                            </FormItem>
                          ))}
                        </RadioGroup>
                      </FormItem>
                    )}
                  />

                  <div className="flex gap-8">
                    <FormField
                      control={form.control}
                      name="startDate"
                      render={({ field }) => (
                        <FormItem className="flex flex-col flex-1">
                          <FormLabel>Start Date</FormLabel>
                          <Popover>
                            <PopoverTrigger asChild>
                              <button
                                type="button"
                                className={cn(
                                  "w-full text-left border border-gray-300 rounded-lg p-3",
                                  !field.value && "text-muted-foreground"
                                )}
                              >
                                {field.value ? (
                                  format(field.value, "PPP")
                                ) : (
                                  <span>Pick a date</span>
                                )}
                              </button>
                            </PopoverTrigger>
                            <PopoverContent className="w-auto p-0">
                              <Calendar
                                mode="single"
                                selected={field.value}
                                onSelect={field.onChange}
                                initialFocus
                              />
                            </PopoverContent>
                          </Popover>
                          <FormDescription>
                            The event start date
                          </FormDescription>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                    <div className="flex flex-1 flex-col gap-2">
                      <FormField
                        control={form.control}
                        name="startTime"
                        render={({ field }) => (
                          <FormItem className="flex flex-col flex-1">
                            <FormLabel>Start Time</FormLabel>
                            <Input
                              type="time"
                              value={field.value}
                              onChange={field.onChange}
                              className="h-12"
                            />
                            <FormDescription>
                              The event start time
                            </FormDescription>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                    <div className="flex flex-1 flex-col gap-2">
                      <FormField
                        control={form.control}
                        name="endTime"
                        render={({ field }) => (
                          <FormItem className="flex flex-col flex-1">
                            <FormLabel>End Time</FormLabel>
                            <Input
                              type="time"
                              value={field.value}
                              onChange={field.onChange}
                              className="h-12"
                            />
                            <FormDescription>
                              The event end time
                            </FormDescription>
                            <FormMessage />
                          </FormItem>
                        )}
                      />
                    </div>
                  </div>
                </div>

                <div className="flex flex-col gap-4">
                  <span className="text-2xl font-bold">Location</span>
                  <FormField
                    key="location"
                    control={form.control}
                    name="location"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Where will your event take place?</FormLabel>
                        <FormControl>
                          <TextInput
                            type="text"
                            id="location"
                            placeholder="Location"
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
                </div>

                <div className="flex flex-col gap-4">
                  <span className="text-2xl font-bold">Description</span>
                  <FormField
                    key="description"
                    control={form.control}
                    name="description"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Event Description</FormLabel>
                        <FormControl>
                          <Textarea
                            value={field.value}
                            onChange={field.onChange}
                            className={"resize-none h-40"}
                            rows={40}
                            placeholder="Description"
                          />
                        </FormControl>
                        <FormDescription>
                          Enter your event's title
                        </FormDescription>
                      </FormItem>
                    )}
                  />
                </div>

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
                <FormField
                  key="banner"
                  control={form.control}
                  name="banner"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Upload Image</FormLabel>
                      <FormControl>
                        <div className="grid w-full max-w-sm items-center gap-1.5">
                          <Input
                            id="picture"
                            type="file"
                            accept="image/*"
                            onChange={(e) => {
                              const file = e.target.files?.[0];
                              if (file) {
                                field.onChange(file);
                              }
                            }}
                          />
                        </div>
                      </FormControl>
                      <FormDescription>
                        Feature Image must be less than 5MB.
                        <br />
                        Valid file formats: .jpg, .jpeg, .png
                      </FormDescription>
                    </FormItem>
                  )}
                />

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
                    {step === totalSteps - 1 ? "Submit" : "Next"}
                  </Button>
                </div>
              </form>
            </Form>
          )}

          {step === 2 && (
            <Form {...form}>
              <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="grid gap-y-12"
              >
                <FormField
                  control={form.control}
                  name="ticketType"
                  render={({ field }) => (
                    <FormItem className="space-y-3 w-11/12">
                      <FormLabel className="text-2xl">
                        What type of event are you running?
                      </FormLabel>
                      <FormControl>
                        <RadioGroup
                          onValueChange={field.onChange}
                          value={field.value}
                          className="grid grid-cols-1 sm:grid-cols-2 gap-4"
                        >
                          {[
                            {
                              label: "Ticketed Event",
                              value: "ticketed",
                              icon: TicketStamp,
                              desc: "My event requires tickets for entry",
                            },
                            {
                              label: "Free Event",
                              value: "free",
                              icon: FreeStamp,
                              desc: "I'm running a free event",
                            },
                          ].map((option) => (
                            <label
                              key={option.value}
                              htmlFor={option.value}
                              className={cn(
                                "cursor-pointer rounded-xl border px-4 py-8 transition-all",
                                field.value === option.value
                                  ? "border-primary bg-blue-200/10"
                                  : "border-primary/30"
                              )}
                            >
                              <div className="flex items-center">
                                <RadioGroupItem
                                  value={option.value}
                                  id={option.value}
                                  className="hidden"
                                />
                                <div className="flex flex-col items-center justify-center w-full">
                                  <img
                                    src={option.icon}
                                    alt=""
                                    className="w-[20%] mb-4"
                                  />
                                  <span className="font-semibold text-2xl">
                                    {option.label}
                                  </span>
                                  <span className="font-light text-xl text-muted-foreground">
                                    {option.desc}
                                  </span>
                                </div>
                              </div>
                            </label>
                          ))}
                        </RadioGroup>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                {getValues("ticketType") === "ticketed" && (
                  <div className="space-y-6">
                    <div className="py-2">
                      <span className="text-2xl font-semibold">
                        Ticketing Details
                      </span>
                    </div>
                    {fields.map((field, index) => (
                      <div
                        key={field.id}
                        className="flex flex-row w-full gap-4 justify-center items-end"
                      >
                        <FormField
                          control={form.control}
                          name={`tickets.${index}.name`}
                          render={({ field }) => (
                            <FormItem className="flex-1">
                              <FormLabel>Ticket Name</FormLabel>
                              <FormControl>
                                <Input
                                  {...field}
                                  placeholder="Name"
                                  className="border border-gray-300 rounded-lg px-4 py-6"
                                />
                              </FormControl>
                            </FormItem>
                          )}
                        />

                        <FormField
                          control={form.control}
                          name={`tickets.${index}.price`}
                          render={({ field }) => (
                            <FormItem className="flex-1">
                              <FormLabel>Ticket Price</FormLabel>
                              <FormControl>
                                <Input
                                  value={field.value}
                                  onChange={(e) =>
                                    field.onChange(e.target.valueAsNumber || 0)
                                  }
                                  type="number"
                                  placeholder="Price"
                                  className="border border-gray-300 rounded-lg px-4 py-6"
                                />
                              </FormControl>
                            </FormItem>
                          )}
                        />

                        <FormField
                          control={form.control}
                          name={`tickets.${index}.quantity`}
                          render={({ field }) => (
                            <FormItem className="flex-1">
                              <FormLabel>Ticket Quantity</FormLabel>
                              <FormControl>
                                <Input
                                  value={field.value}
                                  onChange={(e) =>
                                    field.onChange(e.target.valueAsNumber || 0)
                                  }
                                  type="number"
                                  placeholder="Quantity"
                                  className="border border-gray-300 rounded-lg px-4 py-6"
                                />
                              </FormControl>
                            </FormItem>
                          )}
                        />

                        <Button
                          type="button"
                          variant="destructive"
                          className="hover:bg-red-500/80"
                          size="icon"
                          disabled={fields.length <= 1}
                          onClick={() => remove(index)}
                        >
                          ✕
                        </Button>
                      </div>
                    ))}

                    <Button
                      type="button"
                      variant="outline"
                      onClick={() => {
                        append({ name: "", price: 0, quantity: 0 });
                        console.log(getValues("tickets"));
                      }}
                    >
                      + Add Ticket
                    </Button>
                  </div>
                )}

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
                    {step === totalSteps - 1 ? "Submit" : "Next"}
                  </Button>
                </div>
              </form>
            </Form>
          )}
          {step === 3 && (
            <Form {...form}>
              {" "}
              {/* Still good to wrap with Form for context if needed */}
              <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="grid gap-y-8" // Increased gap slightly
              >
                <div className="border-2 border-gray-200 rounded-2xl p-6 md:p-10 shadow-md">
                  {" "}
                  {/* Softer border, shadow */}
                  <div className="flex flex-col gap-6">
                    {formValuesForReview.banner && ( // Check if banner exists
                      <img
                        className="w-full h-auto max-h-[400px] object-cover rounded-lg" // max-h for responsiveness
                        src={URL.createObjectURL(
                          formValuesForReview.banner as File
                        )}
                        alt={`${formValuesForReview.title || "Event"} banner`}
                      />
                    )}
                    <div className="flex flex-col gap-8 md:gap-10">
                      {" "}
                      {/* Adjusted gaps */}
                      <div className="flex flex-col gap-1 py-3">
                        <span className="font-bold text-3xl md:text-5xl break-words">
                          {formValuesForReview.title}
                        </span>
                        <span className="text-lg md:text-xl text-gray-500">
                          {formValuesForReview.eventType === "recurring"
                            ? "Recurring Event"
                            : "Single Event"}
                        </span>
                      </div>
                      <div className="flex flex-col gap-2">
                        <span className="text-xl md:text-2xl font-semibold">
                          Date and Time
                        </span>
                        {formValuesForReview.startDate && (
                          <span className="text-lg md:text-xl flex flex-row gap-3 items-center">
                            {" "}
                            {/* Reduced gap */}
                            <LucideCalendarDays className="text-primary" />
                            {format(
                              new Date(formValuesForReview.startDate),
                              "PPP"
                            )}{" "}
                            {/* Ensure startDate is a Date */}
                          </span>
                        )}
                        <span className="text-lg md:text-xl flex flex-row gap-3 items-center">
                          <LucideClock className="text-primary" />
                          {formValuesForReview.startTime || "Not set"}
                          {formValuesForReview.endTime &&
                            ` - ${formValuesForReview.endTime}`}
                        </span>
                      </div>
                      <div className="flex flex-col gap-2">
                        <span className="text-xl md:text-2xl font-semibold">
                          Location
                        </span>
                        <span className="text-lg md:text-xl flex flex-row gap-3 items-center">
                          <LucideMapPin className="text-primary" />
                          {formValuesForReview.location || "Not specified"}
                        </span>
                      </div>
                      <div className="flex flex-col gap-2">
                        <span className="text-xl md:text-2xl font-semibold">
                          Event Description
                        </span>
                        <pre className="text-base md:text-lg font-sans text-gray-700 whitespace-pre-wrap break-words bg-gray-50 p-3 rounded-md">
                          {" "}
                          {/* Improved styling */}
                          {formValuesForReview.description ||
                            "No description provided."}
                        </pre>
                      </div>
                    </div>
                  </div>
                  {/* MODIFIED TICKETING DISPLAY */}
                  {formValuesForReview.ticketType === "ticketed" &&
                    formValuesForReview.tickets &&
                    formValuesForReview.tickets.length > 0 && (
                      <div className="flex flex-col gap-6 mt-8 pt-6 border-t border-gray-200">
                        {" "}
                        {/* Added spacing and separator */}
                        <span className="text-xl md:text-2xl font-semibold">
                          Ticketing
                        </span>
                        {formValuesForReview.tickets.map((ticket, index) => (
                          <div
                            // Use the id from the `fields` array for a stable key if available,
                            // otherwise fallback to index. This assumes `fields` and `formValuesForReview.tickets`
                            // maintain the same order.
                            key={fields[index]?.id || `review-ticket-${index}`}
                            className="p-4 border border-gray-200 rounded-lg shadow-sm bg-white" // Card-like appearance for each ticket
                          >
                            <p className="font-medium text-lg">
                              Ticket:{" "}
                              <span className="font-normal">
                                {ticket.name || "Unnamed Ticket"}
                              </span>
                            </p>
                            <p className="text-gray-600">
                              Price:{" "}
                              <span className="font-normal">
                                ${(ticket.price || 0).toFixed(2)}
                              </span>
                            </p>
                            <p className="text-gray-600">
                              Quantity:{" "}
                              <span className="font-normal">
                                {ticket.quantity || 0}
                              </span>
                            </p>
                          </div>
                        ))}
                      </div>
                    )}
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
                    Submit Event
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
