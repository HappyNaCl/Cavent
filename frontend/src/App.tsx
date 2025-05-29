import { Route, Routes } from "react-router";
import AuthPage from "./page/auth/AuthPage";
import HomePage from "./page/home/HomePage";
import PreferencePage from "./page/profile/PreferencePage";
import GeneralLayout from "./layout/GeneralLayout";
import CreateEventPage from "./page/event/CreateEventPage";

function App() {
  return (
    <main className="w-full min-h-screen">
      <Routes>
        <Route path="/" element={<GeneralLayout />}>
          <Route index element={<HomePage />} />
          <Route path="profile">
            <Route path="interest" element={<PreferencePage />} />
          </Route>
          <Route path="event">
            <Route path="create" element={<CreateEventPage />} />
          </Route>
        </Route>
        <Route path="/auth" element={<AuthPage />} />
      </Routes>
    </main>
  );
}

export default App;
