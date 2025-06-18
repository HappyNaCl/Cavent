import { useState, useMemo } from "react";
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

type TicketFormProps = {
  tickets: Ticket[];
};

type SelectedQuantities = {
  [ticketId: string]: number;
};

export default function TicketForm({ tickets }: TicketFormProps) {
  const [selectedQuantities, setSelectedQuantities] =
    useState<SelectedQuantities>({});

  const handleQuantityChange = (ticketId: string, amount: number) => {
    setSelectedQuantities((prev) => {
      const currentQuantity = prev[ticketId] || 0;
      const newQuantity = currentQuantity + amount;

      const ticket = tickets.find((t) => t.id === ticketId);
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
        const ticket = tickets.find((t) => t.id === ticketId);
        if (!ticket) return total;
        return total + ticket.price * quantity;
      },
      0
    );
  }, [selectedQuantities, tickets]);

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    console.log("Submitting order:", {
      selectedTickets: selectedQuantities,
      totalPrice: totalPrice.toFixed(2),
    });
    alert(`Order submitted! Total: $${totalPrice.toFixed(2)}`);
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
          {tickets.map((ticket) => {
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
