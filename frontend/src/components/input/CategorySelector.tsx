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
import { Category } from "@/interface/Category";

type CategorySelectorProps = {
  onChange: (selectedCategory: Category | null) => void;
};

export default function CategorySelector({ onChange }: CategorySelectorProps) {
  const [selectedId, setSelectedId] = useState<string>("");
  const [categories, setCategories] = useState<CategoryType[]>([]);

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

  useEffect(() => {
    const allCategories = categories.flatMap((type) => type.categories);
    const selected = allCategories.find((cat) => cat.id === selectedId) || null;
    onChange(selected);
  }, [selectedId, categories, onChange]);

  const selectCategory = (categoryId: string) => {
    setSelectedId((prev) => (prev === categoryId ? "" : categoryId));
  };

  return (
    <div className="w-full max-w-md space-y-4">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-semibold">Select A Category</h2>
      </div>
      <ScrollArea className="h-64 rounded-md border p-2">
        <Accordion type="multiple" className="w-full">
          {categories.map((type) => (
            <AccordionItem key={type.id} value={type.id}>
              <AccordionTrigger className="text-lg">
                {type.name}
              </AccordionTrigger>
              <AccordionContent>
                {type.categories.map((category) => (
                  <label
                    key={category.id}
                    className="flex items-center space-x-2 border p-4"
                  >
                    <Checkbox
                      checked={selectedId === category.id}
                      onCheckedChange={() => selectCategory(category.id)}
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
