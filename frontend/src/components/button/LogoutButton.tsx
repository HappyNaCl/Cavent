import { toast } from "sonner";
import { useNavigate } from "react-router";
import api from "@/lib/axios";
import { useAuth } from "../provider/AuthProvider";
import { LogOutIcon } from "lucide-react";

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
        nav("/");
      }
    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  };
  return (
    <div className="flex items-center gap-2 w-full py-2" onClick={handleClick}>
      <LogOutIcon />
      Logout
    </div>
  );
}
