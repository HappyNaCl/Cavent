import User from "@/interface/User";
import api from "@/lib/axios";
import { InternalAxiosRequestConfig } from "axios";
import { createContext, useContext, useEffect, useLayoutEffect, useState } from "react";

interface AuthContextType {
  user: User | null;
  token: string | null;
  login: (user: User | null, token: string | null) => void;
  logout: () => void;
}

interface AuthInterceptorConfig extends InternalAxiosRequestConfig {
  retry?: boolean;
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
  const [token, setToken] = useState<string | null>(null)
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const fetchMe = async () => {
      try {
        const response = await api.get("/auth/refresh");
        setUser(response.data.data) 
        setToken(response.data.accessToken);
      }
      catch (error) {
        setUser(null);
        setToken(null);
        console.error("Failed to fetch user data:", error);
      }
    }

    fetchMe();
  }, [])

  useLayoutEffect(() => {
    const authInterceptor = api.interceptors.request.use((config: AuthInterceptorConfig) => {
      config.headers.Authorization = 
      !config.retry && token 
        ? `Bearer ${token}` 
        : config.headers.Authorization;

      return config;
    })

    return () => {
      api.interceptors.request.eject(authInterceptor);
    }
  }, [token]);

  useLayoutEffect(() => {
    const refreshInterceptor = api.interceptors.response.use(
      (response) => response,
      async (error) => {
        const originalRequest = error.config;

        if (error.response && error.response.status === 401 && error.response.data.message === "unauthorized") {
          try {
            const response = await api.get("/auth/refresh");
            setUser(response.data.data);
            setToken(response.data.accessToken);

            originalRequest.headers.Authorization = `Bearer ${response.data.accessToken}`;
            originalRequest.retry = true;

            return api(originalRequest);
          } catch {
            setToken(null);
            setUser(null);
          }
        }
        return Promise.reject(error);
      }
    );

    return () => {
      api.interceptors.response.eject(refreshInterceptor);
    };
  }, []);

  return (
    <AuthContext.Provider
      value={{
        user,
        token,
        login: (user, token) => {
          setUser(user)
          setToken(token);
        },
        logout: () => {
          setUser(null);
          setToken(null);
        },
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
