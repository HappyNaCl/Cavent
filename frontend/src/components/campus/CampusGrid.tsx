import { Campus } from "@/interface/Campus";
import { useEffect, useState } from "react";
import CampusCard from "../cards/CampusCard";
import axios from "axios";
import api from "@/lib/axios";

export default function CampusGrid() {
  const [campuses, setCampuses] = useState<Campus[]>([]);

  useEffect(() => {
    async function fetchCampuses() {
      try {
        const res = await api.get("/campus");
        if (res.status === 200) {
          setCampuses(res.data.data);
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          const errorMessage =
            error.response?.data?.error || "An error occurred";
          console.error(`Error: ${errorMessage}`);
        }
      }
    }

    fetchCampuses();
  }, []);

  return (
    <section className="flex flex-col w-full items-center gap-6">
      <span className="text-4xl font-semibold self-start py-2">
        Available Campus
      </span>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 gap-y-8 py-4">
        {campuses.map((campus) => (
          <CampusCard key={campus.id} campus={campus} />
        ))}
      </div>
    </section>
  );
}
