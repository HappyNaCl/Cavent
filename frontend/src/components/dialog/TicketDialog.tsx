import { Ticket } from "@/interface/Ticket";
import { Dialog, DialogContent, DialogHeader } from "../ui/dialog";
import TicketForm from "../form/TicketForm";
import { DialogTitle } from "@radix-ui/react-dialog";

type TicketDialogProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  tickets: Ticket[];
  eventId: string;
};

export function TicketDialog({
  open,
  onOpenChange,
  tickets,
  eventId,
}: TicketDialogProps) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-sm">
        <DialogHeader>
          <DialogTitle className="text-center"></DialogTitle>
        </DialogHeader>
        <TicketForm tickets={tickets} eventId={eventId} />
      </DialogContent>
    </Dialog>
  );
}
