import { JoinCampusSchema } from "@/lib/schema/JoinCampusSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { z } from "zod";
import { InputOTP, InputOTPGroup, InputOTPSlot } from "../ui/input-otp";
import { REGEXP_ONLY_CHARS } from "input-otp";
import axios from "axios";
import { useEffect } from "react";
import api from "@/lib/axios";
import { toast } from "sonner";
import { useAuth } from "../provider/AuthProvider";

type JoinCampusFormProps = {
  onSuccess?: () => void;
};

export default function JoinCampusForm({ onSuccess }: JoinCampusFormProps) {
  const form = useForm({
    resolver: zodResolver(JoinCampusSchema),
    defaultValues: {
      inviteCode: "",
    },
  });

  const { setUser } = useAuth();

  const onSubmit = async (data: z.infer<typeof JoinCampusSchema>) => {
    try {
      const res = await api.put("/user/campus", {
        inviteCode: data.inviteCode,
      });
      console.log(res);
      if (res.status === 200) {
        toast.success("Successfully joined the campus!");
        setUser(res.data.data);
        onSuccess?.();
      }

      onSuccess?.();
    } catch (error) {
      let errorMessage = "An unexpected error occurred. Please try again.";
      if (axios.isAxiosError(error)) {
        errorMessage =
          error.response?.data?.error || "An error occurred with the request.";
        form.setError("inviteCode", {
          type: "server",
          message: errorMessage,
        });
      }
    }
  };

  const inviteCode = form.watch("inviteCode");

  useEffect(() => {
    if (inviteCode.length === 6) {
      form.handleSubmit(onSubmit)();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [inviteCode]);

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)}>
        <FormField
          control={form.control}
          name="inviteCode"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-center block">Invite Code</FormLabel>
              <FormControl>
                <div className="flex justify-center">
                  <InputOTP
                    maxLength={6}
                    pattern={REGEXP_ONLY_CHARS}
                    {...field}
                  >
                    <InputOTPGroup>
                      <InputOTPSlot index={0} className="w-14 h-14 text-xl" />
                      <InputOTPSlot index={1} className="w-14 h-14 text-xl" />
                      <InputOTPSlot index={2} className="w-14 h-14 text-xl" />
                      <InputOTPSlot index={3} className="w-14 h-14 text-xl" />
                      <InputOTPSlot index={4} className="w-14 h-14 text-xl" />
                      <InputOTPSlot index={5} className="w-14 h-14 text-xl" />
                    </InputOTPGroup>
                  </InputOTP>
                </div>
              </FormControl>
              <FormMessage className="text-center" />
            </FormItem>
          )}
        />
      </form>
    </Form>
  );
}
