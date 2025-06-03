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
import { toast } from "sonner";
import axios from "axios";
import { useEffect } from "react";

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

  const onSubmit = async (data: z.infer<typeof JoinCampusSchema>) => {
    try {
      console.log("Submitting invite code:", data.inviteCode);

      onSuccess?.();
    } catch (error) {
      if (axios.isAxiosError(error)) {
        toast.error(
          `Error: ${error.response?.data?.error || "An error occurred"}`
        );
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
              <FormMessage />
            </FormItem>
          )}
        />
      </form>
    </Form>
  );
}
