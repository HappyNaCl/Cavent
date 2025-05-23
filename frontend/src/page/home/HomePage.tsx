import LogoutButton from "@/components/button/LogoutButton";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";

export default function HomePage() {
  const user = useAuthGuard();

  if (!user) return null;

  return (
    <div>
      <h1>Hello, {user.name}</h1>
      <LogoutButton />
    </div>
  );
}
