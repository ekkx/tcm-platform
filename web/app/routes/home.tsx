import { format } from "date-fns";
import { ja } from "date-fns/locale/ja";
import { useEffect, useState } from "react";
import { CreateReservationButton } from "~/components/create-reservation-button";
import { FilterReservationsButton } from "~/components/filter-reservations-button";
import { LogoutButton } from "~/components/logout-button";
import { ReservationItem } from "~/components/reservation-item";
import { CampusType, convertRoomToComponent, type Reservation, type HomeLoaderData } from "~/types/api";
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

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const intent = formData.get("intent")?.toString();

  if (intent === "create-reservation") {
    const { createAuthenticatedClient } = await import("~/api/grpc-client");
    const { ReservationServiceClient } = await import("~/proto/v1/reservation/reservation.js");
    const { CampusType: GrpcCampusType } = await import("~/proto/v1/room/room.js");

    const cookieHeader = request.headers.get("Cookie");
    const cookies = new Map<string, string>();
    
    if (cookieHeader) {
      cookieHeader.split(';').forEach(cookie => {
        const [key, value] = cookie.split('=').map(s => s.trim());
        if (key && value) {
          cookies.set(key, value);
        }
      });
    }

    const accessToken = cookies.get('access-token');
    if (!accessToken) {
      return { error: "認証が必要です" };
    }

    try {
      const reservationClient = createAuthenticatedClient(
        ReservationServiceClient,
        accessToken
      );

      const campusCode = formData.get("campus_code")?.toString();
      const dateStr = formData.get("date")?.toString();
      const fromHour = Number(formData.get("from_hour"));
      const fromMinute = Number(formData.get("from_minute"));
      const toHour = Number(formData.get("to_hour"));
      const toMinute = Number(formData.get("to_minute"));
      const roomId = formData.get("room_id")?.toString();
      const bookerName = formData.get("booker_name")?.toString();

      const campusType = campusCode === "2" ? GrpcCampusType.IKEBUKURO : GrpcCampusType.NAKAMEGURO;

      const result = await new Promise<any>((resolve, reject) => {
        reservationClient.createReservation(
          {
            reservation: {
              campusType,
              date: dateStr ? new Date(dateStr) : undefined,
              fromHour,
              fromMinute,
              toHour,
              toMinute,
              roomId: roomId || "",
              bookerName
            }
          },
          (error, response) => {
            if (error) {
              console.error("[Create Reservation] Failed:", error);
              reject(error);
            } else {
              resolve(response);
            }
          }
        );
      });

      return { success: true, reservations: result.reservations };
    } catch (error: any) {
      console.error("[Create Reservation] Error:", error);
      return { error: error.message || "予約の作成に失敗しました" };
    }
  }

  if (intent === "delete-reservation") {
    const { createAuthenticatedClient } = await import("~/api/grpc-client");
    const { ReservationServiceClient } = await import("~/proto/v1/reservation/reservation.js");

    const cookieHeader = request.headers.get("Cookie");
    const cookies = new Map<string, string>();
    
    if (cookieHeader) {
      cookieHeader.split(';').forEach(cookie => {
        const [key, value] = cookie.split('=').map(s => s.trim());
        if (key && value) {
          cookies.set(key, value);
        }
      });
    }

    const accessToken = cookies.get('access-token');
    if (!accessToken) {
      return { error: "認証が必要です" };
    }

    try {
      const reservationClient = createAuthenticatedClient(
        ReservationServiceClient,
        accessToken
      );

      const reservationId = Number(formData.get("reservation_id"));

      await new Promise<any>((resolve, reject) => {
        reservationClient.deleteReservation(
          { reservationId },
          (error, response) => {
            if (error) {
              console.error("[Delete Reservation] Failed:", error);
              reject(error);
            } else {
              resolve(response);
            }
          }
        );
      });

      return { success: true, deleted: true };
    } catch (error: any) {
      console.error("[Delete Reservation] Error:", error);
      return { error: error.message || "予約の削除に失敗しました" };
    }
  }

  if (intent === "update-reservation") {
    const { createAuthenticatedClient } = await import("~/api/grpc-client");
    const { ReservationServiceClient } = await import("~/proto/v1/reservation/reservation.js");
    const { CampusType: GrpcCampusType } = await import("~/proto/v1/room/room.js");

    const cookieHeader = request.headers.get("Cookie");
    const cookies = new Map<string, string>();
    
    if (cookieHeader) {
      cookieHeader.split(';').forEach(cookie => {
        const [key, value] = cookie.split('=').map(s => s.trim());
        if (key && value) {
          cookies.set(key, value);
        }
      });
    }

    const accessToken = cookies.get('access-token');
    if (!accessToken) {
      return { error: "認証が必要です" };
    }

    try {
      const reservationClient = createAuthenticatedClient(
        ReservationServiceClient,
        accessToken
      );

      const reservationId = Number(formData.get("reservation_id"));
      const campusCode = formData.get("campus_code")?.toString();
      const dateStr = formData.get("date")?.toString();
      const fromHour = Number(formData.get("from_hour"));
      const fromMinute = Number(formData.get("from_minute"));
      const toHour = Number(formData.get("to_hour"));
      const toMinute = Number(formData.get("to_minute"));
      const roomId = formData.get("room_id")?.toString();
      const bookerName = formData.get("booker_name")?.toString();

      const campusType = campusCode === "2" ? GrpcCampusType.IKEBUKURO : GrpcCampusType.NAKAMEGURO;

      const result = await new Promise<any>((resolve, reject) => {
        reservationClient.updateReservation(
          {
            reservationId,
            reservation: {
              campusType,
              date: dateStr ? new Date(dateStr) : undefined,
              fromHour,
              fromMinute,
              toHour,
              toMinute,
              roomId: roomId || "",
              bookerName
            }
          },
          (error, response) => {
            if (error) {
              console.error("[Update Reservation] Failed:", error);
              reject(error);
            } else {
              resolve(response);
            }
          }
        );
      });

      return { success: true, reservation: result.reservation };
    } catch (error: any) {
      console.error("[Update Reservation] Error:", error);
      return { error: error.message || "予約の更新に失敗しました" };
    }
  }

  return { error: "Invalid intent" };
}

export async function loader({ request }: Route.LoaderArgs): Promise<HomeLoaderData> {
  // Dynamic imports for server-side only
  const { createAuthenticatedClient } = await import("~/api/grpc-client");
  const { ReservationServiceClient } = await import("~/proto/v1/reservation/reservation.js");
  const { RoomServiceClient } = await import("~/proto/v1/room/room.js");
  const { CampusType: GrpcCampusType } = await import("~/proto/v1/room/room.js");
  const cookieHeader = request.headers.get("Cookie");
  const cookies = new Map<string, string>();
  
  if (cookieHeader) {
    cookieHeader.split(';').forEach(cookie => {
      const [key, value] = cookie.split('=').map(s => s.trim());
      if (key && value) {
        cookies.set(key, value);
      }
    });
  }

  const accessToken = cookies.get('access-token');
  if (!accessToken) {
    return { authenticated: false, rooms: [], reservations: [] };
  }

  try {
    // Create authenticated clients
    const reservationClient = createAuthenticatedClient(
      ReservationServiceClient,
      accessToken
    );
    const roomClient = createAuthenticatedClient(
      RoomServiceClient,
      accessToken
    );

    // Fetch data in parallel
    const [reservationsResult, roomsResult] = await Promise.all([
      new Promise<any>((resolve, reject) => {
        reservationClient.getUserReservations(
          {},
          (error, response) => {
            if (error) {
              console.error("[Home Loader] Failed to fetch reservations:", error);
              reject(error);
            } else {
              resolve(response);
            }
          }
        );
      }),
      new Promise<any>((resolve, reject) => {
        roomClient.getRooms(
          {},
          (error, response) => {
            if (error) {
              console.error("[Home Loader] Failed to fetch rooms:", error);
              reject(error);
            } else {
              resolve(response);
            }
          }
        );
      })
    ]);

    // Convert gRPC types to shared types
    const rooms = (roomsResult.rooms || []).map((room: any) => ({
      id: room.id,
      name: room.name,
      pianoType: room.pianoType,
      pianoNumber: room.pianoNumber,
      isClassroom: room.isClassroom,
      isBasement: room.isBasement,
      campusType: room.campusType === GrpcCampusType.IKEBUKURO ? CampusType.IKEBUKURO : CampusType.NAKAMEGURO,
      floor: room.floor
    }));

    const reservations = (reservationsResult.reservations || []).map((r: any) => ({
      id: r.id,
      externalId: r.externalId,
      campusType: r.campusType === GrpcCampusType.IKEBUKURO ? CampusType.IKEBUKURO : CampusType.NAKAMEGURO,
      date: r.date,
      roomId: r.roomId,
      fromHour: r.fromHour,
      fromMinute: r.fromMinute,
      toHour: r.toHour,
      toMinute: r.toMinute,
      bookerName: r.bookerName,
      createdAt: r.createdAt
    }));

    return {
      authenticated: true,
      rooms,
      reservations
    };
  } catch (error: any) {
    console.error("[Home Loader] Error fetching data:", error);
    // Check if it's an authentication error
    if (error.code === 16 || error.code === 7) {
      return { authenticated: false, rooms: [], reservations: [] };
    }
    // For other errors, still return data structure but with empty arrays
    return { authenticated: true, rooms: [], reservations: [] };
  }
}


const groupReservationsByDate = (
  reservations: Reservation[]
) => {
  const grouped: Record<string, Reservation[]> = {};
  const now = new Date();
  const nowJST = new Date(
    now.getTime() + (9 * 60 + now.getTimezoneOffset()) * 60 * 1000
  );

  reservations.forEach((r) => {
    console.log(r.externalId);

    if (!r.date) return;
    const date = new Date(r.date);
    const end = new Date(date);
    end.setHours(r.toHour, r.toMinute, 0, 0);

    console.log("nowJST", nowJST);
    console.log("end", end);

    if (end <= nowJST) return; // 終了してたらスキップ

    const dateKey = format(date, "yyyy-MM-dd");
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
        const aStart = a.fromHour * 60 + a.fromMinute;
        const bStart = b.fromHour * 60 + b.fromMinute;
        return aStart - bStart;
      }),
    }));

  return sorted;
};

export default function Home({ loaderData, actionData }: Route.ComponentProps) {
  const data = loaderData as HomeLoaderData | undefined;
  
  const [rooms] = useState<any[]>((data?.rooms || []).map(convertRoomToComponent));
  const [reservations, setReservations] = useState<Reservation[]>(
    data?.reservations || []
  );

  useEffect(() => {
    if (data && !data.authenticated) {
      window.location.href = "/login";
    }
  }, [data]);

  useEffect(() => {
    if (actionData?.success && actionData?.reservations) {
      // 予約作成成功後、ページをリロード
      window.location.reload();
    }
  }, [actionData]);

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
                      isConfirmed={r.externalId !== undefined}
                      campusName={
                        r.campusType === CampusType.IKEBUKURO ? "池袋" : "中目黒・代官山"
                      }
                      date={group.formattedDate}
                      timeRange={`${r.fromHour
                        .toString()
                        .padStart(2, "0")}:${r.fromMinute
                        .toString()
                        .padStart(2, "0")} 〜 ${r.toHour
                        .toString()
                        .padStart(2, "0")}:${r.toMinute
                        .toString()
                        .padStart(2, "0")}`}
                      userName={r.bookerName}
                      roomName={rooms.find(room => room.id === r.roomId)?.name || r.roomId || "未定"}
                      roomId={r.roomId}
                      pianoType={"グランドピアノ"}
                      reservationId={r.id}
                      rooms={rooms}
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
