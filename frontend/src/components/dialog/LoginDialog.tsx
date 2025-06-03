import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from "@/components/ui/dialog";
import { Link } from "react-router";
import { forwardRef, useImperativeHandle, useState } from "react";
import GoogleButton from "../button/GoogleButton";

export interface LoginModalRef {
  open: () => void;
  close: () => void;
}

const LoginModal = forwardRef<LoginModalRef>((_, ref) => {
  const [open, setOpen] = useState(false);

  useImperativeHandle(ref, () => ({
    open: () => setOpen(true),
    close: () => setOpen(false),
  }));

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle className="text-xl">Login</DialogTitle>
          <DialogDescription className="text-lg">
            Please log in to continue.
          </DialogDescription>
        </DialogHeader>

        <div className="mt-4 flex flex-col gap-4">
          <div className="flex flex-col justify-between gap-4">
            <GoogleButton />
            <Link
              className="w-full bg-white text-black hover:bg-gray-200 border-2 text-center border-gray-300/50 py-3 px-8"
              to={"/auth"}
              state={{ login: true }}
            >
              Login
            </Link>
            <Link
              to="/forgot-password"
              className="text-sm text-blue-600 underline hover:text-blue-800"
            >
              Forgot your password?
            </Link>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
});

export default LoginModal;
