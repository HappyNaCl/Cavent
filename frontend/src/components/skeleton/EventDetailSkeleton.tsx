export default function EventDetailSkeleton() {
  return (
    <div className="mx-auto space-y-6 animate-pulse w-full">
      <div className="w-full h-[300px] bg-gray-200 rounded-lg" />

      <div className="flex justify-between items-center">
        <div className="h-6 w-1/3 bg-gray-300 rounded" />
        <div className="flex gap-2">
          <div className="h-10 w-10 bg-gray-300 rounded-full" />
          <div className="h-10 w-10 bg-gray-300 rounded-full" />
        </div>
      </div>

      <div className="space-y-2">
        <div className="h-5 w-1/4 bg-gray-300 rounded" />
        <div className="flex gap-2 items-center">
          <div className="w-4 h-4 bg-gray-300 rounded-full" />
          <div className="h-4 w-1/3 bg-gray-200 rounded" />
        </div>
        <div className="flex gap-2 items-center">
          <div className="w-4 h-4 bg-gray-300 rounded-full" />
          <div className="h-4 w-1/3 bg-gray-200 rounded" />
        </div>
        <div className="h-8 w-40 bg-gray-300 rounded" />
      </div>

      <div className="space-y-2">
        <div className="h-5 w-1/4 bg-gray-300 rounded" />
        <div className="h-4 w-2/3 bg-gray-200 rounded" />
      </div>

      <div className="space-y-2">
        <div className="h-5 w-1/4 bg-gray-300 rounded" />
        <div className="h-4 w-1/2 bg-gray-200 rounded" />
        <div className="h-8 w-32 bg-gray-300 rounded" />
      </div>

      <div className="space-y-2">
        <div className="h-5 w-1/4 bg-gray-300 rounded" />
        <div className="flex items-center gap-3">
          <div className="w-10 h-10 bg-gray-300 rounded-full" />
          <div className="space-y-2">
            <div className="h-4 w-32 bg-gray-300 rounded" />
            <div className="flex gap-2">
              <div className="h-8 w-20 bg-gray-300 rounded" />
              <div className="h-8 w-20 bg-gray-300 rounded" />
            </div>
          </div>
        </div>
      </div>

      <div className="space-y-2">
        <div className="h-5 w-1/4 bg-gray-300 rounded" />
        <div className="h-4 w-full bg-gray-200 rounded" />
        <div className="h-4 w-3/4 bg-gray-200 rounded" />
      </div>
    </div>
  );
}
