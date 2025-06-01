import { Link } from "react-router";

export default function NotFound() {
  return (
    <div className="relative min-h-screen bg-yellow-300 text-gray-900 flex items-center justify-center overflow-hidden">
      <div className="absolute inset-0 opacity-10 pointer-events-none">
        <img
          src="/wavy-lines.svg"
          alt="Wavy lines background"
          className="w-full h-full object-cover"
        />
      </div>

      <div className="relative z-10 text-center max-w-xl px-6">
        <h1 className="text-5xl font-bold mb-4">404 - Page Not Found</h1>
        <p className="text-lg mb-6">
          We can't find the page you're looking for. Let's get you back on
          track.
        </p>
        <Link
          to="/"
          className="inline-block bg-gray-900 text-white px-6 py-3 rounded-md text-lg font-semibold hover:bg-gray-800 transition"
        >
          Go Home →
        </Link>
      </div>
    </div>
  );
}
