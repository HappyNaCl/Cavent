import Navbar from "@/components/ui/navbar";
import { Link, Outlet, useLocation } from "react-router";

export default function SettingsLayout() {
  const location = useLocation();
  const isAccountPage = location.pathname.includes("account");
  return (
    <>
      <Navbar />
      <div className="bg-gray-50 min-h-screen font-sans">
        <div className="p-4 sm:p-6 lg:p-8 h-full">
          <div className="bg-white shadow-sm rounded-lg overflow-hidden flex flex-col lg:flex-row h-full">
            <aside className="w-full lg:w-80 border-b lg:border-r lg:border-b-0 border-gray-200">
              <div className="p-5 bg-gray-100">
                <h1 className="text-2xl font-bold text-gray-800">
                  Account Settings
                </h1>
              </div>
              <div className="bg-gray-50/50 flex-1">
                <nav>
                  <ul>
                    <li>
                      <Link
                        to="/settings/account"
                        className={`block py-4 px-5 bg-white text-gray-900 font-semibold border-l-4 ${
                          isAccountPage
                            ? "border-slate-700"
                            : "border-transparent"
                        }`}
                      >
                        Account Info
                      </Link>
                    </li>
                    <li>
                      <Link
                        to="/settings/password"
                        className={`block py-4 px-5 bg-white text-gray-900 font-semibold border-l-4 ${
                          !isAccountPage
                            ? "border-slate-700"
                            : "border-transparent"
                        }`}
                      >
                        Password
                      </Link>
                    </li>
                  </ul>
                </nav>
              </div>
            </aside>
            <Outlet />
          </div>
        </div>
      </div>
    </>
  );
}
