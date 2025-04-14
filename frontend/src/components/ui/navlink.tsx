import { Link, useLocation } from "react-router";

interface NavlinkProps {
  to: string;
  name: string;
}

export default function Navlink({ to, name }: NavlinkProps) {
  const location = useLocation();
  const isActive = location.pathname === to;

  return (
    <div className="h-full w-fit relative group flex items-center">
      <Link
        to={to}
        className={`relative px-2 py-1 text-white ${
          isActive ? "text-yellow-400" : ""
        }`}
      >
        {name}
      </Link>
      <div
        className={`absolute bottom-0 left-0 w-full h-[6px] bg-yellow-400 transition-all duration-300 ease-in-out ${
          isActive || "opacity-0 group-hover:opacity-100"
        }`}
      ></div>
    </div>
  );
}
