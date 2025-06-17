import { ChangePasswordForm } from "@/components/form/ChangePasswordForm";
import { SetPasswordForm } from "@/components/form/SetPasswordForm";
import api from "@/lib/axios";
import axios from "axios";
import { useEffect, useState } from "react";
import { toast } from "sonner";

export default function PasswordPage() {
  const [hasPassword, setHasPassword] = useState(false);

  useEffect(() => {
    async function checkHasPassword() {
      try {
        const res = await api.get("/user/has-password");
        console.log(res);
        if (res.status === 200) {
          setHasPassword(res.data.data);
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast(
            error.response?.data.error ||
              "An error occurred while checking password status."
          );
        }
      }
    }

    checkHasPassword();
  }, []);

  return (
    <div className="flex w-full justify-center py-24 min-h-screen">
      {hasPassword ? <ChangePasswordForm /> : <SetPasswordForm />}
    </div>
  );
}
