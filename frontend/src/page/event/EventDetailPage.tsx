import { Button } from "@/components/ui/button";
import { Calendar, Clock, MapPin, Star } from "lucide-react";
import { useParams } from "react-router";

export default function EventDetailPage() {
  const { id } = useParams();

  return (
    <div className="mx-auto space-y-6">
      <img
        src="https://placehold.co/1600x600"
        alt="Sound Of Christmas 2023"
        className="w-full rounded-lg"
      />

      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-800">
          Sound Of Christmas 2023
        </h1>
        <div className="flex items-center gap-2">
          <Button variant="ghost" size="icon">
            <Star className="w-5 h-5" />
          </Button>
          <Button variant="ghost" size="icon">
            <svg
              width="20"
              height="20"
              fill="currentColor"
              className="text-gray-600"
            >
              <path d="M15 8a3 3 0 10-2.83-2H8a3 3 0 000 6h4.17A3 3 0 1015 8zM8 10a2 2 0 110-4 2 2 0 010 4zm7 0a2 2 0 110-4 2 2 0 010 4z"></path>
            </svg>
          </Button>
        </div>
      </div>

      {/* Date and Time */}
      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">Date and Time</h2>
        <div className="flex items-center gap-2 text-gray-600">
          <Calendar className="w-4 h-4" />
          <span>Saturday, 2 December 2023</span>
        </div>
        <div className="flex items-center gap-2 text-gray-600">
          <Clock className="w-4 h-4" />
          <span>6:30PM - 9:30PM</span>
        </div>
        <Button variant="outline" size="sm" className="mt-2">
          + Add to Calendar
        </Button>
      </div>

      {/* Location */}
      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">Location</h2>
        <div className="flex items-center gap-2 text-gray-600">
          <MapPin className="w-4 h-4" />
          <span>
            St Gonsalo Garcia Rang Mandir, Near Junction Of Link Road, Vasai
            West, Mumbai, MH
          </span>
        </div>
        {/* <img
          src="https://maps.googleapis.com/maps/api/staticmap?center=Vasai+West+Mumbai&zoom=15&size=600x300&key=YOUR_API_KEY"
          alt="Map"
          className="rounded-lg border"
        /> */}
      </div>

      {/* Ticket Info */}
      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">
          Ticket Information
        </h2>
        <p className="text-gray-700">🎫 Standard Ticket: ₹200 each</p>
        <Button className="bg-yellow-400 hover:bg-yellow-500 text-black font-semibold mt-2">
          Buy Tickets
        </Button>
      </div>

      {/* Hosted By */}
      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">Hosted by</h2>
        <div className="flex items-center gap-3">
          <img
            src="/cry-logo.png"
            alt="Cry Youth Movement"
            className="w-10 h-10 rounded-full"
          />
          <div>
            <p className="font-medium text-gray-800">Cry Youth Movement</p>
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

      {/* Event Description */}
      <div className="space-y-2">
        <h2 className="text-lg font-semibold text-gray-700">
          Event Description
        </h2>
        <p className="text-gray-700 leading-relaxed">
          Get ready to kick off the Christmas season in Mumbai with SOUND OF
          CHRISTMAS – your favourite LIVE Christmas concert!
        </p>
      </div>
    </div>
  );
}
