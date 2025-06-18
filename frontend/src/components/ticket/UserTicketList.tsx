import { UserTicket } from "@/interface/UserTicket";
import { Ticket } from "lucide-react";

type UserTicketListProps = {
  userTickets: UserTicket[];
};

export default function UserTicketList({ userTickets }: UserTicketListProps) {
  return (
    <div className="flex flex-col items-center">
      {userTickets.map((event) => (
        <div
          key={event.eventId}
          className="rounded-xl border shadow-sm p-4 bg-white w-fit"
        >
          <div className="flex flex-col gap-3">
            <h2 className="text-xl font-semibold mb-2">{event.eventTitle}</h2>
            <span className="text-lg font-normal text-gray-600 mb-4">
              {event.startTime} - {event.endTime}
            </span>
          </div>
          <ul className="space-y-1">
            {event.tickets.map((ticket) => (
              <li
                key={ticket.id}
                className="px-3 py-2 min-w-2xl bg-gray-100 rounded-md text-md flex justify-around items-center gap-2 hover:bg-gray-200 transition-colors cursor-pointer"
              >
                <Ticket /> <span>{ticket.name}</span> <span>{ticket.id}</span>
              </li>
            ))}
          </ul>
        </div>
      ))}
    </div>
  );
}
