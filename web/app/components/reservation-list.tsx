import type { Reservation } from "~/api/pb/reservation/v1/reservation_pb";
import { ReservationListItem } from "./reservation-list-item";

export function ReservationList({
  reservations,
  onDelete,
  onNoteUpdated,
}: {
  reservations: Reservation[];
  onDelete?: (reservationId: string) => void;
  onNoteUpdated?: (reservationId: string, note: string | undefined) => void;
}) {
  const sorted = reservations.slice().sort((a, b) => {
    const getTime = (r: Reservation) => {
      const hour = String(r.fromHour).padStart(2, "0");
      const minute = String(r.fromMinute ?? 0).padStart(2, "0");
      return new Date(`${r.date}T${hour}:${minute}`).getTime();
    };
    return getTime(a) - getTime(b);
  });

  return (
    <div className="flex flex-col gap-6 p-6">
      {sorted.map((reservation) => (
        <ReservationListItem
          key={reservation.id}
          reservation={reservation}
          onDelete={onDelete}
          onNoteUpdated={onNoteUpdated}
        />
      ))}
    </div>
  );
}
