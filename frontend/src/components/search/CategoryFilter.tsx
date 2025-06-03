import { useEffect, useState } from "react";
import { Checkbox } from "@/components/ui/checkbox";
import { Label } from "../ui/label";
import { Category } from "@/interface/Category";
import { toast } from "sonner";
import axios from "axios";
import api from "@/lib/axios";

interface CategoryFilterProps {
  values: string[] | null;
  onChange?: (values: string[]) => void;
}

export function CategoryFilter({ values, onChange }: CategoryFilterProps) {
  const [internalValues, setInternalValues] = useState<string[]>(values ?? []);
  const [categories, setCategories] = useState<Category[]>([]);

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

  useEffect(() => {
    async function fetchCategory() {
      try {
        const res = await api.get("/category/all");
        if (res.status === 200) {
          setCategories(res.data.data);
        }
      } catch (error) {
        if (axios.isAxiosError(error)) {
          toast.error(
            `${error.response?.data.error}` || "An error has occured"
          );
        }
      }
    }

    fetchCategory();
  }, []);

  return (
    <div className="flex flex-col gap-4">
      {categories.map((option) => (
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
            {option.name}
          </Label>
        </div>
      ))}
    </div>
  );
}
