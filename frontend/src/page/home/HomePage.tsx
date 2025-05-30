import CategoryLinkCarousel from "@/components/categories/CategoryLinkCarousel";
import EventGrid from "@/components/events/EventGrid";
import RecommendedEventGrid from "@/components/events/RecommendedEventGrid";

export default function HomePage() {
  return (
    <>
      <CategoryLinkCarousel />
      <RecommendedEventGrid />
      <EventGrid />
    </>
  );
}
