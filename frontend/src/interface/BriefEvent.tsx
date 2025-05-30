export interface BriefEvent {
  id: string;
  title: string;
  startDate: number;
  endDate?: number;
  location: string;
  ticketType: "ticketed" | "free";
  campusName: string;
  ticketPrice?: number;
  bannerUrl: string;
  categoryName: string;
  favoriteCount: number;
  isFavorited?: boolean;
}
