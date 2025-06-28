import { Button, DatePicker, Select, SelectItem, Slider } from "@heroui/react";
import { today } from "@internationalized/date";
import { useState } from "react";
import { I18nProvider } from "react-aria-components";

export function ReservationForm() {
  const [selectedRoom, setSelectedRoom] = useState<string | null>(null);

  const times = [
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
  ];

  const rooms = [
    "楽屋3（G）旧楽屋202",
    "P 205（G）",
    "P 220（G）",
    "P 222（U）",
    "P 441（G）",
    "P 456（G）",
    "P 460（G2台）",
  ];

  return (
    <div className="flex flex-col gap-6 w-full h-full">
      <div className="grid gap-3">
        <h4 className="text-sm text-default-700 opacity-60">キャンパス</h4>
        <Select isRequired placeholder="キャンパスを選択" name="campus">
          <SelectItem key="nakameguro">中目黒・代官山キャンパス</SelectItem>
          <SelectItem key="ikebukuro">池袋キャンパス</SelectItem>
        </Select>
      </div>
      <div className="grid gap-3">
        <h4 className="text-sm text-default-700 opacity-60">予約日</h4>
        <I18nProvider locale="ja">
          <DatePicker
            labelPlacement="outside"
            isRequired
            fullWidth
            minValue={today("Asia/Tokyo").add({ days: 3 })}
          />
        </I18nProvider>
      </div>
      <div className="grid gap-3">
        <h4 className="text-sm text-default-700 opacity-60">開始時刻</h4>
        <Select isRequired placeholder="開始時刻を選択" name="campus">
          {times.map((time) => {
            return <SelectItem key={time}>{time}</SelectItem>;
          })}
        </Select>
      </div>
      <div className="grid">
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
        />
      </div>
      <div className="grid gap-3">
        <h4 className="text-sm text-default-700 opacity-60">練習室</h4>
        <div className="grid grid-cols-2 gap-2">
          {rooms.map((room) => {
            const isSelected = selectedRoom === room;

            return (
              <Button
                key={room}
                fullWidth
                onPress={() => setSelectedRoom(room)}
                className={`text-xs font-semibold ${
                  isSelected
                    ? "bg-primary text-foreground"
                    : "bg-default-100 text-default-500"
                }`}
              >
                {room}
              </Button>
            );
          })}
        </div>
      </div>
      <Button fullWidth color="primary" className="mt-auto flex-shrink-0">
        予約する
      </Button>
      <span className="h-2 flex-shrink-0" />
    </div>
  );
}
