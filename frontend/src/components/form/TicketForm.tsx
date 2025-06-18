"use client";

import { useState, useMemo, useEffect } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Minus, Plus } from "lucide-react";
import { Ticket } from "@/interface/Ticket";
import { toast } from "sonner";
import axios from "axios";
import api from "@/lib/axios";

type TicketFormProps = {
  tickets: Ticket[];
  eventId: string;
  onSubmit?: () => void;
};

type SelectedQuantities = {
  [ticketId: string]: number;
};

export default function TicketForm({
  tickets,
  eventId,
  onSubmit,
}: TicketFormProps) {
  const [componentTickets, setComponentTickets] = useState<Ticket[]>(tickets);
  const [selectedQuantities, setSelectedQuantities] =
    useState<SelectedQuantities>({});

  useEffect(() => {
    setComponentTickets(tickets);
  }, [tickets]);

  const handleQuantityChange = (ticketId: string, amount: number) => {
    setSelectedQuantities((prev) => {
      const currentQuantity = prev[ticketId] || 0;
      const newQuantity = currentQuantity + amount;

      const ticket = componentTickets.find((t) => t.id === ticketId);
      if (!ticket) return prev;

      const availableStock = ticket.quantity - ticket.sold;

      if (newQuantity < 0 || newQuantity > availableStock) {
        return prev;
      }

      return {
        ...prev,
        [ticketId]: newQuantity,
      };
    });
  };

  const totalPrice = useMemo(() => {
    return Object.entries(selectedQuantities).reduce(
      (total, [ticketId, quantity]) => {
        const ticket = componentTickets.find((t) => t.id === ticketId);
        if (!ticket) return total;
        return total + ticket.price * quantity;
      },
      0
    );
  }, [selectedQuantities, componentTickets]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const result = Object.entries(selectedQuantities)
      .filter(([, quantity]) => quantity > 0)
      .map(([ticketId, quantity]) => ({
        ticketId,
        quantity,
      }));

    if (result.length === 0) {
      toast.info("Please select at least one ticket to purchase.");
      return;
    }

    const formData = {
      eventId,
      tickets: result,
    };

    try {
      const res = await api.post("/checkout", formData);

      if (res.status === 200) {
        toast.success("Tickets purchased successfully!");

        setComponentTickets((prevTickets) =>
          prevTickets.map((ticket) => {
            const purchasedQuantity = selectedQuantities[ticket.id];
            if (purchasedQuantity > 0) {
              return {
                ...ticket,
                sold: ticket.sold + purchasedQuantity,
              };
            }
            return ticket;
          })
        );

        setSelectedQuantities({});
        onSubmit?.();
      } else {
        toast.error("Failed to purchase tickets");
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        toast.error(error.response?.data?.error || "An error occurred");
      }
    }
  };

  return (
    <Card className="w-full max-w-2xl mx-auto">
      <CardHeader>
        <CardTitle>Buy Tickets</CardTitle>
        <CardDescription>
          Select the number of tickets you would like to purchase.
        </CardDescription>
      </CardHeader>
      <form onSubmit={handleSubmit}>
        <CardContent className="space-y-6">
          {componentTickets.map((ticket) => {
            const availableStock = ticket.quantity - ticket.sold;
            const currentSelection = selectedQuantities[ticket.id] || 0;

            if (availableStock <= 0) {
              return (
                <div
                  key={ticket.id}
                  className="flex items-center justify-between p-4 rounded-lg bg-muted/50 opacity-60"
                >
                  <div>
                    <h3 className="font-semibold">{ticket.name}</h3>
                    <p className="text-sm text-muted-foreground">
                      Price: ${ticket.price.toFixed(2)}
                    </p>
                  </div>
                  <div className="font-bold text-destructive">SOLD OUT</div>
                </div>
              );
            }

            return (
              <div
                key={ticket.id}
                className="flex flex-col sm:flex-row items-start sm:items-center justify-between p-4 border rounded-lg"
              >
                <div className="mb-4 sm:mb-0">
                  <h3 className="text-lg font-semibold">{ticket.name}</h3>
                  <p className="text-sm text-muted-foreground">
                    Price: ${ticket.price.toFixed(2)}
                  </p>
                  <p className="text-xs text-muted-foreground">
                    {availableStock} available
                  </p>
                </div>

                <div className="flex items-center gap-2">
                  <Button
                    type="button"
                    variant="outline"
                    size="icon"
                    className="h-8 w-8"
                    onClick={() => handleQuantityChange(ticket.id, -1)}
                    disabled={currentSelection <= 0}
                  >
                    <Minus className="h-4 w-4" />
                    <span className="sr-only">Decrease quantity</span>
                  </Button>

                  <span className="w-10 text-center text-lg font-medium">
                    {currentSelection}
                  </span>

                  <Button
                    type="button"
                    variant="outline"
                    size="icon"
                    className="h-8 w-8"
                    onClick={() => handleQuantityChange(ticket.id, 1)}
                    disabled={currentSelection >= availableStock}
                  >
                    <Plus className="h-4 w-4" />
                    <span className="sr-only">Increase quantity</span>
                  </Button>
                </div>
              </div>
            );
          })}
        </CardContent>
        <CardFooter className="flex flex-col sm:flex-row items-center justify-between gap-4 bg-muted/50 px-6 py-4">
          <div className="text-xl font-semibold">
            Total:{" "}
            <span className="text-primary">${totalPrice.toFixed(2)}</span>
          </div>
          <Button type="submit" size="lg" disabled={totalPrice <= 0}>
            Proceed to Checkout
          </Button>
        </CardFooter>
      </form>
    </Card>
  );
}
