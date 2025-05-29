import EventCard from "../cards/EventCard";
import { useEffect, useState } from "react";
import { BriefEvent } from "@/interface/BriefEvent";
import api from "@/lib/axios";
import axios from "axios";
import { toast } from "sonner";

export default function RecommendedEventGrid() {
  const [events, setEvents] = useState<BriefEvent[]>([]);

  useEffect(() => {
    async function fetchEvents() {
      try {
        const res = await api.get("/event/recommendation");
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

  return (
    <section className="flex flex-col w-full items-center gap-6">
      <span className="text-4xl font-semibold self-start py-2">
        Events You May Like
      </span>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 gap-y-8 py-4">
        {events.map((event) => (
          <EventCard event={event} />
        ))}
      </div>
    </section>
  );
}
