import { format } from "date-fns";
import { ja } from "date-fns/locale/ja";
import { useEffect, useState } from "react";
import client from "~/api";
import type { components } from "~/api/client";
import { CreateReservationButton } from "~/components/create-reservation-button";
import { FilterReservationsButton } from "~/components/filter-reservations-button";
import { LogoutButton } from "~/components/logout-button";
import { ReservationItem } from "~/components/reservation-item";
import type { Route } from "./+types/home";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "練習室予約 ｜ 東京音楽大学" },
    {
      name: "description",
      content: "東京音楽大学の非公式練習室予約サイトです。",
    },
  ];
}

const groupReservationsByDate = (
  reservations: components["schemas"]["Reservation"][]
) => {
  const grouped: Record<string, components["schemas"]["Reservation"][]> = {};
  const now = new Date();
  const nowJST = new Date(
    now.getTime() + (9 * 60 + now.getTimezoneOffset()) * 60 * 1000
  );

  reservations.forEach((r) => {
    console.log(r.external_id);

    const date = new Date(r.date);
    const end = new Date(date);
    end.setHours(r.to_hour, r.to_minute, 0, 0);

    console.log("nowJST", nowJST);
    console.log("end", end);

    if (end <= nowJST) return; // 終了してたらスキップ

    const dateKey = format(new Date(r.date), "yyyy-MM-dd");
    if (!grouped[dateKey]) grouped[dateKey] = [];
    grouped[dateKey].push(r);
  });

  const sorted = Object.entries(grouped)
    .sort(([a], [b]) => a.localeCompare(b))
    .map(([date, reservations]) => ({
      date,
      formattedDate: format(new Date(date), "yyyy年M月d日(EEE)", {
        locale: ja,
      }),
      reservations: reservations.sort((a, b) => {
        const aStart = a.from_hour * 60 + a.from_minute;
        const bStart = b.from_hour * 60 + b.from_minute;
        return aStart - bStart;
      }),
    }));

  return sorted;
};

export default function Home() {
  const [rooms, setRooms] = useState<components["schemas"]["Room"][]>([]);
  const [reservations, setReservations] = useState<
    components["schemas"]["Reservation"][]
  >([]);

  useEffect(() => {
    const fetch = async () => {
      const [rsvsResponse, roomsResponse] = await Promise.all([
        client.GET("/reservations/mine"),
        client.GET("/rooms"),
      ]);

      if (rsvsResponse.data?.ok) {
        setReservations(rsvsResponse.data.data.reservations);
      }

      if (roomsResponse.data?.ok) {
        setRooms(roomsResponse.data.data.rooms);
      }
    };

    fetch();
  }, []);

  return (
    <>
      <div className="w-full min-h-dvh">
        <div className="flex justify-center items-center pt-8 pb-4">
          <span className="text-2xl">予約一覧</span>
        </div>
        <div className="mt-4 mx-4">
          <div className="grid gap-4">
            {groupReservationsByDate(reservations).map((group) => (
              <div key={group.date} className="grid">
                <span>{group.formattedDate}</span>
                <div className="flex gap-4 py-3 overflow-x-auto">
                  {group.reservations.map((r) => (
                    <ReservationItem
                      key={r.id}
                      isConfirmed={r.external_id !== undefined}
                      campusName={
                        r.campus_code === "1" ? "池袋" : "中目黒・代官山"
                      }
                      date={group.formattedDate}
                      timeRange={`${r.from_hour
                        .toString()
                        .padStart(2, "0")}:${r.from_minute
                        .toString()
                        .padStart(2, "0")} 〜 ${r.to_hour
                        .toString()
                        .padStart(2, "0")}:${r.to_minute
                        .toString()
                        .padStart(2, "0")}`}
                      userName={r.booker_name}
                      // roomName={roomIdToRoomNameMap[r.room_id] ?? "未定"}
                      roomName={r.room_id ?? "未定"}
                      pianoType={"グランドピアノ"}
                      reservationId={r.id}
                      onDelete={() => {
                        setReservations((prev) =>
                          prev.filter((reservation) => reservation.id !== r.id)
                        );
                      }}
                    />
                  ))}
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
      <div className="sticky bottom-0 border-t backdrop-blur-xl">
        <div className="flex justify-between items-center h-16 px-14 max-w-[500px] mx-auto">
          <FilterReservationsButton />
          <CreateReservationButton rooms={rooms} />
          <LogoutButton />
        </div>
      </div>
    </>
  );
}
