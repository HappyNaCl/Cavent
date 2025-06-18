import { Ticket } from "@/interface/Ticket";
import { Dialog, DialogContent } from "../ui/dialog";
import TicketForm from "../form/TicketForm";

type TicketDialogProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  tickets: Ticket[];
};

export function TicketDialog({
  open,
  onOpenChange,
  tickets,
}: TicketDialogProps) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-sm">
        <TicketForm tickets={tickets} />
      </DialogContent>
    </Dialog>
  );
}
