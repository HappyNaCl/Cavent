import CategoryLinkCarousel from "@/components/categories/CategoryLinkCarousel";
import EventGrid from "@/components/events/EventGrid";
import RecommendedEventGrid from "@/components/events/RecommendedEventGrid";
import { useAuthGuard } from "@/lib/hook/useAuthGuard";

export default function HomePage() {
  useAuthGuard();

  return (
    <>
      <CategoryLinkCarousel />
      <EventGrid />
      <RecommendedEventGrid />
    </>
  );
}
