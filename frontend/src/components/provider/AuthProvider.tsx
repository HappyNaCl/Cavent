import User from "@/interface/User";
import api from "@/lib/axios";
import { AxiosError, InternalAxiosRequestConfig, AxiosResponse } from "axios";
import {
  createContext,
  useContext,
  useEffect,
  useLayoutEffect,
  useState,
  ReactNode,
} from "react";
import { useNavigate } from "react-router";

interface RefreshTokenData {
  accessToken: string;
}

interface RefreshTokenApiResponse {
  data: RefreshTokenData;
}

interface UserData {
  user: User;
}

interface MeApiResponse {
  data: UserData;
}

interface AuthContextType {
  user: User | null;
  token: string | null;
  login: (user: User | null, token: string | null) => void;
  logout: () => void;
  setUser: (user: User | null) => void;
  isLoading: boolean;
}

interface AuthInterceptorConfig extends InternalAxiosRequestConfig {
  _retry?: boolean;
  _isRefreshTokenRequest?: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

let refreshTokenPromise: Promise<string | null> | null = null;

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [token, setToken] = useState<string | null>(null);
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  const nav = useNavigate();

  useEffect(() => {
    const fetchSession = async () => {
      try {
        const refreshRes = await api.get<RefreshTokenApiResponse>(
          "/auth/refresh",
          {
            withCredentials: true,
            _isRefreshTokenRequest: true,
          } as AuthInterceptorConfig
        );
        const newToken = refreshRes.data.data.accessToken;
        setToken(newToken);

        const meRes = await api.get<MeApiResponse>("/auth/me", {
          headers: {
            Authorization: `Bearer ${newToken}`,
          },
        });
        setUser(meRes.data.data.user);

        const user = meRes.data.data.user;
        if (user.firstTimeLogin) {
          nav("/profile/interest");
        }
      } catch (err) {
        setUser(null);
        setToken(null);
        console.error("Initial session fetch/refresh failed:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchSession();
  }, [nav]);

  useLayoutEffect(() => {
    const requestInterceptorId = api.interceptors.request.use(
      (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
        const authConfig = config as AuthInterceptorConfig;

        if (
          token &&
          !authConfig.headers.Authorization &&
          !authConfig._isRefreshTokenRequest
        ) {
          authConfig.headers.Authorization = `Bearer ${token}`;
        }
        return authConfig;
      },
      (error: AxiosError) => Promise.reject(error)
    );

    return () => {
      api.interceptors.request.eject(requestInterceptorId);
    };
  }, [token]);

  useLayoutEffect(() => {
    const responseInterceptorId = api.interceptors.response.use(
      (response: AxiosResponse) => response,
      async (error: AxiosError) => {
        const originalRequest = error.config as
          | AuthInterceptorConfig
          | undefined;

        if (
          originalRequest &&
          error.response?.status === 401 &&
          !originalRequest._retry &&
          !originalRequest._isRefreshTokenRequest
        ) {
          originalRequest._retry = true;

          if (!refreshTokenPromise) {
            refreshTokenPromise = api
              .get<RefreshTokenApiResponse>("/auth/refresh", {
                withCredentials: true,
                _isRefreshTokenRequest: true,
              } as AuthInterceptorConfig)
              .then((refreshRes: AxiosResponse<RefreshTokenApiResponse>) => {
                const newToken = refreshRes.data.data.accessToken;
                setToken(newToken);
                return newToken;
              })
              .catch((refreshError: AxiosError) => {
                console.error(
                  "Token refresh failed in interceptor:",
                  refreshError
                );
                setToken(null);
                setUser(null);
                return null;
              })
              .finally(() => {
                refreshTokenPromise = null;
              });
          }

          try {
            const newToken = await refreshTokenPromise;
            if (newToken && originalRequest.headers) {
              originalRequest.headers.Authorization = `Bearer ${newToken}`;
              return api(originalRequest);
            }
          } catch (e) {
            console.error("Error during retry logic after token refresh:", e);
            return Promise.reject(error);
          }
        }
        return Promise.reject(error);
      }
    );

    return () => {
      api.interceptors.response.eject(responseInterceptorId);
    };
  }, []);

  const login = (newUser: User | null, newToken: string | null) => {
    setUser(newUser);
    setToken(newToken);
  };

  const logout = () => {
    setUser(null);
    setToken(null);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center h-screen">
        <div className="w-10 h-10 border-4 border-blue-500 border-t-transparent rounded-full animate-spin" />
      </div>
    );
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        token,
        login,
        logout,
        isLoading: loading,
        setUser: (newUser: User | null) => {
          setUser(newUser);
        },
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
