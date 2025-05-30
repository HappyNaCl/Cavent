// ... other imports
import {
  Popover,
  PopoverTrigger,
  PopoverContent,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { Command, CommandItem, CommandList } from "@/components/ui/command";
import { Link } from "react-router";
import {
  BookHeart,
  ChevronDown,
  Settings,
  Star,
  Ticket,
  User,
} from "lucide-react";
import Navlink from "./navlink";
import { useAuth } from "../provider/AuthProvider";
import Logo from "@/assets/Logo.png";
import LogoutButton from "../button/LogoutButton";
import { PopoverArrow } from "@radix-ui/react-popover";
import { useState } from "react";

export default function Navbar() {
  const midLinks = [
    { name: "Home", to: "/" },
    { name: "Events", to: "/event" },
    { name: "Campus", to: "/campus" },
    { name: "About", to: "/about" },
  ];

  const [popoverOpen, setPopoverOpen] = useState(false);

  const { user } = useAuth();

  return (
    <div className="sticky top-0 z-50 flex justify-around h-24 px-16 bg-slate-800 items-center">
      <img src={Logo} alt="Logo" className="h-[50%] object-contain" />

      <nav className="flex flex-row gap-8 h-full text-white list-none text-2xl justify-center items-center">
        {midLinks.map((link) => (
          <Navlink key={link.to} to={link.to} name={link.name} />
        ))}
      </nav>

      <nav className="relative flex flex-row gap-8 text-xl text-white list-none items-center justify-center">
        {user ? (
          <>
            <Link
              to="/event/create"
              className="hover:text-yellow-400 transition"
            >
              Create Event
            </Link>
            <Link
              to="/tickets"
              className="flex flex-col items-center gap-2 hover:text-yellow-400 transition"
            >
              <Ticket className="w-4 h-4" />
              Tickets
            </Link>
            <Link
              to="/tickets"
              className="flex flex-col items-center gap-2 hover:text-yellow-400 transition"
            >
              <Star className="w-4 h-4" />
              Favorites
            </Link>

            <Popover open={popoverOpen} onOpenChange={setPopoverOpen}>
              <PopoverTrigger>
                <Button
                  variant="ghost"
                  className="flex flex-row items-center gap-2 h-fit text-white font-normal hover:text-yellow-600"
                >
                  <div className="flex flex-col items-center gap-2 text-xl">
                    <User className="w-4 h-4" />
                    Profile
                  </div>
                  <ChevronDown
                    className={`w-4 h-4 transition-transform duration-200 ${
                      popoverOpen ? "rotate-180" : "rotate-0"
                    }`}
                  />
                </Button>
              </PopoverTrigger>
              <PopoverContent
                align="center"
                side="bottom"
                className="w-48 p-0 z-[51] sm:z-[60]"
              >
                <PopoverArrow className="fill-white" />
                <Command>
                  <CommandList>
                    <CommandItem>
                      <Link
                        to="/favorites"
                        className="flex items-center gap-2 w-full py-2"
                      >
                        <BookHeart className="w-4 h-4" />
                        Interests
                      </Link>
                    </CommandItem>
                    <CommandItem>
                      <Link
                        to="/settings"
                        className="flex items-center gap-2 w-full py-2"
                      >
                        <Settings className="w-4 h-4" />
                        Settings
                      </Link>
                    </CommandItem>
                    <CommandItem>
                      <LogoutButton />
                    </CommandItem>
                  </CommandList>
                </Command>
              </PopoverContent>
            </Popover>
          </>
        ) : (
          <>
            <Link
              to={user ? "/event/create" : "/auth"}
              className="hover:text-yellow-400 transition text-2xl"
            >
              Create Event
            </Link>
            <Link
              to="/auth"
              state={{ login: true }}
              className="hover:text-yellow-400 transition text-2xl"
            >
              Login
            </Link>
            <Link
              to="/auth"
              state={{ login: false }}
              className="bg-yellow-400 transition text-2xl py-2 px-4 rounded-lg hover:bg-yellow-300 text-slate-800 font-semibold"
            >
              Sign Up
            </Link>
          </>
        )}
      </nav>
    </div>
  );
}
