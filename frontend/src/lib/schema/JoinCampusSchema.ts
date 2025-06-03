import { z } from "zod";

export const JoinCampusSchema = z.object({
  inviteCode: z.string().min(1, "Invite code is required"),
});
