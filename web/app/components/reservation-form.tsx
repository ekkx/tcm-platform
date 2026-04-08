import {
  Button,
  Card,
  DatePicker,
  Select,
  SelectItem,
  Spinner,
  Textarea,
  type DateValue,
} from "@heroui/react";
import { today } from "@internationalized/date";
import { useMemo, useState } from "react";
import { I18nProvider } from "react-aria-components";
import { reservationClient, roomClient } from "~/api";
import type { Reservation } from "~/api/pb/reservation/v1/reservation_pb";
import { CampusType, type Room } from "~/api/pb/room/v1/room_pb";

// 30分刻みの時刻リストを生成する
function generateTimes(
  startHour: number,
  startMinute: number,
  endHour: number,
  endMinute: number
): string[] {
  const times: string[] = [];
  let h = startHour;
  let m = startMinute;
  while (h < endHour || (h === endHour && m <= endMinute)) {
    times.push(`${h}:${String(m).padStart(2, "0")}`);
    m += 30;
    if (m >= 60) {
      h += 1;
      m = 0;
    }
  }
  return times;
}

function parseTime(time: string): { hour: number; minute: number } {
  const [h, m] = time.split(":");
  return { hour: parseInt(h, 10), minute: parseInt(m, 10) };
}

// キャンパスごとの最終時刻
const CAMPUS_END_TIME: Record<string, { hour: number; minute: number }> = {
  nakameguro: { hour: 21, minute: 30 },
  ikebukuro: { hour: 22, minute: 30 },
};

const getCampusType = (key: "nakameguro" | "ikebukuro" | null): CampusType => {
  switch (key) {
    case "nakameguro":
      return CampusType.NAKAMEGURO;
    case "ikebukuro":
      return CampusType.IKEBUKURO;
    default:
      return CampusType.UNSPECIFIED;
  }
};

export function ReservationForm({
  onReservationCreated,
  onReservationFailed,
}: {
  onReservationCreated?: (reservation: Reservation) => void;
  onReservationFailed?: (error: Error) => void;
}) {
  const [selectedCampus, setSelectedCampus] = useState<
    "nakameguro" | "ikebukuro" | null
  >(null);
  const [selectedDate, setSelectedDate] = useState<DateValue | null>(null);
  const [selectedStartTime, setSelectedStartTime] = useState<string | null>(
    null
  );
  const [selectedEndTime, setSelectedEndTime] = useState<string | null>(null);
  const [availableRooms, setAvailableRooms] = useState<Room[]>([]);
  const [selectedRoomId, setSelectedRoomId] = useState<string | null>(null);
  const [isLoadingRooms, setIsLoadingRooms] = useState(false);
  const [note, setNote] = useState("");
  const [isReservating, setIsReservating] = useState(false);

  // キャンパスに応じた開始時刻の選択肢
  const startTimes = useMemo(() => {
    if (!selectedCampus) return [];
    const end = CAMPUS_END_TIME[selectedCampus];
    // 開始時刻は最終時刻の30分前まで（最低30分の予約を確保）
    return generateTimes(
      7,
      30,
      end.hour,
      end.minute - 30 >= 0 ? end.minute - 30 : end.minute
    );
  }, [selectedCampus]);

  // 開始時刻の後の終了時刻の選択肢
  const endTimes = useMemo(() => {
    if (!selectedCampus || !selectedStartTime) return [];
    const start = parseTime(selectedStartTime);
    const end = CAMPUS_END_TIME[selectedCampus];
    // 開始時刻の30分後から最終時刻まで
    let startM = start.minute + 30;
    let startH = start.hour;
    if (startM >= 60) {
      startH += 1;
      startM = 0;
    }
    return generateTimes(startH, startM, end.hour, end.minute);
  }, [selectedCampus, selectedStartTime]);

  const canCreateReservation = () => {
    return (
      selectedCampus !== null &&
      selectedDate !== null &&
      selectedStartTime !== null &&
      selectedEndTime !== null &&
      selectedRoomId !== null
    );
  };

  const clearRoomSelection = () => {
    setSelectedRoomId(null);
    setAvailableRooms([]);
  };

  const resetEndTimeAndRooms = () => {
    setSelectedEndTime(null);
    clearRoomSelection();
  };

  const handleSelectEndTime = async (endTime: string) => {
    if (!selectedStartTime || !selectedCampus || !selectedDate) return;

    setSelectedEndTime(endTime);
    setIsLoadingRooms(true);
    setAvailableRooms([]);
    setSelectedRoomId(null);

    const from = parseTime(selectedStartTime);
    const to = parseTime(endTime);

    try {
      const response = await roomClient.listAvailableRooms({
        campusType: getCampusType(selectedCampus),
        date: selectedDate.toString(),
        fromHour: from.hour,
        fromMinute: from.minute,
        toHour: to.hour,
        toMinute: to.minute,
      });

      setAvailableRooms(response.rooms);
    } catch (error) {
      clearRoomSelection();
    } finally {
      setIsLoadingRooms(false);
    }
  };

  const handleCreateReservation = async () => {
    if (!canCreateReservation()) return;

    const from = parseTime(selectedStartTime!);
    const to = parseTime(selectedEndTime!);

    setIsReservating(true);

    try {
      const response = await reservationClient.createReservation({
        campusType: getCampusType(selectedCampus),
        date: selectedDate?.toString(),
        fromHour: from.hour,
        fromMinute: from.minute,
        toHour: to.hour,
        toMinute: to.minute,
        roomId: selectedRoomId!,
        note: note.trim() || undefined,
      });
      onReservationCreated?.(response.reservation!);
    } catch (error) {
      onReservationFailed?.(
        error instanceof Error ? error : new Error(String(error))
      );
    }

    clearRoomSelection();
    setIsReservating(false);
  };

  return (
    <div className="flex flex-col gap-6 w-full h-full">
      <div className="grid gap-3 px-6">
        <h4 className="text-sm text-default-700 opacity-60">キャンパス</h4>
        <Select
          isRequired
          placeholder="キャンパスを選択"
          name="campus"
          selectedKeys={selectedCampus ? [selectedCampus] : []}
          onChange={(event) => {
            setSelectedCampus(event.target.value as "nakameguro" | "ikebukuro");
            setSelectedStartTime(null);
            resetEndTimeAndRooms();
          }}
        >
          <SelectItem key="nakameguro">中目黒・代官山キャンパス</SelectItem>
          <SelectItem key="ikebukuro">池袋キャンパス</SelectItem>
        </Select>
      </div>
      <div className="grid gap-3 px-6">
        <h4 className="text-sm text-default-700 opacity-60">予約日</h4>
        <I18nProvider locale="ja">
          <DatePicker
            labelPlacement="outside"
            isRequired
            fullWidth
            minValue={today("Asia/Tokyo").add({ days: 3 })}
            value={selectedDate}
            onChange={(value) => {
              setSelectedDate(value);
              setSelectedStartTime(null);
              resetEndTimeAndRooms();
            }}
          />
        </I18nProvider>
      </div>
      <div className="grid gap-3 px-6">
        <h4 className="text-sm text-default-700 opacity-60">開始時刻</h4>
        <Select
          isRequired
          isDisabled={!selectedCampus}
          placeholder="開始時刻を選択"
          name="startTime"
          selectedKeys={selectedStartTime ? [selectedStartTime] : []}
          onChange={(event) => {
            setSelectedStartTime(event.target.value);
            resetEndTimeAndRooms();
          }}
        >
          {startTimes.map((time) => (
            <SelectItem key={time}>{time}</SelectItem>
          ))}
        </Select>
      </div>
      <div className="grid gap-3 px-6">
        <h4 className="text-sm text-default-700 opacity-60">終了時刻</h4>
        <Select
          isRequired
          isDisabled={!selectedStartTime}
          placeholder="終了時刻を選択"
          name="endTime"
          selectedKeys={selectedEndTime ? [selectedEndTime] : []}
          onChange={(event) => {
            handleSelectEndTime(event.target.value);
          }}
        >
          {endTimes.map((time) => (
            <SelectItem key={time}>{time}</SelectItem>
          ))}
        </Select>
      </div>
      <div className="grid gap-3 px-6">
        <h4 className="text-sm text-default-700 opacity-60">メモ</h4>
        <Textarea
          placeholder="メモを入力"
          value={note}
          onValueChange={setNote}
          maxRows={3}
          classNames={{ input: "text-base" }}
        />
      </div>
      <div className="grid gap-3 px-6 pb-24">
        <h4 className="text-sm text-default-700 opacity-60">練習室</h4>
        <div>
          {!selectedEndTime ? (
            <div className="grid place-items-center h-32">
              <p className="text-sm text-default-400">
                終了時刻を選択してください
              </p>
            </div>
          ) : isLoadingRooms ? (
            <div className="grid place-items-center h-32">
              <Spinner size="lg" variant="wave" />
            </div>
          ) : availableRooms.length === 0 ? (
            <div className="grid place-items-center h-32">
              <p className="text-sm text-default-400">
                利用可能な練習室が見つかりませんでした
              </p>
            </div>
          ) : (
            <div className="grid grid-cols-2 gap-2">
              {availableRooms.map((room) => {
                const isSelected = selectedRoomId === room.id;

                return (
                  <Button
                    key={room.id}
                    fullWidth
                    onPress={() => setSelectedRoomId(room.id)}
                    className={`text-xs font-semibold ${
                      isSelected
                        ? "bg-primary text-foreground"
                        : "bg-default-100 text-default-500"
                    }`}
                  >
                    {room.name}
                  </Button>
                );
              })}
            </div>
          )}
        </div>
      </div>
      <Card className="absolute w-full left-0 bottom-0 p-4 rounded-t-none rounded-r-none bg-content1/20 backdrop-blur-xl">
        <Button
          fullWidth
          color="primary"
          className="mt-auto flex-shrink-0"
          isDisabled={!canCreateReservation()}
          isLoading={isReservating}
          onPress={handleCreateReservation}
        >
          予約する
        </Button>
      </Card>
    </div>
  );
}
