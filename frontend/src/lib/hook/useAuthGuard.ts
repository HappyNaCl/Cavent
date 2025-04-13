import { useAuth } from "@/components/provider/AuthProvider";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export const useAuthGuard = () => {
  const { user } = useAuth();
  const nav = useNavigate();

  useEffect(() => {
    if (!user) {
      nav("/auth");
    }
  }, [user, nav]);

  return user;
};
