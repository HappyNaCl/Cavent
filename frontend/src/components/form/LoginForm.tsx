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
import { loginSchema } from "@/lib/schema/LoginSchema";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import LabeledCheckbox from "../input/LabeledCheckbox";
import { useAuth } from "../provider/AuthProvider";
import { useNavigate } from "react-router";
import { toast } from "sonner";
import api from "@/lib/axios";

export default function LoginForm() {
  const { login } = useAuth();
  const nav = useNavigate();

  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      email: "",
      password: "",
      rememberMe: false,
    },
  });

  const onSubmit = async (data: z.infer<typeof loginSchema>) => {
    const formData = new FormData();
    formData.append("email", data.email);
    formData.append("password", data.password);
    formData.append("rememberMe", String(data.rememberMe));

    try {
      const res = await api.post("/auth/login", formData, {
        withCredentials: true,
      });

      if (res.status === 200) {
        const { accessToken, user } = res.data;
        login(user, accessToken);
        toast.success("Login successful!");
        nav("/");
      }
    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
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
          name="rememberMe"
          render={({ field }) => (
            <FormItem>
              <FormControl>
                <LabeledCheckbox
                  id="rememberMe"
                  label="Remember Me"
                  disabled={false}
                  value={field.value}
                  onChange={field.onChange}
                />
              </FormControl>
            </FormItem>
          )}
        />
        <Button
          className="w-full text-2xl py-8 rounded-2xl hover:bg-gray-800"
          type="submit"
        >
          Login
        </Button>
      </form>
    </Form>
  );
}
