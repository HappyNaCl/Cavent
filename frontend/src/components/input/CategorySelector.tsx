import { useEffect, useState } from "react";
import {
  Accordion,
  AccordionItem,
  AccordionTrigger,
  AccordionContent,
} from "@/components/ui/accordion";
import { Checkbox } from "@/components/ui/checkbox";
import { ScrollArea } from "@/components/ui/scroll-area";
import { CategoryType } from "@/interface/CategoryType";
import api from "@/lib/axios";
import { toast } from "sonner";

type Category = CategoryType["categories"][0];

type CategorySelectorProps = {
  onChange: (selectedCategories: Category[]) => void;
};

export default function CategorySelector({ onChange }: CategorySelectorProps) {
  const [selectedIds, setSelectedIds] = useState<string[]>([]);
  const [categories, setCategories] = useState<CategoryType[]>([]);
  const maxSelection = 3;

  useEffect(() => {
    async function fetchCategoryTypes() {
      try {
        const res = await api.get("/category");
        if (res.status === 200 && res.data.data) {
          setCategories(res.data.data);
        }
      } catch (error) {
        toast.error(`Error: ${error}`);
      }
    }

    fetchCategoryTypes();
  }, []);

  // Send selected category objects to parent
  useEffect(() => {
    const allCategories = categories.flatMap((type) => type.categories);
    const selected = allCategories.filter((cat) =>
      selectedIds.includes(cat.id)
    );
    onChange(selected);
  }, [selectedIds, categories, onChange]);

  const toggleCategory = (categoryId: string) => {
    setSelectedIds((prev) =>
      prev.includes(categoryId)
        ? prev.filter((id) => id !== categoryId)
        : prev.length < maxSelection
        ? [...prev, categoryId]
        : prev
    );
  };

  return (
    <div className="w-full max-w-md space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-semibold">Select Categories</h2>
        <span className="text-sm text-muted-foreground">
          {selectedIds.length} / {maxSelection} selected
        </span>
      </div>
      <ScrollArea className="h-64 rounded-md border p-2">
        <Accordion type="multiple" className="w-full">
          {categories.map((type) => (
            <AccordionItem key={type.id} value={type.id}>
              <AccordionTrigger className="text-lg">
                {type.name}
              </AccordionTrigger>
              <AccordionContent className="">
                {type.categories.map((category) => (
                  <label
                    key={category.id}
                    className="flex items-center space-x-2 border p-4 border-b-gray-800"
                  >
                    <Checkbox
                      checked={selectedIds.includes(category.id)}
                      onCheckedChange={() => toggleCategory(category.id)}
                      disabled={
                        !selectedIds.includes(category.id) &&
                        selectedIds.length >= maxSelection
                      }
                    />
                    <span>{category.name}</span>
                  </label>
                ))}
              </AccordionContent>
            </AccordionItem>
          ))}
        </Accordion>
      </ScrollArea>
    </div>
  );
}
