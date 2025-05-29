import { Search, X } from "lucide-react";

type SearchBarProps = {
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  onClear?: () => void;
};

export default function SearchBar({
  value,
  onChange,
  onClear,
}: SearchBarProps) {
  return (
    <div className="flex items-center w-full mx-auto bg-white rounded-full shadow-md px-4 py-2 border border-gray-300 focus-within:ring-2 focus-within:ring-yellow-400">
      <Search className="text-gray-500 w-5 h-5" />
      <input
        type="text"
        value={value}
        onChange={onChange}
        placeholder="Search..."
        className="flex-grow px-4 py-2 text-xl text-gray-800 placeholder-gray-400 bg-transparent focus:outline-none"
      />
      {value && (
        <X
          className="text-gray-500 w-5 h-5 cursor-pointer hover:text-yellow-500 transition-colors duration-200"
          onClick={onClear}
        />
      )}
    </div>
  );
}
