import { z } from "zod";

export const EventSchema = z
  .object({
    title: z.string().nonempty("Event title is required"),
    category: z.string().nonempty("Event category is required"),
    eventType: z.enum(["single", "recurring"]).default("single").optional(),
    ticketType: z.enum(["ticketed", "free"]).default("ticketed").optional(),
    startDate: z.date().refine(
      (date) => {
        const now = new Date();
        const today = new Date(
          now.getFullYear(),
          now.getMonth(),
          now.getDate()
        );
        const inputDate = new Date(
          date.getFullYear(),
          date.getMonth(),
          date.getDate()
        );
        console.log(today, inputDate);
        return inputDate >= today;
      },
      {
        message: "Start date must be today or in the future",
      }
    ),
    startTime: z.string().nonempty("Start time is required"),
    endTime: z.string().optional(),
    location: z.string().nonempty("Location is required"),
    description: z.string().optional(),
    banner: z
      .any()
      .refine((file) => file?.size <= 5000000, `Max image size is 5MB.`)
      .refine(
        (file) =>
          ["image/jpeg", "image/jpg", "image/png", "image/webp"].includes(
            file?.type
          ),
        "Only .jpg, .jpeg, .png and .webp formats are supported."
      ),
    tickets: z
      .array(
        z.object({
          name: z.string().min(1, "Ticket name is required"),
          price: z.number().min(0, "Price must be a positive number"),
          quantity: z.number().int().min(1, "Quantity must be at least 1"),
        })
      )
      .optional(),
  })
  .refine(
    (data) =>
      data.ticketType === "free" || (data.tickets && data.tickets.length > 0),
    {
      path: ["tickets"],
      message: "At least one ticket is required when ticket type is ticketed",
    }
  );

export const EventDetailsSchema = z.object({
  title: z.string().nonempty("Event title is required"),
  category: z.string().nonempty("Event category is required"),
  eventType: z.enum(["single", "recurring"]).default("single").optional(),
  ticketType: z.enum(["ticketed", "free"]).default("ticketed").optional(),
  startDate: z.date().refine((date) => {
    const now = new Date();
    const today = new Date(now.getFullYear(), now.getMonth(), now.getDate());
    const inputDate = new Date(
      date.getFullYear(),
      date.getMonth(),
      date.getDate()
    );
    console.log(today, inputDate);
    return inputDate >= today;
  }),
  startTime: z.string().nonempty("Start time is required"),
  endTime: z.string().optional(),
  location: z.string().nonempty("Location is required"),
  description: z.string().optional(),
});

export const EventBannerSchema = z.object({
  banner: z
    .any()
    .refine((file) => file?.size <= 5000000, `Max image size is 5MB.`)
    .refine(
      (file) =>
        ["image/jpeg", "image/jpg", "image/png", "image/webp"].includes(
          file?.type
        ),
      "Only .jpg, .jpeg, .png and .webp formats are supported."
    ),
});

export const EventTicketingSchema = z.object({
  tickets: z
    .array(
      z.object({
        name: z.string().min(1, "Ticket name is required"),
        price: z.number().min(0, "Price must be a positive number"),
        quantity: z.number().int().min(1, "Quantity must be at least 1"),
      })
    )
    .optional(),
});
