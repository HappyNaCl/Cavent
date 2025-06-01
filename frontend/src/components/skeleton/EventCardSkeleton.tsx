export default function EventCardSkeleton() {
  return (
    <div className="max-w-sm rounded-xl overflow-hidden shadow bg-white animate-pulse">
      <div className="relative w-full h-40 bg-gray-200">
        <div className="absolute top-2 left-2 h-5 w-16 rounded bg-yellow-300/80" />
        <div className="absolute top-2 right-2 h-6 w-6 rounded-full bg-gray-300" />
      </div>

      <div className="p-4">
        <div className="flex gap-3">
          <div className="text-center space-y-1">
            <div className="h-4 w-10 bg-purple-200 rounded" />
            <div className="h-6 w-10 bg-gray-300 rounded" />
          </div>

          <div className="flex-1 space-y-2">
            <div className="h-5 w-full bg-gray-300 rounded" />
            <div className="h-4 w-1/2 bg-gray-200 rounded" />
            <div className="h-4 w-2/3 bg-gray-200 rounded" />

            <div className="flex gap-4 mt-2">
              <div className="flex items-center gap-2">
                <div className="w-4 h-4 bg-gray-300 rounded-full" />
                <div className="h-4 w-10 bg-gray-300 rounded" />
              </div>
              <div className="flex items-center gap-2">
                <div className="w-4 h-4 bg-gray-300 rounded-full" />
                <div className="h-4 w-16 bg-gray-300 rounded" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
