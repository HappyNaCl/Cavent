import { BriefEvent } from "@/interface/BriefEvent";
import { Calendar, MapPin, Ticket } from "lucide-react";
import Image from "../ui/image";

type Props = {
  event: BriefEvent;
};

export default function EventCard({ event }: Props) {
  const startDate = new Date(event.startDate * 1000).toLocaleDateString(
    "en-US",
    {
      weekday: "short",
      month: "short",
      day: "numeric",
      year: "numeric",
    }
  );

  return (
    <div className="max-w-sm rounded-2xl overflow-hidden shadow-md bg-white hover:shadow-lg transition">
      <div className="relative w-full h-48 overflow-hidden rounded-t-2xl">
        <Image
          src={event.bannerUrl}
          alt={event.title}
          className="w-full h-48 object-cover"
        />
        <div className="absolute bottom-2 left-2 bg-yellow-500 bg-opacity-60 text-white text-xs font-semibold px-3 py-1 rounded-full select-none">
          {event.categoryName}
        </div>
      </div>

      <div className="p-4 space-y-2">
        <div className="flex justify-between items-center">
          <h2 className="text-xl font-semibold">{event.title}</h2>
          <span className="text-xs bg-yellow-100 text-yellow-600 px-2 py-1 rounded-full capitalize">
            {event.ticketType}
          </span>
        </div>
        <p className="text-gray-500 text-sm">{event.campusName}</p>

        <div className="flex items-center gap-2 text-gray-700 text-sm">
          <Calendar className="w-4 h-4" /> {startDate}
        </div>

        <div className="flex items-center gap-2 text-gray-700 text-sm">
          <MapPin className="w-4 h-4" /> {event.location}
        </div>

        <div className="flex items-center gap-2 text-sm font-medium">
          <Ticket className="w-4 h-4 text-gray-700" />
          {event.ticketType === "free" ? (
            <span className="text-green-600">Free Entry</span>
          ) : (
            <span className="text-gray-900">
              ${event.ticketPrice?.toFixed(2)}
            </span>
          )}
        </div>
      </div>
    </div>
  );
}
