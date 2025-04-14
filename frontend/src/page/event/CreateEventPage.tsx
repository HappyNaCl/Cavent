import { EventForm } from "@/components/form/EventForm";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";

export default function CreateEventPage() {
  const user = useAuthGuard();

  return (
    <div className="flex flex-col gap-8 items-center justify-center px-36 py-24">
      <EventForm />
    </div>
  );
}
