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

export default function RegisterForm() {
  const form = useForm({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      email: "",
      password: "",
      confirmPassword: "",
    },
  });

  const onSubmit = async (data: z.infer<typeof registerSchema>) => {
    console.log(data);
    // Perform login action here
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
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
                  onChange={field.onChange}
                  value={field.value}
                  placeholder="Enter your password"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button
          className="w-full text-3xl py-8 rounded-2xl hover:bg-gray-800"
          type="submit"
        >
          Register
        </Button>
      </form>
    </Form>
  );
}
