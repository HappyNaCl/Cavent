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
