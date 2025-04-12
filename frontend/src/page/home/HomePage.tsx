import LogoutButton from "@/components/button/LogoutButton";
import { useAuth } from "@/components/provider/AuthProvider";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export default function HomePage() {
  const { user } = useAuth();
  const nav = useNavigate();

  useEffect(() => {
    if (!user) {
      nav("/auth");
    }
  }, [user, nav]);

  if (!user) return null;

  return (
    <div>
      <h1>Hello, {user.name}</h1>
      <LogoutButton />
    </div>
  );
}
