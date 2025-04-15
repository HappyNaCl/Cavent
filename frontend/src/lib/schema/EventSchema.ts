import { z } from "zod";

export const EventSchema = z.object({
  title: z.string().nonempty("Event title is required"),
  category: z.string().nonempty("Event category is required"),
  eventType: z.enum(["single", "recurring"]).default("single").optional(),
  startDate: z.date().refine((date) => date > new Date(), {
    message: "Start date must be in the future",
  }),
  startTime: z.string().nonempty("Start time is required"),
  location: z.string().nonempty("Location is required"),
  description: z.string().optional(),
});

export const EventDetailsSchema = z.object({
  title: z.string().nonempty("Event title is required"),
  category: z.string().nonempty("Event category is required"),
});

export const EventBannerSchema = z.object({
  banner: z.any().refine((file) => file instanceof File, {
    message: "Banner image is required",
  }),
});

export const EventTicketingSchema = z.object({
  ticketType: z.enum(["free", "paid"]).default("free").optional(),
  ticketPrice: z
    .number()
    .optional()
    .refine((price) => {
      if (price) {
        return price > 0;
      }
      return false;
    }),
  ticketQuantity: z
    .number()
    .optional()
    .refine((quantity) => {
      if (quantity) {
        return quantity > 0;
      }
      return false;
    }),
});
