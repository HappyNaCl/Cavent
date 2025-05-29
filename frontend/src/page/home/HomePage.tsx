import EventCard from "@/components/cards/EventCard";
import { BriefEvent } from "@/interface/BriefEvent";
import api from "@/lib/axios";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";
import axios from "axios";
import { useEffect, useState } from "react";
import { toast } from "sonner";

export default function HomePage() {
  const user = useAuthGuard();

  const [events, setEvents] = useState<BriefEvent[]>([]);

  useEffect(() => {
    async function fetchEvents() {
      try {
        const res = await api.get("/event", {
          params: {
            limit: 8,
          },
        });
        setEvents(res.data.data);
      } catch (error) {
        if (axios.isAxiosError(error)) {
          const errorMessage =
            error.response?.data?.error || "An error occurred";
          toast.error(`Error: ${errorMessage}`);
        }
      }
    }

    fetchEvents();
  }, []);

  if (!user) return null;

  return (
    <section className="flex flex-col w-full items-center py-12 px-24">
      <span className="text-4xl font-semibold self-start py-2">Events</span>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 gap-y-8 py-4">
        {events.map((event) => (
          <EventCard event={event} />
        ))}
      </div>
    </section>
  );
}
