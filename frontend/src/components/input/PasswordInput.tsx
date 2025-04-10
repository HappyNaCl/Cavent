import { useState } from "react";
import { Button } from "../ui/button";
import { LucideEye, LucideEyeClosed } from "lucide-react";
import { Input } from "../ui/input";

interface PasswordInputProps {
  label?: string;
  placeholder?: string;
  value?: string;
  id?: string;
  className?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

export default function PasswordInput({
  label,
  placeholder,
  value,
  id,
  className,
  onChange,
}: PasswordInputProps) {
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const [isFocused, setIsFocused] = useState(false);

  return (
    <div className="flex flex-col w-full gap-1">
      <label htmlFor={id} className="text-gray-700">
        {label}
      </label>
      <div
        className={`flex flex-row w-full items-center border border-gray-300 rounded-lg transition-all duration-100 ${
          isFocused ? "outline-3 outline-gray-300 " : ""
        }`}
      >
        <Input
          type={isPasswordVisible ? "text" : "password"}
          id={id}
          value={value}
          onChange={onChange}
          onFocus={() => setIsFocused(true)}
          onBlur={() => setIsFocused(false)}
          className={
            className ||
            "rounded-lg px-4 py-6 w-full border-0 focus:outline-none focus:border-0 focus:ring-0 focus-visible:ring-transparent transition-color duration-200 ease-in-out"
          }
          placeholder={placeholder}
        />
        <Button
          type="button"
          className="bg-white text-black hover:bg-white rounded-full aspect-square"
          onClick={() => setIsPasswordVisible(!isPasswordVisible)}
        >
          {isPasswordVisible ? (
            <LucideEye size={40} />
          ) : (
            <LucideEyeClosed size={40} />
          )}
        </Button>
      </div>
    </div>
  );
}
