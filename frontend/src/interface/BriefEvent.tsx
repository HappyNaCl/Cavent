export interface BriefEvent {
  id: string;
  title: string;
  startDate: number;
  location: string;
  ticketType: "ticketed" | "free";
  campusName: string;
  ticketPrice?: number;
  bannerUrl: string;
  categoryName: string;
}
