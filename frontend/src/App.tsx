import { Route, Routes } from "react-router";
import AuthPage from "./page/auth/AuthPage";

function App() {
  return (
    <main className="w-screen h-fit min-h-screen">
      <Routes>
        <Route path="/auth" element={<AuthPage />} />
      </Routes>
    </main>
  );
}

export default App;
