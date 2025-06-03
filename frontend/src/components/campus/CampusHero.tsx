import { ArrowRight } from "lucide-react";
import { useAuth } from "../provider/AuthProvider";

type CampusHeroProps = {
  onUnauthorized: () => void;
  onNoCampusUserClick: () => void;
  onAlreadyJoinedClick: () => void;
};

export default function CampusHero({
  onUnauthorized,
  onNoCampusUserClick,
  onAlreadyJoinedClick,
}: CampusHeroProps) {
  const { user } = useAuth();

  const handleClick = () => {
    if (!user) {
      onUnauthorized();
    } else if (!user.campusId) {
      onNoCampusUserClick();
    } else {
      onAlreadyJoinedClick();
    }
  };

  return (
    <section className="bg-amber-400 relative overflow-hidden py-16 px-6 md:px-20 w-full">
      <div className="relative z-10 max-w-4xl mx-auto text-center">
        <h1 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
          Explore Top Campuses Around You!
        </h1>
        <p className="text-gray-800 text-lg md:text-xl mb-8">
          Discover vibrant student communities, events, and hidden gems at your
          favorite campuses. Your journey starts here!
        </p>
        <button
          onClick={handleClick}
          className="inline-flex items-center gap-2 bg-gray-900 text-white font-semibold px-6 py-3 rounded-md shadow-md hover:bg-gray-800 transition duration-300"
        >
          Join a Campus
          <ArrowRight />
        </button>
      </div>
    </section>
  );
}
