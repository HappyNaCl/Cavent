import Navbar from "@/components/ui/navbar";
import { Outlet } from "react-router";
import { Toaster } from "sonner";

export default function GeneralLayout() {
  return (
    <div className="flex flex-col h-screen">
      <Navbar />
      <div className="flex-1 overflow-y-auto">
        <Outlet />
      </div>
      <Toaster />
      {/* <Footer /> */}
    </div>
  );
}
