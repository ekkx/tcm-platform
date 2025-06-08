import {
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  useDisclosure,
} from "@heroui/react";
import { ReservationForm } from "./reservation-form";

type Props = {
  rooms: any[];
  isConfirmed?: boolean;
};

export function CreateReservationButton(props: Props) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <Button
        isIconOnly
        color="primary"
        className="w-12 h-12 rounded-full"
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
              strokeWidth="2"
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
          closeButton: "top-4 right-4 scale-125 border",
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
                <ReservationForm type="create" rooms={props.rooms} />
              </DrawerBody>
            </>
          )}
        </DrawerContent>
      </Drawer>
    </>
  );
}
