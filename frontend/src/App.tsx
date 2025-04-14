import { Route, Routes, useNavigate } from "react-router";
import AuthPage from "./page/auth/AuthPage";
import HomePage from "./page/home/HomePage";
import { useEffect, useRef, useState } from "react";
import axios from "axios";
import { env } from "@/lib/schema/EnvSchema";
import { useAuth } from "./components/provider/AuthProvider";
import PreferencePage from "./page/profile/PreferencePage";
import GeneralLayout from "./layout/GeneralLayout";

function App() {
  const hasRun = useRef(false);
  const nav = useNavigate();
  const { login } = useAuth();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (hasRun.current) return;
    hasRun.current = true;

    async function checkToken() {
      try {
        const res = await axios.get(`${env.VITE_BACKEND_URL}/api/auth/me`, {
          withCredentials: true,
        });

        if (res.status === 200) {
          login(res.data.user);
          if (res.data.user.firstTimeLogin) {
            nav("/profile/preference");
          } else {
            nav("/");
          }
        } else {
          nav("/auth");
        }
      } catch {
        nav("/auth");
      }
      setLoading(false);
    }

    checkToken();
  }, [login, nav]);

  if (loading) {
    return (
      <main className="w-screen h-screen flex items-center justify-center">
        <span className="text-gray-500">Checking authentication...</span>
      </main>
    );
  }

  return (
    <main className="w-screen h-fit min-h-screen">
      <Routes>
        <Route path="/" element={<GeneralLayout />}>
          <Route index element={<HomePage />} />
          <Route path="/profile">
            <Route path="preference" element={<PreferencePage />} />
          </Route>
          {/* <Route path="home" element={<HomePage />} />
          <Route path="preference" element={<PreferencePage />} /> */}
        </Route>
        <Route path="/auth" element={<AuthPage />} />
      </Routes>
    </main>
  );
}

export default App;
