import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Facebook, Linkedin, Copy, X, MessageCircleMore } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useState } from "react";
import { env } from "@/lib/schema/EnvSchema";

type ShareDialogProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  eventId: string;
};

export function ShareModal({ open, onOpenChange, eventId }: ShareDialogProps) {
  const [copied, setCopied] = useState(false);
  const url = `${env.VITE_APP_URL}/event/${eventId}`;

  const handleCopy = () => {
    navigator.clipboard.writeText(url);
    setCopied(true);
    setTimeout(() => setCopied(false), 1500);
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-sm">
        <DialogHeader>
          <DialogTitle className="text-center">Share with friends</DialogTitle>
        </DialogHeader>

        <div className="flex justify-center items-center gap-4 mt-2">
          <Button size="icon" variant="ghost">
            <Facebook className="text-blue-600 w-5 h-5" />
          </Button>
          <Button size="icon" variant="ghost">
            <X className="w-5 h-5" />
          </Button>
          <Button size="icon" variant="ghost">
            <MessageCircleMore className="text-green-500 w-5 h-5" />
          </Button>
          <Button size="icon" variant="ghost">
            <Linkedin className="text-blue-700 w-5 h-5" />
          </Button>
        </div>

        <div className="mt-4">
          <label className="block text-sm font-medium mb-1">Event URL</label>
          <div className="relative">
            <Input readOnly value={url} className="pr-10" />
            <Button
              size="icon"
              variant="ghost"
              onClick={handleCopy}
              className="absolute right-1 top-1/2 -translate-y-1/2"
            >
              <Copy className="w-4 h-4" />
            </Button>
          </div>
          {copied && <p className="text-green-500 text-xs mt-1">Copied!</p>}
        </div>
      </DialogContent>
    </Dialog>
  );
}
