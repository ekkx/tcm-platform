import {
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  useDisclosure,
} from "@heroui/react";
import { ReservationForm } from "./reservation-form";

export function CreateReservationButton() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <Button
        isIconOnly
        className="w-16 h-16 rounded-full bg-foreground/10 backdrop-blur-xl border-[0.5px] border-default-300"
        onPress={onOpen}
        startContent={
          <svg
            className="w-8 h-8 text-white"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <path
              fill="none"
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="1.5"
              d="M5 12h14m-7-7v14"
            />
          </svg>
        }
      />
      <Drawer
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        size="xl"
        classNames={{
          closeButton:
            "top-4 right-4 scale-125 border-[0.5px] border-default-300",
        }}
      >
        <DrawerContent>
          {(onClose) => (
            <>
              <DrawerBody className="gap-8 py-8">
                <div className="grid gap-1">
                  <h3 className="text-xl font-bold">新規予約</h3>
                  <p className="text-xs text-foreground-400">
                    希望の日時と部屋を選択して予約してください。
                  </p>
                </div>
                <ReservationForm />
              </DrawerBody>
            </>
          )}
        </DrawerContent>
      </Drawer>
    </>
  );
}
