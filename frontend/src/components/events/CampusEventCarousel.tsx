import { BriefEvent } from "@/interface/BriefEvent";
import api from "@/lib/axios";
import { useEffect, useState } from "react";
import { useAuth } from "../provider/AuthProvider";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "../ui/carousel";
import EventCard from "../cards/EventCard";
import EventCardSkeleton from "../skeleton/EventCardSkeleton";

type CampusEventCarouselProps = {
  campusId: string;
  campusName?: string;
  unAuthorized: () => void;
};

export default function CampusEventCarousel({
  campusId,
  campusName = "Campus",
  unAuthorized,
}: CampusEventCarouselProps) {
  const [events, setEvents] = useState<BriefEvent[]>([]);
  const [loading, setLoading] = useState(true);
  const { user } = useAuth();

  useEffect(() => {
    const fetchPage = async () => {
      setLoading(true);
      try {
        const response = await api.get(`/campus/${campusId}/events`, {
          params: {
            user: user?.id,
            page: 1,
            limit: 20,
          },
        });
        console.log(response.data);
        if (response.status === 200) {
          const { rows } = response.data.data;
          setEvents(rows);
        }
      } catch (error) {
        console.error("Failed to fetch events:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchPage();
  }, [campusId, user]);

  if (loading) {
    return (
      <div className="flex flex-row gap-4">
        {Array.from({ length: 4 }).map((_, index) => (
          <div key={index} className="basis-1/4">
            <EventCardSkeleton />
          </div>
        ))}
      </div>
    );
  }

  return (
    <section className="relative w-full space-y-12 py-4">
      <h2 className="text-2xl font-semibold">Other Events from {campusName}</h2>
      <Carousel
        className="w-full mx-auto"
        opts={{ align: "start", loop: false }}
      >
        <CarouselContent className="flex">
          {events.map((event) => (
            <CarouselItem className="basis-1/4 flex-shrink-0" key={event.id}>
              <EventCard
                event={event}
                key={event.id}
                onUnauthorized={unAuthorized}
              />
            </CarouselItem>
          ))}
        </CarouselContent>
        <CarouselPrevious className="text-2xl -left-6" />
        <CarouselNext className="text-2xl -right-6" />
      </Carousel>
    </section>
  );
}
