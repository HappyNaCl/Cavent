import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { forwardRef, useImperativeHandle, useState } from "react";
import JoinCampusForm from "../form/JoinCampusForm";

export interface JoinCampusRef {
  open: () => void;
  close: () => void;
}

const JoinCampusDialog = forwardRef<JoinCampusRef>((_, ref) => {
  const [open, setOpen] = useState(false);

  useImperativeHandle(ref, () => ({
    open: () => setOpen(true),
    close: () => setOpen(false),
  }));

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle className="text-xl">Join a Campus</DialogTitle>
          <DialogDescription className="text-lg">
            Please enter the invite code to join a campus community
          </DialogDescription>
        </DialogHeader>

        <div className="mt-4 flex flex-col gap-4">
          <JoinCampusForm onSuccess={() => setOpen(false)} />
        </div>
      </DialogContent>
    </Dialog>
  );
});

export default JoinCampusDialog;
