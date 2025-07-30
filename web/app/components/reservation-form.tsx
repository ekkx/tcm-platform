import {
  Button,
  Card,
  DatePicker,
  Select,
  SelectItem,
  Slider,
  Spinner,
  type DateValue,
} from "@heroui/react";
import { today } from "@internationalized/date";
import { useState } from "react";
import { I18nProvider } from "react-aria-components";
import { reservationClient, roomClient } from "~/api";
import type { Reservation } from "~/api/pb/reservation/v1/reservation_pb";
import { CampusType, type Room } from "~/api/pb/room/v1/room_pb";

const selectableTimes = [
  "7:30",
  "8:00",
  "8:30",
  "9:00",
  "9:30",
  "10:00",
  "10:30",
  "11:00",
  "11:30",
  "12:00",
  "12:30",
  "13:00",
  "13:30",
  "14:00",
  "14:30",
  "15:00",
  "15:30",
  "16:00",
  "16:30",
  "17:00",
  "17:30",
  "18:00",
  "18:30",
  "19:00",
  "19:30",
  "20:00",
  "20:30",
  "21:00",
  "21:30",
  "22:00",
];

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
  const [sliderValue, setSliderValue] = useState<number>(0);
  const [availableRooms, setAvailableRooms] = useState<Room[]>([]);
  const [selectedRoomId, setSelectedRoomId] = useState<string | null>(null);
  const [isLoadingRooms, setIsLoadingRooms] = useState(false);
  const [isReservating, setIsReservating] = useState(false);
  const [selectedFromHour, setSelectedFromHour] = useState<number | null>(null);
  const [selectedFromMinute, setSelectedFromMinute] = useState<number | null>(
    null
  );
  const [selectedToHour, setSelectedToHour] = useState<number | null>(null);
  const [selectedToMinute, setSelectedToMinute] = useState<number | null>(null);

  const canSelectTimeRange = () => {
    return (
      selectedCampus !== null &&
      selectedDate !== null &&
      selectedStartTime !== null
    );
  };

  const canCreateReservation = () => {
    return canSelectTimeRange() && sliderValue > 0 && selectedRoomId !== null;
  };

  const clearRoomSelection = () => {
    setSelectedRoomId(null);
    setAvailableRooms([]);
  };

  const resetTimeAndRooms = () => {
    clearRoomSelection();
    setSliderValue(0);
  };

  const getTimeRangeFromSlider = (startTime: string, value: number) => {
    const [fromHourStr, fromMinuteStr] = startTime.split(":");
    const fromHour = parseInt(fromHourStr, 10);
    const fromMinute = parseInt(fromMinuteStr, 10);
    const startMinutes = fromHour * 60 + fromMinute;

    const durationMinutes = Math.round((value / 100) * 6 * 60);
    const endMinutes = startMinutes + durationMinutes;

    return {
      fromHour,
      fromMinute,
      toHour: Math.floor(endMinutes / 60),
      toMinute: endMinutes % 60,
    };
  };

  const handleSelectTimeRange = async (value: number | number[]) => {
    const actualValue = Array.isArray(value) ? value[0] : value;
    if (!selectedStartTime) return;

    const { fromHour, fromMinute, toHour, toMinute } = getTimeRangeFromSlider(
      selectedStartTime,
      actualValue
    );

    if (fromHour === toHour && fromMinute === toMinute) return;

    setIsLoadingRooms(true);
    setAvailableRooms([]);

    try {
      const response = await roomClient.listAvailableRooms({
        campusType: getCampusType(selectedCampus),
        date: selectedDate?.toString(),
        fromHour: fromHour,
        fromMinute: fromMinute,
        toHour: toHour,
        toMinute: toMinute,
      });

      await new Promise((resolve) => setTimeout(resolve, 1500));

      setAvailableRooms(response.rooms);
    } catch (error) {
      clearRoomSelection();
    } finally {
      setIsLoadingRooms(false);
    }

    setSelectedFromHour(fromHour);
    setSelectedFromMinute(fromMinute);
    setSelectedToHour(toHour);
    setSelectedToMinute(toMinute);
  };

  const handleCreateReservation = async () => {
    if (!canCreateReservation()) return;

    setIsReservating(true);

    try {
      const response = await reservationClient.createReservation({
        campusType: getCampusType(selectedCampus),
        date: selectedDate?.toString(),
        fromHour: selectedFromHour!,
        fromMinute: selectedFromMinute!,
        toHour: selectedToHour!,
        toMinute: selectedToMinute!,
        roomId: selectedRoomId!,
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
            resetTimeAndRooms();
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
              resetTimeAndRooms();
            }}
          />
        </I18nProvider>
      </div>
      <div className="grid gap-3 px-6">
        <h4 className="text-sm text-default-700 opacity-60">開始時刻</h4>
        <Select
          isRequired
          placeholder="開始時刻を選択"
          name="campus"
          selectedKeys={selectedStartTime ? [selectedStartTime] : []}
          onChange={(event) => {
            setSelectedStartTime(event.target.value);
            resetTimeAndRooms();
          }}
        >
          {selectableTimes.map((time) => {
            return <SelectItem key={time}>{time}</SelectItem>;
          })}
        </Select>
      </div>
      <div className="grid px-6">
        <h4 className="text-sm text-default-700 opacity-60">利用時間</h4>
        <Slider
          showSteps
          color="foreground"
          label=" "
          getValue={(value) => {
            const hours = (parseInt(value.toString()) / 100) * 6;
            const rounded = Math.ceil(hours * 2) / 2;
            const h = Math.floor(rounded);
            const m = rounded % 1 === 0.5 ? 30 : 0;
            if (h === 0 && m === 30) return "30分";
            if (h > 0 && m === 0) return `${h}時間`;
            if (h > 0 && m === 30) return `${h}時間30分`;
            return "0分";
          }}
          formatOptions={{ style: "unit", unit: "hour" }}
          marks={[
            {
              value: (100 / 6) * 1,
              label: "1h",
            },
            {
              value: (100 / 6) * 2,
              label: "2h",
            },
            {
              value: (100 / 6) * 3,
              label: "3h",
            },
            {
              value: (100 / 6) * 4,
              label: "4h",
            },
            {
              value: (100 / 6) * 5,
              label: "5h",
            },
          ]}
          size="lg"
          step={(100 / 6) * 0.5}
          isDisabled={!canSelectTimeRange()}
          onChange={(value) => {
            // 選択中のとき
            const actualValue = Array.isArray(value) ? value[0] : value;
            setSliderValue(actualValue);
            clearRoomSelection();
            setIsLoadingRooms(true);
          }}
          onChangeEnd={handleSelectTimeRange}
        />
      </div>
      <div className="grid gap-3 px-6 pb-24">
        <h4 className="text-sm text-default-700 opacity-60">練習室</h4>
        <div>
          {sliderValue === 0 ? (
            <div className="grid place-items-center h-32">
              <p className="text-sm text-default-400">
                利用時間を選択してください
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
