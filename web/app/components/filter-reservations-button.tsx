import {
  Button,
  DatePicker,
  Drawer,
  DrawerBody,
  DrawerContent,
  Form,
  Select,
  SelectItem,
  useDisclosure,
} from "@heroui/react";
import { today, CalendarDate } from "@internationalized/date";
import { I18nProvider, Label } from "react-aria-components";
import { useState } from "react";

type Props = {
  onFilterChange: (filters: {
    campus?: string;
    pianoType?: string;
    date?: string;
  }) => void;
};

export function FilterReservationsButton(props: Props) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const [campus, setCampus] = useState<string>("");
  const [pianoType, setPianoType] = useState<string>("");
  const [selectedDate, setSelectedDate] = useState<CalendarDate | null>(null);

  const handleFilter = () => {
    const filters: { campus?: string; pianoType?: string; date?: string } = {};
    
    if (campus) filters.campus = campus;
    if (pianoType) filters.pianoType = pianoType;
    if (selectedDate) {
      // Create date string directly from CalendarDate components
      const year = selectedDate.year;
      const month = String(selectedDate.month).padStart(2, '0');
      const day = String(selectedDate.day).padStart(2, '0');
      filters.date = `${year}-${month}-${day}`;
    }
    
    props.onFilterChange(filters);
    onOpenChange();
  };

  const handleReset = () => {
    setCampus("");
    setPianoType("");
    setSelectedDate(null);
    props.onFilterChange({});
  };

  return (
    <>
      <Button
        isIconOnly
        variant="light"
        size="lg"
        className="rounded-full"
        onPress={onOpen}
        startContent={
          <svg
            className="w-6 h-6 text-foreground-500"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <g fill="none" stroke="currentColor" strokeWidth="1.5">
              <circle cx="11.5" cy="11.5" r="9.5" />
              <path strokeLinecap="round" d="M18.5 18.5L22 22" />
            </g>
          </svg>
        }
      />
      <Drawer
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        placement="bottom"
        size="xl"
        classNames={{
          closeButton: "top-4 right-4 scale-125 border",
        }}
      >
        <DrawerContent>
          {(onClose) => (
            <>
              <DrawerBody className="gap-8 py-8">
                <div className="grid gap-1">
                  <h3 className="text-xl font-bold">絞り込み</h3>
                  <p className="text-xs text-foreground-400">
                    絞り込み条件を入力してください。
                  </p>
                </div>
                <Form>
                  <Label className="text-sm font-medium text-default-700">
                    キャンパス
                  </Label>
                  <Select
                    className="mb-2"
                    placeholder="キャンパスを選択"
                    size="lg"
                    selectedKeys={campus ? [campus] : []}
                    onSelectionChange={(keys) => setCampus(Array.from(keys)[0] as string)}
                  >
                    <SelectItem key="nakameguro">中目黒・代官山キャンパス</SelectItem>
                    <SelectItem key="ikebukuro">池袋キャンパス</SelectItem>
                  </Select>
                  <Label className="text-sm font-medium text-default-700">
                    ピアノ
                  </Label>
                  <Select
                    className="mb-2"
                    placeholder="ピアノを選択"
                    size="lg"
                    selectedKeys={pianoType ? [pianoType] : []}
                    onSelectionChange={(keys) => setPianoType(Array.from(keys)[0] as string)}
                  >
                    <SelectItem key="grand">グランドピアノ</SelectItem>
                    <SelectItem key="upright">アップライトピアノ</SelectItem>
                    <SelectItem key="none">ピアノなし</SelectItem>
                  </Select>
                  <Label className="text-sm font-medium text-default-700">
                    予約日
                  </Label>
                  <I18nProvider locale="ja">
                    <DatePicker
                      fullWidth
                      className="mb-2"
                      size="lg"
                      value={selectedDate}
                      onChange={setSelectedDate}
                      minValue={today("Asia/Tokyo")}
                    />
                  </I18nProvider>
                  <div className="flex gap-2 w-full">
                    <Button fullWidth color="primary" size="lg" onPress={handleFilter}>
                      絞り込む
                    </Button>
                    <Button variant="flat" size="lg" onPress={handleReset}>
                      リセット
                    </Button>
                  </div>
                </Form>
              </DrawerBody>
            </>
          )}
        </DrawerContent>
      </Drawer>
    </>
  );
}
