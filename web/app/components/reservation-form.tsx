import {
  Button,
  cn,
  DatePicker,
  Form,
  Input,
  Select,
  SelectItem,
  SelectSection,
  Switch,
} from "@heroui/react";
import { today } from "@internationalized/date";
import { useState } from "react";
import { I18nProvider, Label } from "react-aria-components";

type Props = {
  type: "create" | "update";
  defaultCampus?: "nakameguro" | "ikebukuro";
  defaultDate?: string;
  defaultStartTime?: string;
  defaultEndTime?: string;
  defaultRoomName?: string;
  defaultUserName?: string;
};

export function ReservationForm(props: Props) {
  const [isAutoSelectRoom, setIsAutoSelectRoom] = useState(false);

  return (
    <Form>
      <Label className="text-sm font-medium text-default-700">キャンパス</Label>
      <Select
        isRequired
        size="lg"
        className="mb-4"
        placeholder="キャンパスを選択"
      >
        <SelectItem>中目黒・代官山キャンパス</SelectItem>
        <SelectItem>池袋キャンパス</SelectItem>
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
        />
      </I18nProvider>
      <div className="flex gap-4 w-full mb-4">
        <div className="flex flex-col w-full gap-2">
          <Label className="text-sm font-medium text-default-700">
            開始時刻
          </Label>
          <Select isRequired size="lg" placeholder="開始時刻">
            {/* 選択されたキャンパスから生成する */}
            <SelectItem>7:30</SelectItem>
            <SelectItem>8:00</SelectItem>
            <SelectItem>8:30</SelectItem>
            <SelectItem>9:00</SelectItem>
            <SelectItem>9:30</SelectItem>
            <SelectItem>10:00</SelectItem>
            <SelectItem>10:30</SelectItem>
            <SelectItem>11:00</SelectItem>
            <SelectItem>11:30</SelectItem>
            <SelectItem>12:00</SelectItem>
            <SelectItem>12:30</SelectItem>
            <SelectItem>13:00</SelectItem>
            <SelectItem>13:30</SelectItem>
            <SelectItem>14:00</SelectItem>
            <SelectItem>14:30</SelectItem>
            <SelectItem>15:00</SelectItem>
            <SelectItem>15:30</SelectItem>
            <SelectItem>16:00</SelectItem>
            <SelectItem>16:30</SelectItem>
            <SelectItem>17:00</SelectItem>
            <SelectItem>17:30</SelectItem>
            <SelectItem>18:00</SelectItem>
            <SelectItem>18:30</SelectItem>
            <SelectItem>19:00</SelectItem>
            <SelectItem>19:30</SelectItem>
            <SelectItem>20:00</SelectItem>
            <SelectItem>20:30</SelectItem>
            <SelectItem>21:00</SelectItem>
            <SelectItem>21:30</SelectItem>
            <SelectItem>22:00</SelectItem>
          </Select>
        </div>
        <div className="flex flex-col w-full gap-2">
          <Label className="text-sm font-medium text-default-700">
            終了時刻
          </Label>
          <Select isDisabled isRequired size="lg" placeholder="終了時刻">
            <SelectItem>8:00</SelectItem>
            {/* 選択されたキャンパスと開始時刻から自動算出する */}
            <SelectItem>22:30</SelectItem>
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
          size="lg"
          selectionMode="multiple"
          className="mb-4"
          placeholder="希望する条件を選択"
        >
          <SelectItem>グランドピアノ</SelectItem>
          <SelectItem>アップライトピアノ</SelectItem>
          <SelectItem>ピアノ無し</SelectItem>
          <SelectItem>地下希望</SelectItem>
          <SelectItem>4階希望</SelectItem>
        </Select>
      ) : (
        <Select isRequired className="mb-4" size="lg" placeholder="部屋を選択">
          {/* 選択されたキャンパスから生成する */}
          <SelectSection title="地下">
            <SelectItem>A地下103（G）</SelectItem>
            <SelectItem>A地下104（U）</SelectItem>
          </SelectSection>
          <SelectSection title="3階（休日のみ）">
            <SelectItem>A301</SelectItem>
            <SelectItem>A302</SelectItem>
            <SelectItem>A303</SelectItem>
            <SelectItem>A304</SelectItem>
          </SelectSection>
          <SelectSection title="4階">
            <SelectItem>A418（無）</SelectItem>
            <SelectItem>A401（G2台）</SelectItem>
          </SelectSection>
          <SelectSection title="B5F">
            <SelectItem>B503（G）</SelectItem>
          </SelectSection>
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
      />
      <Button isDisabled fullWidth size="lg" color="primary">
        {props.type === "create" ? "予約する" : "変更する"}
      </Button>
    </Form>
  );
}
