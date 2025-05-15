import { useAuth } from "@/components/provider/AuthProvider";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export const useUnauthGuard = () => {
  const { user } = useAuth();
  const nav = useNavigate();

  useEffect(() => {
    if (user) {
      console.log("User found, redirecting to home page");
      nav("/");
    }
  }, [user, nav]);

  return user;
};
