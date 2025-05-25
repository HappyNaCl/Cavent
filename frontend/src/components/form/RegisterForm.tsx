import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import PasswordInput from "@/components/input/PasswordInput";
import TextInput from "@/components/input/TextInput";
import { Button } from "@/components/ui/button";
import { registerSchema } from "@/lib/schema/RegisterSchema";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { toast } from "sonner";
import { useNavigate } from "react-router";
import { useAuth } from "../provider/AuthProvider";
import api from "@/lib/axios";

export default function RegisterForm() {
  const nav = useNavigate();
  const { login } = useAuth();

  const form = useForm({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      fullName: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  const onSubmit = async (data: z.infer<typeof registerSchema>) => {
    const formData = new FormData();
    formData.append("email", data.email.toLowerCase());
    formData.append("password", data.password);
    formData.append("name", data.fullName);
    formData.append("confirmPassword", data.confirmPassword);

    try {
      const res = await api.post("/auth/register", formData, {
        withCredentials: true,
      })

      if (res.status === 201) {
        const { accessToken, user } = res.data;
        login(user, accessToken);
        toast.success("Registration successful!");
        nav("/");
      }
    } catch (error) {
      toast.error(`Error: ${error}}`);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-5">
        <FormField
          control={form.control}
          name="fullName"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Full Name</FormLabel>
              <FormControl>
                <TextInput
                  type="text"
                  id="fullName"
                  placeholder="Enter your full name"
                  onChange={field.onChange}
                  value={field.value}
                  className="border border-gray-300 rounded-lg px-4 py-6"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <TextInput
                  type="email"
                  id="email"
                  placeholder="Enter your email"
                  onChange={field.onChange}
                  value={field.value}
                  className="border border-gray-300 rounded-lg px-4 py-6"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <PasswordInput
                  id="password"
                  onChange={field.onChange}
                  value={field.value}
                  placeholder="Enter your password"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="confirmPassword"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Confirm Password</FormLabel>
              <FormControl>
                <PasswordInput
                  id="confirm-password"
                  onChange={field.onChange}
                  value={field.value}
                  placeholder="Re-enter your password"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button
          className="w-full text-2xl py-8 rounded-2xl hover:bg-gray-800"
          type="submit"
        >
          Register
        </Button>
      </form>
    </Form>
  );
}
