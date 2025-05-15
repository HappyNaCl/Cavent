import { useAuth } from "@/components/provider/AuthProvider";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export const useAuthGuardWithRole = (role: string) => {
  const { user } = useAuth();
  const nav = useNavigate();

  useEffect(() => {
    if (!user) {
      console.log("User not found, redirecting to auth page");
      nav("/auth");
    }
    // } else if (user.role !== role) {
    //   console.log("User role not authorized, redirecting to home page");
    //   nav("/");
    // }
  }, [user, nav, role]);

  return user;
};
