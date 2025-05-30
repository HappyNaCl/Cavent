import api from "@/lib/axios";
import EventCard from "../cards/EventCard";
import axios from "axios";
import { toast } from "sonner";
import { useEffect, useState } from "react";
import { BriefEvent } from "@/interface/BriefEvent";
import { Link } from "react-router";
import { useAuth } from "../provider/AuthProvider";

type EventGridProps = {
  onUnauthorized: () => void;
};

export default function EventGrid({ onUnauthorized }: EventGridProps) {
  const [events, setEvents] = useState<BriefEvent[]>([]);
  const { user } = useAuth();

  useEffect(() => {
    async function fetchEvents() {
      try {
        const res = await api.get("/event", {
          params: {
            limit: 8,
            user: user?.id,
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
  }, [user]);

  return (
    <section className="flex flex-col w-full items-center gap-6">
      <span className="text-4xl font-semibold self-start py-2">
        Upcoming Events
      </span>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 gap-y-8 py-4">
        {events.map((event) => (
          <EventCard
            key={event.id}
            event={event}
            onUnauthorized={onUnauthorized}
          />
        ))}
      </div>
      <Link to={"/events"} className="text-blue-600 hover:underline text-lg">
        Show More
      </Link>
    </section>
  );
}
