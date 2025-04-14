import { z } from "zod";

export const InterestSchema = z.object({
  interests: z.array(z.string()).nonempty("Interest is required"),
});
