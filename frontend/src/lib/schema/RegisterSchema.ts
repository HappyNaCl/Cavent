import { z } from "zod";

export const registerSchema = z
  .object({
    fullName: z.string().nonempty("Full name is required"),
    email: z
      .string()
      .email("Invalid Email Address")
      .nonempty("Email is required")
      .toLowerCase(),
    password: z
      .string()
      .min(8, "Password must be at least 8 characters long")
      .nonempty("Password is required"),
    confirmPassword: z
      .string()
      .min(8, "Password does not match")
      .nonempty("Please re-enter your password"),
  })
  .superRefine(({ confirmPassword, password }, ctx) => {
    if (confirmPassword !== password) {
      ctx.addIssue({
        code: "custom",
        message: "Passwords do not match",
        path: ["confirmPassword"],
      });
    }
  });
