import User from "@/interface/User";
import { createContext, useContext, useState } from "react";

interface AuthContextType {
  user: User | null;
  login: (user: User | null) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);

  return (
    <AuthContext.Provider
      value={{
        user,
        login: (user) => setUser(user),
        logout: () => setUser(null),
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
