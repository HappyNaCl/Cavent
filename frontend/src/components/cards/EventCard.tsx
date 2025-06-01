import { BriefEvent } from "@/interface/BriefEvent";
import { Ticket, Star } from "lucide-react";
import Image from "../ui/image";
import FavoriteButton from "../button/FavoriteButton";
import { MouseEvent, useState } from "react";
import { useNavigate } from "react-router";

type Props = {
  event: BriefEvent;
  onUnauthorized: () => void;
};

export default function EventCard({ event, onUnauthorized }: Props) {
  const nav = useNavigate();

  const dateObj = new Date(event.startDate * 1000);

  const day = dateObj.getDate();
  const month = dateObj
    .toLocaleDateString("en-US", { month: "short" })
    .toUpperCase();
  const startTime = dateObj.toLocaleTimeString("en-US", {
    hour: "2-digit",
    minute: "2-digit",
    hour12: true,
  });

  const endDateObj = new Date(event.endDate ? event.endDate * 1000 : 1000);
  const endTime = endDateObj.toLocaleTimeString("en-US", {
    hour: "2-digit",
    minute: "2-digit",
    hour12: true,
  });

  const [favCount, setFavCount] = useState(event.favoriteCount);

  const handleCardClick = (e: MouseEvent<HTMLDivElement>) => {
    e.stopPropagation();
    nav("/event/" + event.id);
  };

  return (
    <div
      onClick={handleCardClick}
      className="max-w-sm rounded-xl overflow-hidden shadow cursor-pointer bg-white hover:shadow-lg transition-shadow duration-300"
    >
      <div className="relative w-full h-40">
        <Image
          src={event.bannerUrl}
          alt={event.title}
          className="w-full h-full object-cover"
        />
        <div className="absolute top-2 left-2 bg-yellow-400 text-black text-xs font-semibold px-2 py-1 rounded">
          {event.categoryName}
        </div>
        <div className="absolute top-2 right-2">
          <FavoriteButton
            eventId={event.id}
            onUnauthorized={onUnauthorized}
            onSuccess={(newCount) => {
              setFavCount(newCount);
            }}
            isFavorited={event.isFavorited}
          />
        </div>
      </div>

      <div className="p-4">
        <div className="flex gap-3">
          <div className="text-center">
            <div className="text-xs font-medium text-purple-600">{month}</div>
            <div className="text-lg font-bold text-gray-800 leading-none">
              {day}
            </div>
          </div>

          <div className="flex-1 space-y-1">
            <h2 className="text-lg font-semibold text-gray-900 leading-snug line-clamp-2">
              {event.title}
            </h2>
            <p className="text-md text-gray-500">{event.campusName}</p>

            <p className="text-sm text-gray-500">
              {startTime || "00:00 AM"} - {event.endDate ? endTime : "Done"}
            </p>

            <div className="flex items-center gap-3 text-sm text-gray-700 font-medium mt-1">
              <div className="flex items-center gap-1">
                <Ticket className="w-4 h-4 relative top-[1px]" />
                {event.ticketType === "free"
                  ? "Free"
                  : `$${event.ticketPrice?.toFixed(0)}`}
              </div>
              <div className="flex items-center gap-1 text-blue-600">
                <Star className="w-4 h-4 fill-current relative top-[1px]" />
                {favCount} interested
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
