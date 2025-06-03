import { Campus } from "@/interface/Campus";

interface CampusCardProps {
  campus: Campus;
}

export default function CampusCard({ campus }: CampusCardProps) {
  const imageUrl =
    campus.logoUrl || "https://via.placeholder.com/400x225?text=Campus+Image";

  return (
    <div className="w-full bg-white rounded-lg shadow-md overflow-hidden transform transition duration-300 hover:scale-105 max-h-96">
      <div className="relative">
        <img
          src={imageUrl}
          alt={campus.name}
          className="w-full h-48 object-cover"
        />
      </div>
      <div className="p-4">
        <div className="flex items-start mb-2">
          <h3 className="text-xl font-semibold leading-tight">{campus.name}</h3>
        </div>
        <p className="text-gray-600 text-sm mb-2 text-justify line-clamp-3">
          {campus.description}
        </p>

        <div className="flex items-center justify-between text-gray-700 text-sm">
          <div className="flex items-center gap-2 font-bold text-base">
            Invite Code:{" "}
            <span className="text-yellow-600 font-mono text-xl">
              {campus.inviteCode}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
}
