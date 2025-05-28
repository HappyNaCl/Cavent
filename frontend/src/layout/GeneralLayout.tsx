import Navbar from "@/components/ui/navbar";
import { Outlet } from "react-router";
import { Toaster } from "sonner";

export default function GeneralLayout() {
  return (
    <>
      <Navbar />
      <div className="flex-1">
        <Outlet />
      </div>
      <Toaster />
      {/* <Footer /> */}
    </>
  );
}
