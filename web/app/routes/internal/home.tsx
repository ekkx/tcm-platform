import { useMemo, useEffect, useState } from "react";
import { useLocation, useOutletContext } from "react-router";
import { reservationClient } from "~/api";
import type { Reservation } from "~/api/pb/reservation/v1/reservation_pb";
import type { FilterValues } from "~/components/header";
import { ReservationList } from "~/components/reservation-list";
import { SkeletonReservationList } from "~/components/skeleton-reservation-list";

export default function Home() {
  const location = useLocation();
  const { filters } = useOutletContext<{ filters: FilterValues }>();
  const newReservation = location.state?.newReservation as
    | Reservation
    | undefined;
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [isFetchingReservations, setIsFetchingReservations] = useState(true);

  useEffect(() => {
    (async () => {
      setIsFetchingReservations(true);
      const response = await reservationClient.listReservations({});
      setReservations(response.reservations);
      setIsFetchingReservations(false);
    })();
  }, []);

  useEffect(() => {
    if (newReservation) {
      setReservations((prev) => {
        const exists = prev.some((r) => r.id === newReservation.id);
        return exists ? prev : [...prev, newReservation];
      });
    }
  }, [newReservation]);

  const filteredReservations = useMemo(() => {
    return reservations.filter((r) => {
      if (filters.dateRange) {
        const start = filters.dateRange.start.toString();
        const end = filters.dateRange.end.toString();
        if (r.date < start || r.date > end) return false;
      }
      if (filters.campusType !== null) {
        if (r.campusType !== filters.campusType) return false;
      }
      if (filters.pianoType !== null) {
        if (r.room?.pianoType !== filters.pianoType) return false;
      }
      return true;
    });
  }, [reservations, filters]);

  const handleDeleteReservation = (reservationId: string) => {
    setReservations((prev) => prev.filter((r) => r.id !== reservationId));
  };

  const handleNoteUpdated = (
    reservationId: string,
    note: string | undefined
  ) => {
    setReservations((prev) =>
      prev.map((r) => (r.id === reservationId ? { ...r, note } : r))
    );
  };

  return (
    <div>
      {isFetchingReservations ? (
        <SkeletonReservationList />
      ) : filteredReservations.length === 0 ? (
        <div className="grid place-items-center h-36">
          <p className="text-sm text-default-400">予約はありません</p>
        </div>
      ) : (
        <ReservationList
          reservations={filteredReservations}
          onDelete={handleDeleteReservation}
          onNoteUpdated={handleNoteUpdated}
        />
      )}
    </div>
  );
}
