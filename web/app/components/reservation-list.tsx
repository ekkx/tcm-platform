import type { Reservation } from "~/api/pb/reservation/v1/reservation_pb";
import { ReservationListItem } from "./reservation-list-item";

// "2025-08" → "2025年8月"
const getYearMonthLabel = (yearMonth: string) => {
  const [_, month] = yearMonth.split("-");
  // return `${year}年${parseInt(month, 10)}月`;
  return `${parseInt(month, 10)}月`;
};

// 月ごとにグルーピング
const groupByYearMonth = (reservations: Reservation[]) => {
  return reservations.reduce((acc, reservation) => {
    const yearMonth = reservation.date.slice(0, 7); // "YYYY-MM"
    if (!acc[yearMonth]) {
      acc[yearMonth] = [];
    }
    acc[yearMonth].push(reservation);
    return acc;
  }, {} as Record<string, Reservation[]>);
};

export function ReservationList({
  reservations,
  onDelete,
}: {
  reservations: Reservation[];
  onDelete?: (reservationId: string) => void;
}) {
  const grouped = groupByYearMonth(reservations);
  const sortedYearMonths = Object.keys(grouped).sort();

  return (
    <div className="grid gap-6 px-4">
      {sortedYearMonths.map((yearMonth) => (
        <div key={yearMonth} className="grid gap-2">
          <h4 className="ml-2 text-sm text-foreground-400">
            {getYearMonthLabel(yearMonth)}
          </h4>
          <div className="flex flex-col gap-2">
            {grouped[yearMonth]
              .slice()
              .sort((a, b) => {
                const getTime = (r: Reservation) => {
                  const hour = String(r.fromHour).padStart(2, "0");
                  const minute = String(r.fromMinute ?? 0).padStart(2, "0");
                  return new Date(`${r.date}T${hour}:${minute}`).getTime();
                };
                return getTime(a) - getTime(b);
              })
              .map((reservation) => (
                <ReservationListItem
                  key={reservation.id}
                  reservation={reservation}
                  onDelete={onDelete}
                />
              ))}
          </div>
        </div>
      ))}
    </div>
  );
}
