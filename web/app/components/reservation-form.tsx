import {
  addToast,
  Button,
  cn,
  DatePicker,
  Form,
  Input,
  Select,
  SelectItem,
  Switch,
} from "@heroui/react";
import { CalendarDate, today } from "@internationalized/date";
import { useState } from "react";
import { I18nProvider, Label } from "react-aria-components";
import client from "~/api";

type Props = {
  rooms: any[];
  type: "create" | "update";
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
  "21:30",
  "22:00",
  "22:30",
];

const startTimeOptions = timeOptions.slice(0, -1);

export function ReservationForm(props: Props) {
  const [isLoading, setIsLoading] = useState(false);
  const [isAutoSelectRoom, setIsAutoSelectRoom] = useState(false);
  const [selectedCampus, setSelectedCampus] = useState<
    "nakameguro" | "ikebukuro" | null
  >(null);
  const [selectedDate, setSelectedDate] = useState<CalendarDate | null>(null);
  const [startTime, setStartTime] = useState<string | null>(null);
  const [endTime, setEndTime] = useState<string | null>(null);
  const [selectedRoom, setSelectedRoom] = useState<string | null>(null);
  const [selectedRoomOptions, setSelectedRoomOptions] = useState<string[]>([]);
  const [bookerName, setBookerName] = useState<string>("");

  const handleSubmit = async () => {
    setIsLoading(true);

    // const now = new Date();
    // const nowJST = new Date(
    //   now.getTime() + (9 * 60 + now.getTimezoneOffset()) * 60 * 1000
    // );

    if (!selectedDate) {
      setSelectedDate(today("Asia/Tokyo").add({ days: 3 }));
      return;
    }

    const date = selectedDate
      ? selectedDate.toDate("Asia/Tokyo")
      : today("Asia/Tokyo").add({ days: 3 }).toDate("Asia/Tokyo");
    const jst = new Date(date.getTime() + 9 * 60 * 60 * 1000);
    const jstISOString = jst.toISOString().replace("Z", "+09:00");

    console.log("Selected date:", jstISOString);

    if (props.type === "create") {
      // 予約を作成する処理
      let response;

      if (isAutoSelectRoom) {
        response = await client.POST("/reservations/create", {
          body: {
            campus_code: selectedCampus === "ikebukuro" ? "1" : "2",
            date: jstISOString,
            from_hour: Number(startTime?.split(":")[0]),
            from_minute: Number(startTime?.split(":")[1]),
            to_hour: Number(endTime?.split(":")[0]),
            to_minute: Number(endTime?.split(":")[1]),
            is_auto_select: isAutoSelectRoom,
            room_id: selectedRoom ?? undefined,
            booker_name: bookerName,
          },
        });
      } else {
        const pianoNumbers: number[] = [];
        const pianoTypes: ("grand" | "upright")[] = [];

        if (selectedRoomOptions.includes("multiple-piano")) {
          pianoNumbers.push(2);
        }
        if (pianoNumbers.length === 0) {
          pianoNumbers.push(1);
        }

        if (selectedRoomOptions.includes("grand")) {
          pianoTypes.push("grand");
        }
        if (selectedRoomOptions.includes("upright")) {
          pianoTypes.push("upright");
        }

        response = await client.POST("/reservations/create", {
          body: {
            campus_code: selectedCampus === "ikebukuro" ? "1" : "2",
            date: jstISOString,
            from_hour: Number(startTime?.split(":")[0]),
            from_minute: Number(startTime?.split(":")[1]),
            to_hour: Number(endTime?.split(":")[0]),
            to_minute: Number(endTime?.split(":")[1]),
            is_auto_select: isAutoSelectRoom,
            room_id: selectedRoom ?? undefined,
            booker_name: bookerName,
            piano_numbers: pianoNumbers,
            piano_types: pianoTypes,
            is_basement: selectedRoomOptions.includes("basement"),
          },
        });
      }

      if (!response.data?.ok) {
        addToast({
          title: "予約に失敗しました",
          color: "danger",
        });
        return;
      }

      props.onSuccess?.();
    } else {
      // 予約を更新する処理
      console.log("Updating reservation...");
    }

    setIsLoading(false);
  };

  return (
    <Form>
      <Label className="text-sm font-medium text-default-700">キャンパス</Label>
      <Select
        isRequired
        size="lg"
        className="mb-4"
        placeholder="キャンパスを選択"
        onSelectionChange={(key) => {
          const currentKey = key.currentKey as "nakameguro" | "ikebukuro";
          switch (currentKey) {
            case "nakameguro":
              setSelectedCampus("nakameguro");
              break;
            case "ikebukuro":
              setSelectedCampus("ikebukuro");
              break;
            default:
              setSelectedCampus(null);
          }
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
          className="mb-4"
          defaultValue={today("Asia/Tokyo").add({ days: 3 })}
          minValue={today("Asia/Tokyo").add({ days: 3 })}
          onChange={setSelectedDate}
        />
      </I18nProvider>
      <div className="flex gap-4 w-full mb-4">
        <div className="flex flex-col w-full gap-2">
          <Label className="text-sm font-medium text-default-700">
            開始時刻
          </Label>
          <Select
            isRequired
            size="lg"
            placeholder="開始時刻"
            onSelectionChange={(key) => setStartTime(key.currentKey as string)}
          >
            {startTimeOptions.map((time) => (
              <SelectItem key={time}>{time}</SelectItem>
            ))}
          </Select>
        </div>
        <div className="flex flex-col w-full gap-2">
          <Label className="text-sm font-medium text-default-700">
            終了時刻
          </Label>
          <Select
            isRequired
            size="lg"
            placeholder="終了時刻"
            isDisabled={!startTime}
            onSelectionChange={(key) => setEndTime(key.currentKey as string)}
          >
            {startTime
              ? timeOptions
                  .slice(timeOptions.indexOf(startTime) + 1)
                  .map((time) => <SelectItem key={time}>{time}</SelectItem>)
              : null}
          </Select>
        </div>
      </div>
      <Switch
        classNames={{
          base: cn(
            "inline-flex flex-row-reverse min-w-[100%] bg-content2 items-center",
            "justify-between cursor-pointer rounded-lg gap-2 mb-4 p-5 pl-2"
          ),
          wrapper: "p-0 h-4 overflow-visible",
          thumb: cn(
            "w-6 h-6 border-2 shadow-lg",
            "group-data-[hover=true]:border-primary",
            "group-data-[selected=true]:ms-6",
            "group-data-[pressed=true]:w-8",
            "group-data-[selected]:group-data-[pressed]:ms-4"
          ),
        }}
        isSelected={isAutoSelectRoom}
        onValueChange={setIsAutoSelectRoom}
      >
        <div className="grid gap-1">
          <p className="text-sm">部屋おまかせ</p>
          <p className="text-tiny text-default-400">
            条件に合う部屋を自動で予約します
          </p>
        </div>
      </Switch>
      <Label className="flex text-sm font-medium text-default-700">
        部屋
        {isAutoSelectRoom && (
          <span className="text-foreground-400">（任意）</span>
        )}
      </Label>
      {isAutoSelectRoom ? (
        <Select
          isDisabled={selectedCampus === null}
          size="lg"
          selectionMode="multiple"
          className="mb-4"
          placeholder="希望する条件を選択"
          onSelectionChange={(key) => {
            setSelectedRoomOptions(Array.from(key).map(String));
          }}
        >
          <SelectItem key="grand">グランドピアノ</SelectItem>
          <SelectItem key="upright">アップライトピアノ</SelectItem>
          <SelectItem key="single-piano">ピアノ1台</SelectItem>
          <SelectItem key="multiple-piano">ピアノ2台</SelectItem>
          {selectedCampus === "ikebukuro" ? (
            <SelectItem key="basement">地下</SelectItem>
          ) : (
            <SelectItem key="floor2">2階</SelectItem>
          )}
          <SelectItem key="floor4">4階</SelectItem>
          <SelectItem key="classroom">教室</SelectItem>
        </Select>
      ) : (
        <Select
          isRequired
          className="mb-4"
          size="lg"
          placeholder="部屋を選択"
          isDisabled={selectedCampus === null}
          onSelectionChange={(key) => setSelectedRoom(key.currentKey as string)}
        >
          {selectedCampus &&
            (() => {
              const campusCode = selectedCampus === "ikebukuro" ? "1" : "2";
              return props.rooms
                .filter((room) => room.campus_code === campusCode)
                .map((room) => (
                  <SelectItem key={room.id}>{room.name}</SelectItem>
                ));
            })()}
        </Select>
      )}
      <Label className="flex text-sm font-medium text-default-700">
        予約者名<span className="text-foreground-400">（任意）</span>
      </Label>
      <Input
        fullWidth
        size="lg"
        placeholder="予約者名を入力"
        className="mb-4"
        onValueChange={setBookerName}
      />
      <Button
        isDisabled={
          props.type === "create"
            ? selectedCampus === null ||
              startTime === null ||
              endTime === null ||
              (!isAutoSelectRoom && selectedRoom === null)
            : false
        }
        fullWidth
        size="lg"
        color="primary"
        onPress={handleSubmit}
        isLoading={isLoading}
      >
        {props.type === "create" ? "予約する" : "変更する"}
      </Button>
    </Form>
  );
}
