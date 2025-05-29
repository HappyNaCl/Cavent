import { useState } from "react";
import BackgroundImage from "../../assets/image.png";
import SearchBar from "./SearchBar";
import SearchAnimation from "./SearchAnimation";

export default function SearchSection() {
  const [query, setQuery] = useState("");

  const words = [
    "excites you!",
    "motivates you!",
    "inspires you!",
    "makes you smile!",
  ];

  return (
    <div className="relative flex justify-center items-center w-full h-96 bg-black overflow-hidden">
      <img
        className="absolute inset-0 w-full h-full object-cover blur-sm "
        src={BackgroundImage}
        alt=""
      />
      <div className="absolute inset-0 bg-black opacity-50"></div>

      <div className="relative z-10 text-white font-semibold w-2/3 flex flex-col items-center justify-center gap-4">
        <SearchAnimation
          prefix="Explore a world of events. Find what "
          words={words}
        />
        <div className="w-3/4 mt-4">
          <SearchBar
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onClear={() => setQuery("")}
          />
        </div>
      </div>
    </div>
  );
}
