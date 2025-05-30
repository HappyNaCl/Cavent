import { useEffect, useState } from "react";

interface TypingCycleProps {
  words: string[];
  typingSpeed?: number;
  deletingSpeed?: number;
  pauseTime?: number;
  children?: React.ReactNode;
}

export default function SearchAnimation({
  children,
  words,
  typingSpeed = 100,
  deletingSpeed = 50,
  pauseTime = 1500,
}: TypingCycleProps) {
  const [text, setText] = useState("");
  const [isDeleting, setIsDeleting] = useState(false);
  const [loopIndex, setLoopIndex] = useState(0);

  useEffect(() => {
    const currentWord = words[loopIndex % words.length];

    const handleTyping = () => {
      if (isDeleting) {
        setText((prev) => prev.slice(0, -1));
      } else {
        setText((prev) => currentWord.slice(0, prev.length + 1));
      }

      if (!isDeleting && text === currentWord) {
        setTimeout(() => setIsDeleting(true), pauseTime);
      } else if (isDeleting && text === "") {
        setIsDeleting(false);
        setLoopIndex((prev) => prev + 1);
      }
    };

    const timeout = setTimeout(
      handleTyping,
      isDeleting ? deletingSpeed : typingSpeed
    );

    return () => clearTimeout(timeout);
  }, [
    text,
    isDeleting,
    loopIndex,
    words,
    typingSpeed,
    deletingSpeed,
    pauseTime,
  ]);

  return (
    <span className="font-montserrat font-bold text-4xl text-center">
      {children}
      <span>{text}</span>
      <span className="animate-caret-blink">|</span>
    </span>
  );
}
