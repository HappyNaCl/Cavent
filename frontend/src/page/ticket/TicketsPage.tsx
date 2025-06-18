import UserTicketList from "@/components/ticket/UserTicketList";
import { UserTicket } from "@/interface/UserTicket";
import api from "@/lib/axios";
import axios from "axios";
import { TicketIcon } from "lucide-react";
import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { Link } from "react-router";
import { toast } from "sonner";

export default function TicketsPage() {
  const [tickets, setTickets] = useState<UserTicket[]>([]);

  useEffect(() => {
    async function fetchTickets() {
      try {
        const res = await api.get("/user/tickets");
        if (res.status === 200) {
          console.log(res);
          setTickets(res.data.data);
        } else {
          toast.error("Failed to fetch tickets.");
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(
            error.response?.data.error ||
              "An error occurred while fetching tickets."
          );
        }
      }
    }

    fetchTickets();
  }, []);

  return (
    <main className="container mx-auto max-w-4xl px-4 py-12 w-fit flex flex-col items-center">
      <div className="mb-10 text-center">
        <h1 className="text-4xl font-bold tracking-tight">My Tickets</h1>
        <p className="mt-2 text-lg text-muted-foreground">
          Here are all your purchased tickets. Get ready for your next
          experience!
        </p>
      </div>

      {tickets.length > 0 ? (
        <UserTicketList userTickets={tickets} />
      ) : (
        <div className="text-center py-20 px-6 rounded-lg border-2 border-dashed w-full">
          <TicketIcon className="mx-auto h-12 w-12 text-muted-foreground" />
          <h3 className="mt-4 text-xl font-semibold">No Tickets Found</h3>
          <p className="mt-1 text-muted-foreground">
            You haven't purchased any tickets yet.
          </p>
          <div className="mt-6">
            <Button>
              <Link to="/events">Explore Events</Link>
            </Button>
          </div>
        </div>
      )}
    </main>
  );
}
