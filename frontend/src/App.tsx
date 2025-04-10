import { Route, Routes } from "react-router";
import AuthLayout from "./layout/AuthLayout";

function App() {
  return (
    <main className="w-screen h-fit min-h-screen bg-gray-700">
      <Routes>
        <Route path="/auth" element={<AuthLayout />} />
      </Routes>
    </main>
  );
}

export default App;
