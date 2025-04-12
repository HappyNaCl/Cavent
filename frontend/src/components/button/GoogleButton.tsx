import { Button } from "../ui/button";
import GoogleIcon from "@/assets/GoogleLogo.png";
import { env } from "@/lib/schema/EnvSchema";

export default function GoogleButton() {
  const handleClick = () => {
    window.location.href = `${env.VITE_BACKEND_URL}/api/auth/google`;
  };

  return (
    <Button
      onClick={handleClick}
      className="w-full bg-white text-black hover:bg-gray-200 border-2 border-gray-300/50 py-6 px-8"
    >
      <img src={GoogleIcon} alt="Google" className="w-4 h-4 mr-2" />
      <span className="text-sm font-semibold">Continue with Google</span>
    </Button>
  );
}
