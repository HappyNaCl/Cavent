import { Input } from "../ui/input";

interface TextInputProps {
  type: string;
  label?: string;
  placeholder?: string;
  value?: string;
  id?: string;
  className?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

export default function TextInput({
  type,
  label,
  placeholder,
  value,
  id,
  className,
  onChange,
}: TextInputProps) {
  return (
    <div className="flex flex-col w-full gap-1">
      <label htmlFor={id} className="text-gray-700">
        {label}
      </label>
      <Input
        type={type}
        id={id}
        value={value}
        onChange={onChange}
        className={className || "border border-gray-300 rounded-lg px-4 py-6"}
        placeholder={placeholder}
      />
    </div>
  );
}
