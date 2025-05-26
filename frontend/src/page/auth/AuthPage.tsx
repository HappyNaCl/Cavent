import GoogleButton from "@/components/button/GoogleButton";
import Logo from "@/assets/Logo.png";
import Divider from "@/components/ui/divider";
import { useState } from "react";
import LoginForm from "@/components/form/LoginForm";
import RegisterForm from "@/components/form/RegisterForm";
import { Button } from "@/components/ui/button";
import { Toaster } from "@/components/ui/sonner";
import { useUnauthGuard } from "@/lib/hook/useUnauthGuard";

export default function AuthPage() {
  useUnauthGuard();

  const [isLogin, setIsLogin] = useState(true);

  const toggleForm = () => {
    setIsLogin(!isLogin);
  };

  return (
    <div className="flex flex-col lg:flex-row min-h-screen h-fit bg-gray-700 lg:justify-around items-center p-4 sm:p-6 lg:p-0">
      <div className="flex flex-col w-full lg:w-5/12 text-white mb-8 lg:mb-0 px-4 py-6 sm:px-6 sm:py-8 lg:px-12 lg:py-10">
        <div className="flex items-center mb-6 md:mb-8 lg:mb-10 xl:mb-12">
          <img
            src={Logo}
            alt="Company Logo"
            className="h-10 sm:h-12 md:h-14 lg:h-16 w-auto"
          />
        </div>
        <div className="flex-1 mt-4 sm:mt-6 md:mt-8 lg:mt-12 xl:mt-16">
          <h1 className="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-semibold font-montserrat leading-tight">
            Discover tailored events.
            <br />
            Sign up for personalized recommendations today!
          </h1>
        </div>
      </div>

      <div className="flex flex-col w-full max-w-sm sm:max-w-md md:max-w-lg lg:max-w-none lg:flex-1 bg-white shadow-2xl rounded-xl lg:rounded-2xl p-6 sm:p-8 md:p-10 lg:px-16 lg:py-12 xl:px-20 xl:py-16 justify-center gap-4 sm:gap-6">
        <h2 className="font-semibold text-3xl sm:text-4xl lg:text-5xl text-center lg:text-left mb-4 sm:mb-6">
          {isLogin ? "Login" : "Create Account"}
        </h2>
        <div className="w-full">
          <GoogleButton />
        </div>
        <Divider label="OR" />
        {isLogin ? <LoginForm /> : <RegisterForm />}
        {isLogin ? (
          <div className="flex flex-col items-center text-center sm:flex-row sm:justify-start sm:text-left text-sm sm:text-base mt-4 gap-1">
            <span>Doesn't have an account?</span>
            <Button
              onClick={toggleForm}
              variant="link"
              className="p-1 text-blue-600 hover:text-blue-800 font-normal"
            >
              Register here!
            </Button>
          </div>
        ) : (
          <div className="flex flex-col items-center text-center sm:flex-row sm:justify-start sm:text-left text-sm sm:text-base mt-4 gap-1">
            <span>Already have an account?</span>
            <Button
              onClick={toggleForm}
              variant="link"
              className="p-1 text-blue-600 hover:text-blue-800 font-normal"
            >
              Login here!
            </Button>
          </div>
        )}
      </div>
      <Toaster position="top-center" />
    </div>
  );
}
