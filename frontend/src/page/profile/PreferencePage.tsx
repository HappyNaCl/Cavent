import InterestForm from "@/components/form/InterestForm";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";

export default function PreferencePage() {
  const user = useAuthGuard();
  if (!user) return null;

  return (
    <div className="px-20 py-8">
      <InterestForm />
    </div>
  );
}
