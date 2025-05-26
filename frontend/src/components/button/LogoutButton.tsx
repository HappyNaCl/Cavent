import { toast } from "sonner";
import { Button } from "../ui/button";
import { useNavigate } from "react-router";
import api from "@/lib/axios";
import { useAuth } from "../provider/AuthProvider";

export default function LogoutButton() {
  const nav = useNavigate();
  const { logout } = useAuth();

  const handleClick = async () => {
    try {
      const res = await api.get("/auth/logout", {
        withCredentials: true,
      });
      if (res.status === 200) {
        toast.success("Logout successful!");
        logout();
        nav("/login");
      }
    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  };
  return <Button onClick={handleClick}>Logout</Button>;
}
