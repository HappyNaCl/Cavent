"use client";

import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@/components/ui/carousel";
import {
  Music,
  Paintbrush,
  Utensils,
  Dumbbell,
  Briefcase,
  Baby,
  Monitor,
  Clapperboard,
  HeartHandshake,
  HeartPulse,
  Plane,
  GraduationCap,
  Shirt,
} from "lucide-react";
import CategoryLink from "./CategoryLink";

const categories = [
  { name: "Music", color: "#EF4444", icon: Music },
  { name: "Arts & Culture", color: "#8B5CF6", icon: Paintbrush },
  { name: "Food & Drink", color: "#F59E0B", icon: Utensils },
  { name: "Sports & Fitness", color: "#10B981", icon: Dumbbell },
  { name: "Business & Networking", color: "#3B82F6", icon: Briefcase },
  { name: "Family & Kids", color: "#EC4899", icon: Baby },
  { name: "Technology", color: "#6366F1", icon: Monitor },
  { name: "Comedy & Entertainment", color: "#F97316", icon: Clapperboard },
  { name: "Charity & Causes", color: "#84CC16", icon: HeartHandshake },
  { name: "Health & Wellness", color: "#0EA5E9", icon: HeartPulse },
  { name: "Travel & Adventure", color: "#14B8A6", icon: Plane },
  { name: "Education & Learning", color: "#A855F7", icon: GraduationCap },
  { name: "Fashion & Beauty", color: "#F43F5E", icon: Shirt },
];

function slugify(name: string) {
  return name.toLowerCase().replace(/\s+/g, "").replace(/&/g, "-");
}

export default function CategoryLinkCarousel() {
  return (
    <section className="relative w-full space-y-12 py-4">
      <h2 className="text-4xl font-semibold">Explore Categories</h2>
      <Carousel opts={{ align: "start", loop: false }}>
        <CarouselContent className="ml-0 flex">
          {categories.map((cat) => (
            <CarouselItem key={cat.name} className="basis-1/6 flex-shrink-0">
              <CategoryLink
                name={cat.name}
                path={`/category/${slugify(cat.name)}`}
                icon={cat.icon}
                color={cat.color}
              />
            </CarouselItem>
          ))}
        </CarouselContent>
        <CarouselPrevious className="text-2xl -left-6" />
        <CarouselNext className="text-2xl -right-6" />
      </Carousel>
    </section>
  );
}
