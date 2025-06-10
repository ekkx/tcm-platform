import {
  addToast,
  Button,
  DatePicker,
  Input,
  Select,
  SelectItem,
} from "@heroui/react";
import { CalendarDate, today } from "@internationalized/date";
import { useEffect, useState } from "react";
import { I18nProvider, Label } from "react-aria-components";
import { Form, useActionData, useNavigation } from "react-router";

type Props = {
  rooms: any[];
  type: "create" | "update";
  reservationId?: number;
  defaultCampus?: "nakameguro" | "ikebukuro";
  defaultDate?: string;
  defaultStartTime?: string;
  defaultEndTime?: string;
  defaultRoomName?: string;
  defaultUserName?: string;
  onSuccess?: () => void;
};

const timeOptions = [
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
];

export function ReservationForm(props: Props) {
  const navigation = useNavigation();
  const actionData = useActionData() as any;
  const isSubmitting = navigation.state === "submitting";


  const [selectedCampus, setSelectedCampus] = useState<
    "nakameguro" | "ikebukuro" | null
  >(props.defaultCampus ?? null);
  const [selectedDate, setSelectedDate] = useState<CalendarDate | null>(() => {
    if (props.type === "update" && props.defaultDate) {
      // For update mode, use the ISO date string from reservation
      const date = new Date(props.defaultDate);
      return new CalendarDate(date.getFullYear(), date.getMonth() + 1, date.getDate());
    } else if (props.type === "create") {
      // For create mode, default to 3 days from today
      const todayDate = today("Asia/Tokyo");
      return todayDate.add({ days: 3 });
    }
    return null;
  });
  const [startTime, setStartTime] = useState<string | null>(
    props.defaultStartTime ?? null
  );
  const [endTime, setEndTime] = useState<string | null>(
    props.defaultEndTime ?? null
  );
  const [selectedRoom, setSelectedRoom] = useState<string | null>(
    props.defaultRoomName ?? null
  );
  const [bookerName, setBookerName] = useState<string>(
    props.defaultUserName ?? ""
  );

  useEffect(() => {
    if (actionData?.success) {
      addToast({
        title:
          props.type === "create" ? "予約を作成しました" : "予約を更新しました",
        color: "success",
      });
      props.onSuccess?.();
      // Don't reload here - let the parent component handle it
    } else if (actionData?.error) {
      addToast({
        title: actionData.error,
        color: "danger",
      });
    }
  }, [actionData]);

  const campusRooms = props.rooms.filter((room) => {
    if (selectedCampus === "nakameguro") {
      return room.campus_code === "1";
    } else if (selectedCampus === "ikebukuro") {
      return room.campus_code === "2";
    }
    return false;
  });

  // Filter end time options based on selected start time
  const getFilteredEndTimeOptions = () => {
    if (!startTime) return timeOptions;

    const [startHour, startMinute] = startTime.split(":").map(Number);
    const startTotalMinutes = startHour * 60 + startMinute;

    return timeOptions.filter((time) => {
      const [hour, minute] = time.split(":").map(Number);
      const totalMinutes = hour * 60 + minute;
      return totalMinutes > startTotalMinutes;
    });
  };

  return (
    <Form method="post" className="flex flex-col h-full">
      <input type="hidden" name="intent" value={`${props.type}-reservation`} />
      {props.reservationId && (
        <input
          type="hidden"
          name="reservation_id"
          value={props.reservationId}
        />
      )}
      {selectedCampus && (
        <input
          type="hidden"
          name="campus_code"
          value={selectedCampus === "ikebukuro" ? "2" : "1"}
        />
      )}

      <Label className="text-sm font-medium text-default-700">キャンパス</Label>
      <Select
        isRequired
        size="lg"
        className="mt-2 mb-4"
        placeholder="キャンパスを選択"
        name="campus_code"
        selectedKeys={selectedCampus ? [selectedCampus] : []}
        onSelectionChange={(key) => {
          const currentKey = key.currentKey as "nakameguro" | "ikebukuro";
          setSelectedCampus(currentKey);
        }}
      >
        <SelectItem key="nakameguro">中目黒・代官山キャンパス</SelectItem>
        <SelectItem key="ikebukuro">池袋キャンパス</SelectItem>
      </Select>

      <Label className="text-sm font-medium text-default-700">予約日</Label>
      <I18nProvider locale="ja">
        <DatePicker
          isRequired
          fullWidth
          size="lg"
          className="mt-2 mb-4"
          minValue={today("Asia/Tokyo").add({ days: 3 })}
          value={selectedDate}
          onChange={setSelectedDate}
        />
      </I18nProvider>
      {selectedDate && (
        <input
          type="hidden"
          name="date"
          value={(() => {
            const date = selectedDate.toDate("Asia/Tokyo");
            // Create a date string that represents the date in JST
            const year = date.getFullYear();
            const month = String(date.getMonth() + 1).padStart(2, '0');
            const day = String(date.getDate()).padStart(2, '0');
            // Send as JST midnight
            return `${year}-${month}-${day}T00:00:00+09:00`;
          })()}
        />
      )}

      <div className="grid grid-cols-2 gap-3 mb-4">
        <div>
          <Label className="text-sm font-medium text-default-700">
            開始時刻
          </Label>
          <Select
            isRequired
            size="lg"
            className="mt-2"
            placeholder="開始時刻"
            name="start_time"
            selectedKeys={startTime ? [startTime] : []}
            onSelectionChange={(key) => {
              const newStartTime = key.currentKey as string;
              setStartTime(newStartTime);

              // Clear end time if it's now invalid
              if (endTime && newStartTime) {
                const [startHour, startMinute] = newStartTime
                  .split(":")
                  .map(Number);
                const [endHour, endMinute] = endTime.split(":").map(Number);
                const startTotalMinutes = startHour * 60 + startMinute;
                const endTotalMinutes = endHour * 60 + endMinute;

                if (endTotalMinutes <= startTotalMinutes) {
                  setEndTime(null);
                }
              }
            }}
          >
            {timeOptions.map((time) => (
              <SelectItem key={time}>{time}</SelectItem>
            ))}
          </Select>
        </div>
        <div>
          <Label className="text-sm font-medium text-default-700">
            終了時刻
          </Label>
          <Select
            isRequired
            size="lg"
            className="mt-2"
            placeholder="終了時刻"
            name="end_time"
            selectedKeys={endTime ? [endTime] : []}
            onSelectionChange={(key) => {
              setEndTime(key.currentKey as string);
            }}
          >
            {getFilteredEndTimeOptions().map((time) => (
              <SelectItem key={time}>{time}</SelectItem>
            ))}
          </Select>
        </div>
      </div>

      {startTime && (
        <>
          <input
            type="hidden"
            name="from_hour"
            value={startTime.split(":")[0]}
          />
          <input
            type="hidden"
            name="from_minute"
            value={startTime.split(":")[1]}
          />
        </>
      )}
      {endTime && (
        <>
          <input type="hidden" name="to_hour" value={endTime.split(":")[0]} />
          <input type="hidden" name="to_minute" value={endTime.split(":")[1]} />
        </>
      )}

      <Label className="text-sm font-medium text-default-700">部屋</Label>
      <Select
        isRequired
        size="lg"
        placeholder="部屋を選択"
        className="mt-2 mb-4"
        name="room_id"
        selectedKeys={selectedRoom ? [selectedRoom] : []}
        onSelectionChange={(key) => {
          setSelectedRoom(key.currentKey as string);
        }}
        isDisabled={!selectedCampus}
      >
        {campusRooms.map((room) => (
          <SelectItem key={room.id}>{room.name}</SelectItem>
        ))}
      </Select>

      <Label className="flex text-sm font-medium text-default-700">
        予約者名<span className="text-foreground-400">（任意）</span>
      </Label>
      <Input
        size="lg"
        placeholder="名前を入力"
        className="mt-2 mb-4"
        name="booker_name"
        value={bookerName}
        onChange={(e) => setBookerName(e.target.value)}
      />

      <Button
        isLoading={isSubmitting}
        type="submit"
        color="primary"
        size="lg"
        className="w-full mt-auto"
        isDisabled={
          !selectedCampus ||
          !selectedDate ||
          !startTime ||
          !endTime ||
          !selectedRoom
        }
      >
        {props.type === "create" ? "予約する" : "更新する"}
      </Button>
    </Form>
  );
}
