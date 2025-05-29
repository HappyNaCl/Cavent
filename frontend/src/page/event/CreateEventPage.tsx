import { EventForm } from "@/components/form/EventForm";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";

export default function CreateEventPage() {
  useAuthGuard();

  return (
    <>
      <EventForm />
    </>
  );
}
