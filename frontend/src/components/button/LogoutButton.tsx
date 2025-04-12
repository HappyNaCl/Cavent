import { toast } from "sonner";
import { Button } from "../ui/button";
import axios from "axios";
import { env } from "@/lib/schema/EnvSchema";
import { useNavigate } from "react-router";

export default function LogoutButton() {
  const nav = useNavigate();

  const handleClick = async () => {
    try {
      const res = await axios.post(
        `${env.VITE_BACKEND_URL}/api/auth/logout`,
        {},
        { withCredentials: true }
      );

      if (res.status === 200) {
        toast.success("Logged out successfully");
        nav("/auth");
      }
    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  };
  return <Button onClick={handleClick}>Logout</Button>;
}
