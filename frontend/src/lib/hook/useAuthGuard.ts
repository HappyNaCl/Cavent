import { useAuth } from "@/components/provider/AuthProvider";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export const useAuthGuard = () => {
  const { user, token: accessToken } = useAuth();
  const nav = useNavigate();

  useEffect(() => {
    if (!user || !accessToken) {
      console.log("User not found, redirecting to auth page");
      nav("/auth");
    }
  }, [user, accessToken, nav]);

  return user;
};
