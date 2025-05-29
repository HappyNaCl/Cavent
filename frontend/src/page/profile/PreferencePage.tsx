import InterestForm from "@/components/form/InterestForm";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";

export default function PreferencePage() {
  useAuthGuard();

  return (
    <>
      <InterestForm />
    </>
  );
}
