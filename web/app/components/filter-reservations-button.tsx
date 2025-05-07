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
import { today } from "@internationalized/date";
import { I18nProvider, Label } from "react-aria-components";

type Props = {
  isConfirmed?: boolean;
};

export function FilterReservationsButton(props: Props) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

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
            <g fill="none" stroke="currentColor" stroke-width="1.5">
              <circle cx="11.5" cy="11.5" r="9.5" />
              <path stroke-linecap="round" d="M18.5 18.5L22 22" />
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
                    isRequired
                    className="mb-2"
                    placeholder="キャンパスを選択"
                    size="lg"
                  >
                    <SelectItem>中目黒・代官山キャンパス</SelectItem>
                    <SelectItem>池袋キャンパス</SelectItem>
                  </Select>
                  <Label className="text-sm font-medium text-default-700">
                    ピアノ
                  </Label>
                  <Select
                    isRequired
                    className="mb-2"
                    placeholder="ピアノを選択"
                    size="lg"
                  >
                    <SelectItem>グランドピアノ</SelectItem>
                    <SelectItem>アップライトピアノ</SelectItem>
                    <SelectItem>ピアノなし</SelectItem>
                  </Select>
                  <Label className="text-sm font-medium text-default-700">
                    予約日
                  </Label>
                  <I18nProvider locale="ja">
                    <DatePicker
                      isRequired
                      fullWidth
                      className="mb-2"
                      size="lg"
                      defaultValue={today("Asia/Tokyo")}
                      minValue={today("Asia/Tokyo")}
                    />
                  </I18nProvider>
                  <div className="flex gap-2 w-full">
                    <Button fullWidth color="primary" size="lg">
                      絞り込む
                    </Button>
                    <Button variant="flat" size="lg">
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
