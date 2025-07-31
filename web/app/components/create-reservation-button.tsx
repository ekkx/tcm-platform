import {
  addToast,
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  useDisclosure,
} from "@heroui/react";
import { useNavigate } from "react-router";
import type { Reservation } from "~/api/pb/reservation/v1/reservation_pb";
import { CampusType } from "~/api/pb/room/v1/room_pb";
import { ReservationForm } from "./reservation-form";

export function CreateReservationButton() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const navigate = useNavigate();

  const onReservationCreated = (
    reservation: Reservation,
    onClose: () => void
  ) => {
    const campusName =
      reservation.campusType === CampusType.NAKAMEGURO ? "中目黒" : "池袋";
    addToast({
      title: "予約が完了しました",
      description: `【${campusName}キャンパス】${reservation.date}`,
      color: "success",
    });
    onClose();
    navigate("/home", { state: { newReservation: reservation } });
  };

  const onReservationFailed = (error: Error) => {
    addToast({
      title: "予約に失敗しました",
      description: error.message,
      color: "danger",
    });
  };

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
          closeButton: "top-4 right-4 scale-125 z-50 bg-default-100",
        }}
      >
        <DrawerContent>
          {(onClose) => (
            <>
              <DrawerBody className="gap-8 p-0 pt-8">
                <div className="grid gap-1 px-6">
                  <h3 className="text-xl font-bold">新規予約</h3>
                  <p className="text-xs text-foreground-400">
                    希望の日時と部屋を選択して予約してください。
                  </p>
                </div>
                <ReservationForm
                  onReservationCreated={(reservation) =>
                    onReservationCreated(reservation, onClose)
                  }
                  onReservationFailed={onReservationFailed}
                />
              </DrawerBody>
            </>
          )}
        </DrawerContent>
      </Drawer>
    </>
  );
}
