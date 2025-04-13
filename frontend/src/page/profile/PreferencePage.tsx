import { Button } from "@/components/ui/button";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";
import { env } from "@/lib/schema/EnvSchema";
import axios from "axios";
import { useNavigate } from "react-router";
import { toast } from "sonner";

export default function PreferencePage() {
  const user = useAuthGuard();
  const nav = useNavigate();
  if (!user) return null;

  const onSubmit = async () => {
    try {
      const res = await axios.put(
        `${env.VITE_BACKEND_URL}/api/user/preference`,
        {},
        {
          withCredentials: true,
        }
      );

      if (res.status === 200) {
        alert("Preference updated successfully");
        nav("/");
      }
    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  };

  return (
    <div>
      <h1>Set preference pls :c</h1>
      <Button onClick={onSubmit}>Update</Button>
    </div>
  );
}
