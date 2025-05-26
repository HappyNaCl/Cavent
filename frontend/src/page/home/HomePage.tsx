import LogoutButton from "@/components/button/LogoutButton";
import { useAuth } from "@/components/provider/AuthProvider";
import api from "@/lib/axios";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";
import axios from "axios";
import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router";
import { toast } from "sonner";

export default function HomePage() {
  const user = useAuthGuard();

  const location = useLocation();
  const nav = useNavigate();
  const { login } = useAuth();

  useEffect(() => {
    async function fetchMe() {
      const params = new URLSearchParams(location.search);
      const token = params.get("token");

      if (token) {
        try {
          const res = await api.get("/auth/me", {
            headers: {
              Authorization: `Bearer ${token}`,
            },
            withCredentials: true,
          });
          const userData = res.data.data.user;
          login(userData, token);

          nav("/", { replace: true });
        } catch (error) {
          if (axios.isAxiosError(error)) {
            const errorMessage =
              error.response?.data?.error || "An error occurred";
            nav("/login");
            toast.error(`Error: ${errorMessage}`);
          }
        }
      }
    }

    fetchMe();
  }, [location, nav, login]);

  if (!user) return null;

  return (
    <div>
      <h1>Hello, {user.name}</h1>
      <LogoutButton />
    </div>
  );
}
