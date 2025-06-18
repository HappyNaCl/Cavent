import FavoriteButton from "@/components/button/FavoriteButton";
import LoginModal, { LoginModalRef } from "@/components/dialog/LoginDialog";
import { ShareModal } from "@/components/dialog/ShareDialog";
import { TicketDialog } from "@/components/dialog/TicketDialog";
import NotFound from "@/components/error/NotFound";
import CampusEventCarousel from "@/components/events/CampusEventCarousel";
import { useAuth } from "@/components/provider/AuthProvider";
import EventDetailSkeleton from "@/components/skeleton/EventDetailSkeleton";
import { Button } from "@/components/ui/button";
import Image from "@/components/ui/image";
import { Event } from "@/interface/Event";
import { Ticket } from "@/interface/Ticket";
import api from "@/lib/axios";
import axios from "axios";
import {
  Calendar,
  Clock,
  MapPin,
  Share2,
  Ticket as TicketIcon,
} from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { useParams } from "react-router";
import { toast } from "sonner";

export default function EventDetailPage() {
  const { id } = useParams();
  const { user } = useAuth();

  const [openShareDialog, setOpenShareDialog] = useState(false);
  const [openTicketDialog, setOpenTicketDialog] = useState(false);
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [event, setEvent] = useState<Event | null>(null);
  const [loading, setLoading] = useState(true);

  const loginModalRef = useRef<LoginModalRef>(null);

  const handleUnauthorized = () => {
    if (loginModalRef.current) {
      loginModalRef.current.open();
    }
  };

  const handleTicketButtonClick = () => {
    if (!user) {
      handleUnauthorized();
      return;
    }
    setOpenTicketDialog(true);
  };

  const dateObj =
    event && event.startTime !== undefined
      ? new Date(event.startTime * 1000)
      : null;

  const formattedDate = dateObj?.toLocaleDateString("en-GB", {
    weekday: "long",
    day: "2-digit",
    month: "long",
    year: "numeric",
  });

  const startTime = dateObj
    ? dateObj.toLocaleTimeString("en-US", {
        hour: "2-digit",
        minute: "2-digit",
        hour12: true,
      })
    : "";

  const endDateObj =
    event && event.endTime !== undefined
      ? new Date(event.endTime * 1000)
      : null;
  const endTime = endDateObj
    ? endDateObj.toLocaleTimeString("en-US", {
        hour: "2-digit",
        minute: "2-digit",
        hour12: true,
      })
    : "";

  useEffect(() => {
    async function fetchEvent(eventId: string) {
      setLoading(true);
      try {
        const res = await api.get(`/event/${eventId}`, {
          params: {
            user: user?.id,
          },
        });
        setEvent(res.data.data);
        setTickets(res.data.data.tickets || []);
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(
            error.response?.data.error || "Failed to fetch event details"
          );
        }
      } finally {
        setLoading(false);
      }
    }

    if (id) {
      fetchEvent(id);
    }
  }, [user, id]);

  if (!loading && event === null) {
    return <NotFound />;
  }

  if (loading || event === null) {
    return <EventDetailSkeleton />;
  }

  return (
    <div className="w-full mx-auto space-y-6">
      <Image
        src={event.bannerUrl}
        alt="Event Banner"
        className="w-full rounded-lg h-96"
      />

      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-800">{event.title}</h1>
        <div className="flex items-center gap-2">
          <FavoriteButton
            isFavorited={event.isFavorited}
            eventId={event.id}
            onSuccess={() => {}}
            onUnauthorized={handleUnauthorized}
          />
          <Button
            variant="ghost"
            size="icon"
            onClick={() => setOpenShareDialog(true)}
          >
            <Share2 />
          </Button>
        </div>
      </div>

      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">Date and Time</h2>
        <div className="flex items-center gap-2 text-gray-600">
          <Calendar className="w-4 h-4" />
          <span>{formattedDate}</span>
        </div>
        <div className="flex items-center gap-2 text-gray-600">
          <Clock className="w-4 h-4" />
          <span>
            {startTime} - {endTime ? endTime : "Done"}
          </span>
        </div>
        <Button variant="outline" size="sm" className="mt-2">
          + Add to Calendar
        </Button>
      </div>

      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">Location</h2>
        <div className="flex items-center gap-2 text-gray-600">
          <MapPin className="w-4 h-4" />
          <span>{event.location}</span>
        </div>
      </div>

      {event.ticketType === "ticketed" && (
        <div className="space-y-2">
          <h2 className="text-lg font-semibold text-gray-700">
            Ticket Information
          </h2>
          <div className="flex flex-col gap-4 text-gray-600">
            {event.tickets.map((ticket) => (
              <div className="flex flex-col w-fit border-4 border-l-amber-200 px-4 py-2 rounded-2xl shadow-md min-w-sm">
                <TicketIcon />
                <span className="flex items-center gap-2" key={ticket.id}>
                  <span className="font-bold">{ticket.name}</span> : $
                  {ticket.price.toFixed(2)} each
                </span>
                <span>{ticket.quantity - ticket.sold} available</span>
              </div>
            ))}
          </div>
          <Button
            className="bg-yellow-400 hover:bg-yellow-500 text-black font-semibold mt-2"
            onClick={handleTicketButtonClick}
          >
            Buy Tickets
          </Button>
        </div>
      )}

      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">Hosted by</h2>
        <div className="flex items-center gap-3">
          <img
            src={event.campus.profileUrl}
            alt={event.campus.name}
            className="w-10 h-10 rounded-full"
          />
          <div>
            <p className="font-medium text-gray-800">{event.campus.name}</p>
            <div className="flex gap-2 mt-1">
              <Button size="sm" variant="outline">
                Contact
              </Button>
              <Button size="sm" variant="outline">
                More
              </Button>
            </div>
          </div>
        </div>
      </div>

      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">
          Event Description
        </h2>
        <p className="text-gray-700 leading-relaxed text-justify">
          {event.description}
        </p>
      </div>

      <CampusEventCarousel
        campusId={event.campus.id}
        campusName={event.campus.name}
        unAuthorized={handleUnauthorized}
      />

      {!user && <LoginModal ref={loginModalRef} />}
      <ShareModal
        open={openShareDialog}
        onOpenChange={setOpenShareDialog}
        eventId={event.id}
      />

      {event.ticketType === "ticketed" && (
        <TicketDialog
          eventId={event.id}
          open={openTicketDialog}
          onOpenChange={setOpenTicketDialog}
          tickets={tickets}
        />
      )}
    </div>
  );
}
