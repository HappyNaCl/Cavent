"use client";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { forwardRef, useImperativeHandle, useState } from "react";

export interface AlreadyJoinedDialogRef {
  open: () => void;
  close: () => void;
}

export const AlreadyJoinedDialog = forwardRef<AlreadyJoinedDialogRef>(
  (_, ref) => {
    const [open, setOpen] = useState(false);

    useImperativeHandle(ref, () => ({
      open: () => setOpen(true),
      close: () => setOpen(false),
    }));

    return (
      <Dialog open={open} onOpenChange={setOpen}>
        <DialogContent className="sm:max-w-md">
          <DialogHeader>
            <DialogTitle className="text-center">
              You're already part of a campus!
            </DialogTitle>
            <DialogDescription className="text-center">
              You've already joined a campus community. Check out other events
              or visit the dashboard.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter className="flex justify-center">
            <Button onClick={() => setOpen(false)}>OK</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    );
  }
);

AlreadyJoinedDialog.displayName = "AlreadyJoinedDialog";
