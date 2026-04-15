import {
  addToast,
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  Modal,
  ModalBody,
  ModalContent,
  useDisclosure,
} from "@heroui/react";
import { ConnectError } from "@connectrpc/connect";
import { useNavigate } from "react-router";
import type { Reservation } from "~/api/pb/reservation/v1/reservation_pb";
import { CampusType } from "~/api/pb/room/v1/room_pb";
import { ReservationForm } from "./reservation-form";
import { ProfilePlanSection } from "./profile-plan-section";
import { subscriptionClient } from "~/api";

export function CreateReservationButton() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const {
    isOpen: isPlanOpen,
    onOpen: onPlanOpen,
    onOpenChange: onPlanOpenChange,
  } = useDisclosure();
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
    if (
      error instanceof ConnectError &&
      error.message.includes("active subscription required")
    ) {
      onPlanOpen();
      return;
    }
    if (
      error instanceof ConnectError &&
      error.message.includes("usage limit exceeded")
    ) {
      addToast({
        title: "今月の利用枠を超えています",
        description:
          "プランの利用時間上限に達しました。上位プランへの変更をご検討ください。",
        color: "warning",
      });
      return;
    }
    addToast({
      title: "予約に失敗しました",
      description: error.message,
      color: "danger",
    });
  };

  const handleSelectPlan = async (priceId: string) => {
    try {
      const res = await subscriptionClient.createCheckoutSession({ priceId });
      if (res.checkoutUrl) {
        window.location.href = res.checkoutUrl;
      }
    } catch {
      addToast({
        title: "チェックアウトの作成に失敗しました",
        color: "danger",
      });
    }
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

      {/* プラン購入モーダル */}
      <Modal
        isOpen={isPlanOpen}
        onOpenChange={onPlanOpenChange}
        size="full"
        classNames={{
          closeButton: "top-4 right-4 scale-125 z-50 bg-default-100",
        }}
      >
        <ModalContent>
          <ModalBody className="flex flex-col gap-8 p-6 pb-10 overflow-y-auto">
            <div className="grid gap-1 pt-4">
              <h3 className="text-xl font-bold">プランを選択</h3>
              <p className="text-xs text-foreground-400">
                予約するにはプランの契約が必要です。
              </p>
            </div>
            <ProfilePlanSection
              currentPlan="FREE"
              onSelectPlan={handleSelectPlan}
            />
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
}
