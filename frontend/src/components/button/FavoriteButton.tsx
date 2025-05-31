import { Star } from "lucide-react";
import { Button } from "../ui/button";
import { useAuth } from "../provider/AuthProvider";
import { toast } from "sonner";
import axios from "axios";
import { useState, useRef, MouseEvent } from "react";
import api from "@/lib/axios";

type FavoriteButtonProps = {
  isFavorited?: boolean;
  eventId: string;
  onUnauthorized: () => void;
  onSuccess: (newCount: number) => void;
};

export default function FavoriteButton({
  isFavorited: isFavoriteProp,
  eventId,
  onUnauthorized,
  onSuccess,
}: FavoriteButtonProps) {
  const { user } = useAuth();
  const [isFavorite, setIsFavorite] = useState<boolean | undefined>(
    isFavoriteProp
  );
  const lastClickRef = useRef(0);
  const throttleDelay = 1000;

  const handleClick = async (e: MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();

    const now = Date.now();
    if (now - lastClickRef.current < throttleDelay) {
      return;
    }
    lastClickRef.current = now;

    if (!user) {
      onUnauthorized();
      return;
    }

    try {
      if (isFavorite) {
        const response = await api.delete(`/event/${eventId}/favorite`);
        setIsFavorite(false);
        onSuccess(response.data.data);
      } else {
        const response = await api.post(`/event/${eventId}/favorite`);
        setIsFavorite(true);
        onSuccess(response.data.data);
      }
    } catch (error) {
      console.error(error);
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 401) {
          onUnauthorized();
        } else {
          toast.error(error.response?.data?.error || "An error occurred");
        }
      }
    }
  };

  return (
    <Button
      className="bg-white rounded-3xl p-1 hover:bg-gray-100"
      onClick={handleClick}
    >
      <Star
        className={`text-black ${
          isFavorite ? "fill-yellow-400 text-yellow-400" : "fill-none"
        }`}
      />
    </Button>
  );
}
