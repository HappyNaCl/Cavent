import FilterGroup, { Option } from "@/components/search/FilterGroup";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { BriefEvent } from "@/interface/BriefEvent";
import { Separator } from "@radix-ui/react-select";
import { useState } from "react";

export default function SearchResultPage() {
  const [searchedEvents, setSearchedEvents] = useState<BriefEvent[]>([]);
  const [selectedPrice, setSelectedPrice] = useState<string | null>(null);
  const [selectedDate, setSelectedDate] = useState<string | null>(null);

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

  return (
    <div className="p-4 md:p-8 w-full">
      <div className="container mx-auto flex flex-col md:flex-row gap-6 md:gap-8">
        <aside className="w-full md:w-[250px] lg:w-[280px] flex-shrink-0 space-y-12 border-r border-gray-200 pr-4 md:pr-8">
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

            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6"></div>
          </div>
        </main>
      </div>
    </div>
  );
}
