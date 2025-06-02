import { Route, Routes } from "react-router";
import AuthPage from "./page/auth/AuthPage";
import HomePage from "./page/home/HomePage";
import PreferencePage from "./page/profile/PreferencePage";
import GeneralLayout from "./layout/GeneralLayout";
import CreateEventPage from "./page/event/CreateEventPage";
import HomeLayout from "./layout/HomeLayout";
import EventDetailPage from "./page/event/EventDetailPage";
import SearchResultPage from "./page/search/SearchResultPage";

function App() {
  return (
    <main className="w-full min-h-screen">
      <Routes>
        <Route path="/" element={<HomeLayout />}>
          <Route index element={<HomePage />}></Route>
          <Route path="event" element={<SearchResultPage />} />
        </Route>
        <Route path="profile" element={<GeneralLayout />}>
          <Route path="interest" element={<PreferencePage />} />
        </Route>
        <Route path="event" element={<GeneralLayout />}>
          <Route path="create" element={<CreateEventPage />} />
          <Route path=":id" element={<EventDetailPage />}></Route>
        </Route>
        <Route path="/auth" element={<AuthPage />} />
      </Routes>
    </main>
  );
}

export default App;
