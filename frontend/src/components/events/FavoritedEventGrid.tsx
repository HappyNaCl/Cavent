import { BriefEvent } from "@/interface/BriefEvent";
import { useEffect, useState } from "react";
import EventCard from "../cards/EventCard";
import { toast } from "sonner";
import axios from "axios";
import api from "@/lib/axios";

export default function FavoritedEventGrid() {
  const [events, setEvents] = useState<BriefEvent[]>([]);

  useEffect(() => {
    async function fetchEvents() {
      try {
        const res = await api.get("/event/favorite");
        if (res.status === 200) {
          setEvents(res.data.data);
        }
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
      <span className="text-4xl font-semibold self-start py-2">Favorites</span>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 gap-y-8 py-4">
        {events.map((event) => (
          <EventCard key={event.id} event={event} />
        ))}
      </div>
    </section>
  );
}
