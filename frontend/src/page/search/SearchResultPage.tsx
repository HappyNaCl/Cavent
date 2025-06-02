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
  const [searchQuery, setSearchQuery] = useState<string>("");
  const [searchParams] = useSearchParams();

  useEffect(() => {
    const searchQuery = searchParams.get("search_query");
    if (searchQuery) {
      setSearchQuery(searchQuery);
    }
  }, [searchParams]);

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
  const [totalPage, setTotalPage] = useState<number>(2);

  const limit = 20;

  useEffect(() => {
    async function fetchFirstEvents() {
      setLoading(true);
      try {
        const res = await api.get("/event/all", {
          params: {
            limit: limit,
            page: 1,
          },
        });
        if (res.status === 200) {
          setSearchedEvents(res.data.data.rows);
          setTotalPage(res.data.data.totalPages);
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(`${error.response?.data.error}` || "An error occured");
        }
      } finally {
        setLoading(false);
      }
    }

    fetchFirstEvents();
  }, []);

  useEffect(() => {
    if (page === 1) return;

    async function fetchEvents() {
      if (page > totalPage) return;
      setLoading(true);
      try {
        const res = await api.get("/event/all", {
          params: {
            limit: limit,
            page: page,
          },
        });

        if (res.status === 200) {
          setSearchedEvents((prevEvents) => [
            ...prevEvents,
            ...res.data.data.rows,
          ]);
          setTotalPage(res.data.data.totalPages);
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(`${error.response?.data.error}` || "An error occured");
        }
      } finally {
        setLoading(false);
      }
    }

    fetchEvents();
  }, [page, totalPage]);

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
  }, [loading, page, totalPage]);

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
              onChange={setSelectedPrice}
              allowDeselect={true}
            />
          </div>
          <Separator className="my-4 border" />
          <div className="flex flex-col gap-4">
            <h1 className="text-xl font-semibold">Date</h1>
            <FilterGroup
              options={dateOptions}
              value={selectedDate}
              onChange={setSelectedDate}
              allowDeselect={true}
            />
          </div>
          <Separator className="my-4 border" />
          <div className="flex flex-col gap-4">
            <h1 className="text-xl font-semibold">Category</h1>
            <CategoryFilter
              values={selectedCategoryIds}
              onChange={setSelectedCategoryIds}
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
                <EventCard event={event} onUnauthorized={handleUnauthorized} />
              ))}

              {loading &&
                [1, 2, 3, 4, 5, 6, 7, 8].map(() => <EventCardSkeleton />)}
            </div>
          </div>
        </main>
      </div>
      <LoginModal ref={loginModalRef} />
    </div>
  );
}
