import Logo from "../../assets/Logo.png";
import LogoutButton from "../button/LogoutButton";
import Navlink from "./navlink";

export default function Navbar() {
  const midLinks = [
    { name: "Home", to: "/" },
    { name: "Events", to: "/events" },
    { name: "About", to: "/about" },
  ];

  return (
    <div className="sticky top-0 z-50 flex justify-between h-24 px-24 bg-slate-800 items-center">
      <img src={Logo} alt="Logo" className="h-[50%] object-contain" />
      <nav className="flex flex-row gap-18 h-full text-white list-none text-2xl">
        {midLinks.map((link) => (
          <Navlink key={link.to} to={link.to} name={link.name} />
        ))}
      </nav>
      <nav className="flex flex-row gap-12 text-white list-none text-2xl">
        <li>
          <a href="">Login</a>
        </li>
        <li>
          <LogoutButton />
        </li>
      </nav>
    </div>
  );
}
