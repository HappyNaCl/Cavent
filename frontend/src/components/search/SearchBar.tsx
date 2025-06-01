import { EventSearch } from "@/interface/EventSearch";
import api from "@/lib/axios";
import axios from "axios";
import { MapPin, Search, X } from "lucide-react";
import { useEffect, useState, useRef } from "react";
import { toast } from "sonner";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import debounce from "@/lib/debounce";
import { createSearchParams, useNavigate, useSearchParams } from "react-router";

export default function SearchBar() {
  const [query, setQuery] = useState("");
  const [eventSearchResults, setEventSearchResults] = useState<EventSearch[]>(
    []
  );
  const [isOpen, setIsOpen] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  const [searchParams] = useSearchParams();

  useEffect(() => {
    const searchQuery = searchParams.get("search_query");
    if (searchQuery) {
      setQuery(searchQuery);
    }
  }, [searchParams]);

  const handleClear = () => {
    setQuery("");
    setEventSearchResults([]);
    setIsOpen(false);
    inputRef.current?.focus();
  };

  const nav = useNavigate();

  useEffect(() => {
    const handleSearch = async () => {
      if (query.trim() === "") {
        setEventSearchResults([]);
        setIsOpen(false);
        return;
      }

      try {
        const res = await api.get("/event/search", {
          params: { query },
        });

        if (res.status === 200) {
          const results = res.data.data;
          console.log(results);
          setEventSearchResults(results);
          setIsOpen(results.length > 0);
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(
            error.response?.data?.error || "An error occurred while searching."
          );
        }
      }
    };

    const debouncedSearch = debounce(handleSearch, 500);

    debouncedSearch();
  }, [query]);

  return (
    <Popover open={isOpen} onOpenChange={setIsOpen}>
      <PopoverTrigger asChild>
        <div
          className="flex items-center w-full mx-auto bg-white rounded-full shadow-md px-4 py-2 border border-gray-300 focus-within:ring-2 focus-within:ring-yellow-400"
          onClick={() => {
            if (eventSearchResults.length > 0) {
              setIsOpen(true);
            }
            inputRef.current?.focus();
          }}
        >
          <Search className="text-gray-500 w-5 h-5" />
          <input
            ref={inputRef}
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                nav({
                  pathname: "/event/search",
                  search: createSearchParams({ query: query }).toString(),
                });
              }
            }}
            placeholder="Search..."
            className="flex-grow px-4 py-2 text-xl text-gray-800 placeholder-gray-400 bg-transparent focus:outline-none"
          />
          {query && (
            <X
              className="text-gray-500 w-5 h-5 cursor-pointer hover:text-yellow-500 transition-colors duration-200"
              onClick={handleClear}
            />
          )}
        </div>
      </PopoverTrigger>

      {eventSearchResults.length > 0 && (
        <PopoverContent
          className="w-[var(--radix-popover-trigger-width)] mt-2"
          onOpenAutoFocus={(e) => e.preventDefault()}
        >
          <div className="space-y-2">
            {eventSearchResults.map((event) => (
              <div
                key={event.id}
                className="p-2 rounded-md hover:bg-yellow-100 transition-colors cursor-pointer"
                onClick={() => {
                  nav(`/event/${event.id}`);
                  setIsOpen(false);
                }}
              >
                <p className="font-semibold text-gray-800">{event.title}</p>
                <p className="text-sm text-gray-500 flex items-center gap-1">
                  {new Date(event.startTime * 1000).toLocaleString()} •
                  <MapPin className="w-4 h-4 inline text-gray-500" />
                  {event.location}
                </p>
              </div>
            ))}
          </div>
        </PopoverContent>
      )}
    </Popover>
  );
}
