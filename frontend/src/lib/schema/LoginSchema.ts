import { z } from "zod";

export const loginSchema = z.object({
  email: z
    .string()
    .email("Invalid Email Address")
    .nonempty("Email is required")
    .toLowerCase(),
  password: z.string().nonempty("Password is required"),
  rememberMe: z.boolean().optional(),
});
