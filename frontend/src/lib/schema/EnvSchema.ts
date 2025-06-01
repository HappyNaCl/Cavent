import { z } from "zod";

const envSchema = z.object({
  VITE_BACKEND_URL: z.string().nonempty(),
  VITE_APP_URL: z.string().nonempty(),
});

export const env = envSchema.parse(import.meta.env);
