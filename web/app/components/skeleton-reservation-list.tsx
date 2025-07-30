import { Skeleton } from "@heroui/react";
import { SkeletonReservationListItem } from "./skeleton-reservation-list-item";

export function SkeletonReservationList() {
  return (
    <div className="grid gap-6 px-4">
      <div className="grid gap-3">
        <Skeleton className="w-8 h-3 ml-1 rounded-lg" />
        <div className="flex flex-col gap-2">
          <SkeletonReservationListItem />
          <SkeletonReservationListItem />
          <SkeletonReservationListItem />
        </div>
      </div>
      <div className="grid gap-3">
        <Skeleton className="w-8 h-3 ml-1 rounded-lg" />
        <div className="flex flex-col gap-2">
          <SkeletonReservationListItem />
          <SkeletonReservationListItem />
        </div>
      </div>
      <div className="grid gap-3">
        <Skeleton className="w-8 h-3 ml-1 rounded-lg" />
        <div className="flex flex-col gap-2">
          <SkeletonReservationListItem />
          <SkeletonReservationListItem />
          <SkeletonReservationListItem />
          <SkeletonReservationListItem />
        </div>
      </div>
    </div>
  );
}
