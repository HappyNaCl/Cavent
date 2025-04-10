interface DividerProps {
  label?: string | "";
}

export default function Divider({ label }: DividerProps) {
  return (
    <div className="flex items-center justify-center my-4">
      <div className="border-t border-gray-300 flex-grow mr-6"></div>
      <span className="text-gray-400 text-2xl">{label}</span>
      <div className="border-t border-gray-300 flex-grow ml-6"></div>
    </div>
  );
}
