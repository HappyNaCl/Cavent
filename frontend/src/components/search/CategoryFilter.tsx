import { useState } from "react";
import { Checkbox } from "@/components/ui/checkbox";
import { Option } from "./FilterGroup";
import { Label } from "../ui/label";

interface CategoryFilterProps {
  options: Option[];
  values?: string[];
  onChange?: (values: string[]) => void;
}

export function CategoryFilter({
  options,
  values,
  onChange,
}: CategoryFilterProps) {
  const [internalValues, setInternalValues] = useState<string[]>(values ?? []);

  const selectedValues = values ?? internalValues;

  const toggleValue = (id: string) => {
    const isSelected = selectedValues.includes(id);
    const newValues = isSelected
      ? selectedValues.filter((v) => v !== id)
      : [...selectedValues, id];

    onChange?.(newValues);

    if (values === undefined) {
      setInternalValues(newValues);
    }
  };

  return (
    <div className="flex flex-col gap-2">
      {options.map((option) => (
        <div key={option.id} className="flex items-center space-x-2">
          <Checkbox
            id={option.id}
            checked={selectedValues.includes(option.id)}
            onCheckedChange={() => toggleValue(option.id)}
          />
          <Label
            htmlFor={option.id}
            className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
          >
            {option.label}
          </Label>
        </div>
      ))}
    </div>
  );
}
