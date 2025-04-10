import { z } from "zod";

export const loginSchema = z.object({
  email: z
    .string()
    .email("Invalid Email Address")
    .nonempty("Email is required")
    .toLowerCase(),
  password: z
    .string()
    .min(8, "Password must be at least 8 characters long")
    .nonempty("Password is required"),
  rememberMe: z.boolean().optional(),
});
