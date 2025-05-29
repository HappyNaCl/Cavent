import { useState, useEffect } from "react";

type ImageProps = {
  src: string;
  alt: string;
  className?: string;
  maxRetries?: number;
  spinner?: React.ReactNode | null;
};

export default function Image({
  src,
  alt,
  className = "",
  maxRetries = 3,
  spinner = null,
}: ImageProps) {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);
  const [retryCount, setRetryCount] = useState(0);
  const [imgSrc, setImgSrc] = useState(src);

  useEffect(() => {
    if (retryCount > 0) {
      const timer = setTimeout(
        () => setImgSrc(`${src}?retry=${retryCount}`),
        1000
      ); // slight delay between retries
      return () => clearTimeout(timer);
    }
  }, [retryCount, src]);

  const handleLoad = () => {
    setLoading(false);
    setError(false);
  };

  const handleError = () => {
    if (retryCount < maxRetries) {
      setRetryCount((prev) => prev + 1);
    } else {
      setLoading(false);
      setError(true);
    }
  };

  return (
    <div className={`w-full relative inline-block ${className}`}>
      {loading &&
        (spinner || (
          <div className="absolute inset-0 flex items-center justify-center">
            <div className="w-6 h-6 border-4 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
          </div>
        ))}
      {!error ? (
        <img
          src={imgSrc}
          alt={alt}
          className={`transition-opacity duration-300 object-cover w-full h-full ${
            loading ? "opacity-0" : "opacity-100"
          }`}
          onLoad={handleLoad}
          onError={handleError}
        />
      ) : (
        <div className="text-red-500 text-sm p-2">Failed to load image.</div>
      )}
    </div>
  );
}
