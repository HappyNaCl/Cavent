export default function debounce<
  T extends (...args: Parameters<T>) => ReturnType<T>
>(fn: T, delay: number): (...args: Parameters<T>) => void {
  let timeout: ReturnType<typeof setTimeout>;

  return (...args: Parameters<T>) => {
    clearTimeout(timeout);
    timeout = setTimeout(() => {
      fn(...args);
    }, delay);
  };
}
