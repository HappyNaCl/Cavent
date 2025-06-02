import EventCard from "@/components/cards/EventCard";
import LoginModal, { LoginModalRef } from "@/components/dialog/LoginDialog";
import { CategoryFilter } from "@/components/search/CategoryFilter";
import FilterGroup, { Option } from "@/components/search/FilterGroup";
import EventCardSkeleton from "@/components/skeleton/EventCardSkeleton";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Separator } from "@/components/ui/separator";
import { BriefEvent } from "@/interface/BriefEvent";
import api from "@/lib/axios";
import debounce from "@/lib/debounce";
import axios from "axios";
import { useEffect, useRef, useState } from "react";
import { useSearchParams } from "react-router";
import { toast } from "sonner";

export default function SearchResultPage() {
  const [searchedEvents, setSearchedEvents] = useState<BriefEvent[]>([]);
  const [selectedPrice, setSelectedPrice] = useState<string | null>(null);
  const [selectedDate, setSelectedDate] = useState<string | null>(null);
  const [selectedCategoryIds, setSelectedCategoryIds] = useState<
    string[] | null
  >(null);

  const priceOptions: Option[] = [
    { id: "free", label: "Free" },
    { id: "paid", label: "Paid" },
  ];

  const dateOptions: Option[] = [
    { id: "today", label: "Today" },
    { id: "tomorrow", label: "Tomorrow" },
    { id: "this-week", label: "This Week" },
    { id: "this-month", label: "This Month" },
    { id: "this-year", label: "This Year" },
  ];

  const [page, setPage] = useState<number>(1);
  const [loading, setLoading] = useState<boolean>(false);
  const [totalPage, setTotalPage] = useState<number>(1);
  const [searchQuery, setSearchQuery] = useState<string>("");
  const [searchParams, setSearchParams] = useSearchParams();

  const limit = 20;
  const requestCounterRef = useRef(0);

  useEffect(() => {
    const query = searchParams.get("query");
    if (query === null || query === "") {
      if (searchParams.has("query")) {
        const newParams = new URLSearchParams(searchParams);
        newParams.delete("query");
        setSearchParams(newParams, { replace: true });
      }
      setSearchQuery("");
    } else {
      setSearchQuery(query);
    }
  }, [searchParams, setSearchParams]);

  useEffect(() => {
    setPage(1);
    setSearchedEvents([]);
    setTotalPage(1);
  }, [searchQuery, selectedPrice, selectedDate, selectedCategoryIds]);

  useEffect(() => {
    requestCounterRef.current += 1;
    const currentRequestId = requestCounterRef.current;

    const paramBuilder = () => {
      const rawParams: Record<string, string | number> = {
        limit: limit,
        page: page,
        query: searchQuery,
        price: selectedPrice || "",
        date: selectedDate || "",
        categories: selectedCategoryIds ? selectedCategoryIds.join(",") : "",
      };

      const params: Record<string, string> = {};
      for (const [key, value] of Object.entries(rawParams)) {
        if (value !== "" && value !== null && value !== undefined) {
          params[key] = String(value);
        }
      }
      return params;
    };

    const fetchEvents = async () => {
      if (currentRequestId !== requestCounterRef.current) {
        console.log(
          `Skipping outdated fetch request (ID: ${currentRequestId}, Current ID: ${requestCounterRef.current})`
        );
        return;
      }

      if (page > 1 && page > totalPage) {
        setLoading(false);
        return;
      }

      setLoading(true);
      try {
        const res = await api.get("/event/all", {
          params: paramBuilder(),
        });

        if (res.status === 200) {
          const newEvents = res.data.data.rows;
          const newTotalPages = res.data.data.totalPages;

          if (currentRequestId !== requestCounterRef.current) {
            console.log(
              `Fetch completed for outdated request (ID: ${currentRequestId}, Current ID: ${requestCounterRef.current}). Discarding results.`
            );
            setLoading(false);
            return;
          }

          if (page === 1) {
            setSearchedEvents(newEvents);
          } else {
            setSearchedEvents((prevEvents) => [...prevEvents, ...newEvents]);
          }
          setTotalPage(newTotalPages > 0 ? newTotalPages : 1);
        }
      } catch (error) {
        if (currentRequestId === requestCounterRef.current) {
          if (axios.isAxiosError(error)) {
            toast.error(
              error.response?.data?.error ||
                "An error occurred while fetching events."
            );
          } else {
            toast.error("An unexpected error occurred.");
          }
          if (page === 1) {
            setSearchedEvents([]);
            setTotalPage(1);
          }
        } else {
          console.log(
            `Error from outdated request (ID: ${currentRequestId}) ignored.`
          );
        }
      } finally {
        if (currentRequestId === requestCounterRef.current) {
          setLoading(false);
        }
      }
    };

    const debouncedFetch = debounce(fetchEvents, 500);
    debouncedFetch();
  }, [
    page,
    totalPage,
    searchQuery,
    selectedPrice,
    selectedDate,
    selectedCategoryIds /* removed totalPage */,
  ]);

  useEffect(() => {
    const handleScroll = () => {
      const bottom =
        Math.ceil(window.innerHeight + window.scrollY) >=
        document.documentElement.scrollHeight - 200;

      if (bottom && !loading && page < totalPage) {
        setPage((prevPage) => prevPage + 1);
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, [loading, page, totalPage]); // totalPage is correctly a dependency here

  const handleUnauthorized = () => {
    if (loginModalRef.current) {
      loginModalRef.current.open();
    }
  };

  const loginModalRef = useRef<LoginModalRef>(null);

  return (
    <div className="w-full">
      <div className="container mx-auto flex flex-col md:flex-row gap-6 md:gap-8">
        <aside className="w-full md:w-[250px] lg:w-[280px] flex-shrink-0 space-y-12 border-r border-gray-300 pr-4 md:pr-8">
          <h1 className="text-2xl font-semibold">Filters</h1>
          <div className="flex flex-col gap-4">
            <h1 className="text-xl font-semibold">Price</h1>
            <FilterGroup
              options={priceOptions}
              value={selectedPrice}
              onChange={(value) => setSelectedPrice(value as string | null)}
              allowDeselect={true}
            />
          </div>
          <Separator className="my-4 border" />
          <div className="flex flex-col gap-4">
            <h1 className="text-xl font-semibold">Date</h1>
            <FilterGroup
              options={dateOptions}
              value={selectedDate}
              onChange={(value) => setSelectedDate(value as string | null)}
              allowDeselect={true}
            />
          </div>
          <Separator className="my-4 border" />
          <div className="flex flex-col gap-4">
            <h1 className="text-xl font-semibold">Category</h1>
            <CategoryFilter
              values={selectedCategoryIds}
              onChange={(values) =>
                setSelectedCategoryIds(values as string[] | null)
              }
            />
          </div>
        </aside>
        <main className="flex-1">
          <div className="flex-1">
            <div className="flex justify-end items-center mb-6">
              <label htmlFor="sort-by" className="mr-2 text-sm text-gray-700">
                Sort by:
              </label>
              <Select defaultValue="relevance">
                <SelectTrigger id="sort-by" className="w-[180px] bg-white">
                  <SelectValue placeholder="Relevance" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="relevance">Relevance</SelectItem>
                  <SelectItem value="date">Date</SelectItem>
                  <SelectItem value="price-asc">Price: Low to High</SelectItem>
                  <SelectItem value="price-desc">Price: High to Low</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
              {searchedEvents.map((event) => (
                <EventCard
                  key={event.id}
                  event={event}
                  onUnauthorized={handleUnauthorized}
                />
              ))}

              {loading &&
                Array.from({ length: 6 }).map((_, index) => (
                  <EventCardSkeleton key={`skeleton-${index}`} />
                ))}
            </div>
            {!loading && searchedEvents.length === 0 && (
              <div className="text-center py-10 text-gray-500">
                No events found matching your criteria.
              </div>
            )}
          </div>
        </main>
      </div>
      <LoginModal ref={loginModalRef} />
    </div>
  );
}
