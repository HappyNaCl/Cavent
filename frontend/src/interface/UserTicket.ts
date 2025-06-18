export interface Ticket {
  id: string;
  name: string;
}

export interface UserTicket {
  eventId: string;
  eventTitle: string;
  startTime: string;
  endTime: string;
  tickets: Ticket[];
}
