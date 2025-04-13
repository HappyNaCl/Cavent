import GoogleButton from "@/components/button/GoogleButton";
import Logo from "@/assets/Logo.png";
import Divider from "@/components/ui/divider";
import { useState } from "react";
import LoginForm from "@/components/form/LoginForm";
import RegisterForm from "@/components/form/RegisterForm";
import { Button } from "@/components/ui/button";
import { Toaster } from "@/components/ui/sonner";

export default function AuthPage() {
  const [isLogin, setIsLogin] = useState(true);

  const toggleForm = () => {
    setIsLogin(!isLogin);
  };

  return (
    <div className="flex justify-around min-h-screen h-fit bg-gray-700">
      <div className="flex flex-col w-5/12 p-4 shadow-lg px-12 py-8">
        <div className="h-1/12 flex">
          <img src={Logo} alt="" />
        </div>
        <div className="flex-1 mt-16">
          <span className="text-6xl text-white font-semibold">
            Discover tailored events.
            <br />
            Sign up for personalized recommendations today!
          </span>
        </div>
      </div>
      <div className="flex flex-col flex-1 rounded-4xl bg-white shadow-lg px-36 justify-center sm:gap-8 lg:gap-4">
        <span className="font-semibold text-5xl">
          {isLogin ? "Login" : "Create Account"}
        </span>
        <div className="w-full">
          <GoogleButton />
        </div>
        <Divider label="OR" />
        {isLogin ? <LoginForm /> : <RegisterForm />}
        {isLogin ? (
          <div className="flex items-center">
            Doesn't have an account?{" "}
            <Button
              onClick={toggleForm}
              className="bg-white shadow-none text-black font-normal hover:bg-white"
            >
              Register here!
            </Button>
          </div>
        ) : (
          <div className="flex items-center">
            Already have an account?{" "}
            <Button
              onClick={toggleForm}
              className="bg-white shadow-none text-black font-normal hover:bg-white"
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
