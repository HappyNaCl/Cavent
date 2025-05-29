import { LucideIcon } from "lucide-react";
import { Link } from "react-router";

type CategoryLinkProps = {
  name: string;
  path: string;
  color: string;
  icon: LucideIcon;
};

export default function CategoryLink({
  name,
  path,
  icon: Icon,
  color,
}: CategoryLinkProps) {
  return (
    <Link to={path} className="flex flex-col items-center gap-2 group">
      <div
        className={`p-8 rounded-full border-2 flex items-center justify-center`}
        style={{ borderColor: color }}
      >
        <Icon
          size={44}
          className="text-gray-700 group-hover:scale-110 transition-transform"
        />
      </div>
      <span className="text-sm text-center text-gray-800">{name}</span>
    </Link>
  );
}
