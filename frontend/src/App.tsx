import { Route, Routes, useNavigate } from "react-router";
import AuthPage from "./page/auth/AuthPage";
import HomePage from "./page/home/HomePage";
import { useEffect, useRef } from "react";
import axios from "axios";
import { env } from "@/lib/schema/EnvSchema";
import { useAuth } from "./components/provider/AuthProvider";

function App() {
  const hasRun = useRef(false);
  const nav = useNavigate();
  const { login } = useAuth();

  useEffect(() => {
    if (hasRun.current) return;
    hasRun.current = true;

    async function checkToken() {
      try {
        const res = await axios.get(`${env.VITE_BACKEND_URL}/api/auth/me`, {
          withCredentials: true,
        });

        console.log(res);
        if (res.status === 200) {
          login(res.data.user);
          nav("/");
        } else {
          nav("/auth");
        }
      } catch {
        nav("/auth");
      }
    }

    checkToken();
  }, [login, nav]);

  return (
    <main className="w-screen h-fit min-h-screen">
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/auth" element={<AuthPage />} />
      </Routes>
    </main>
  );
}

export default App;
