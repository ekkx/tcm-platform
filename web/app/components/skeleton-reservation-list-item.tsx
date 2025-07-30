import { Card, Skeleton } from "@heroui/react";

export function SkeletonReservationListItem() {
  return (
    <Card className="flex flex-row items-center p-3">
      <Skeleton className="w-12 h-12 rounded-lg" />
      <div className="flex flex-col h-full justify-between ml-5 mr-auto">
        <Skeleton className="w-28 h-3 rounded-lg" />
        <Skeleton className="w-36 h-3 rounded-lg" />
        <Skeleton className="w-16 h-3 rounded-lg" />
      </div>
      <Skeleton className="w-16 h-8 rounded-lg" />
    </Card>
  );
}
