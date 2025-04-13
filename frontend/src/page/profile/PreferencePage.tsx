import { useAuthGuard } from "@/lib/hook/useAuthGuard";

export default function PreferencePage() {
  const user = useAuthGuard();

  if (!user) return null;

  return (
    <div>
      <h1>Set preference pls :c</h1>
    </div>
  );
}
