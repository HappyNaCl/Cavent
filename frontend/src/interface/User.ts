export default interface User {
  provider: string;
  id: string;
  name: string;
  email: string;
  avatarUrl: string;
  campusId?: string;
  firstTimeLogin: boolean;
  role: string;
}
