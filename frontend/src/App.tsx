import { Route, Routes } from "react-router";
import AuthPage from "./page/auth/AuthPage";
import HomePage from "./page/home/HomePage";
import PreferencePage from "./page/profile/PreferencePage";
import GeneralLayout from "./layout/GeneralLayout";
import CreateEventPage from "./page/event/CreateEventPage";
import HomeLayout from "./layout/HomeLayout";
import EventDetailPage from "./page/event/EventDetailPage";
import SearchResultPage from "./page/search/SearchResultPage";
import CampusPage from "./page/campus/CampusPage";
import FavoritePage from "./page/favorite/FavoritePage";

function App() {
  return (
    <main className="w-full min-h-screen">
      <Routes>
        <Route path="/" element={<HomeLayout />}>
          <Route index element={<HomePage />}></Route>
          <Route path="event" element={<SearchResultPage />} />
          <Route path="campus" element={<CampusPage />} />
        </Route>
        <Route path="favorite" element={<GeneralLayout />}>
          <Route index element={<FavoritePage />} />
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
