import CategoryLinkCarousel from "@/components/categories/CategoryLinkCarousel";
import LoginModal, { LoginModalRef } from "@/components/dialog/LoginDialog";
import EventGrid from "@/components/events/EventGrid";
import RecommendedEventGrid from "@/components/events/RecommendedEventGrid";
import { useRef } from "react";

export default function HomePage() {
  const loginModalRef = useRef<LoginModalRef>(null);

  const handleUnauthorized = () => {
    if (loginModalRef.current) {
      loginModalRef.current.open();
    }
  };

  return (
    <>
      <CategoryLinkCarousel />
      <RecommendedEventGrid onUnauthorized={handleUnauthorized} />
      <EventGrid onUnauthorized={handleUnauthorized} />

      <LoginModal ref={loginModalRef} />
    </>
  );
}
