import { useState } from "react";
import { Checkbox } from "../ui/checkbox";
import { Label } from "../ui/label";

export interface Option {
  id: string;
  label: string;
}

type FilterGroup = {
  options: Option[];
  value?: string | null;
  onChange: (value: string | null) => void;
  allowDeselect?: boolean;
};

export default function FilterGroup({
  options,
  value,
  onChange,
  allowDeselect = true,
}: FilterGroup) {
  const [internalValue, setInternalValue] = useState<string | null>(
    value ?? null
  );

  const selectedValue = value !== undefined ? value : internalValue;
  const setValue = (val: string | null) => {
    onChange?.(val);
    if (value === undefined) {
      setInternalValue(val);
    }
  };

  return (
    <div className="flex flex-col gap-4">
      {options.map((option) => (
        <div key={option.id} className="flex items-center space-x-4">
          <Checkbox
            id={option.id}
            checked={selectedValue === option.id}
            onCheckedChange={() => {
              if (allowDeselect && selectedValue === option.id) {
                setValue(null);
              } else {
                setValue(option.id);
              }
            }}
          />
          <Label
            htmlFor={option.id}
            className="text-md font-normal leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
          >
            {option.label}
          </Label>
        </div>
      ))}
    </div>
  );
}
