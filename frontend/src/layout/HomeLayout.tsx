import SearchSection from "@/components/search/SearchSection";
import Navbar from "@/components/ui/navbar";
import { Outlet } from "react-router";
import { Toaster } from "sonner";

export default function HomeLayout() {
  return (
    <>
      <Navbar />
      <SearchSection />
      <div className="flex-1 flex flex-col gap-8 items-center justify-center px-36 py-24">
        <Outlet />
      </div>
      <Toaster />
      {/* <Footer /> */}
    </>
  );
}
