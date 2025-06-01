import { BriefCampus } from "./BriefCampus";
import { Ticket } from "./Ticket";

export interface Event {
  id: string;
  title: string;
  createdById: string;
  campusId: string;
  eventType: string;
  ticketType: string;
  startTime: number;
  endTime: number;
  location: string;
  description: string;
  bannerUrl: string;
  isFavorited: boolean;

  campus: BriefCampus;
  tickets: Ticket[];
}
