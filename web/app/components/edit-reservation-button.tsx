import {
  Button,
  Drawer,
  DrawerBody,
  DrawerContent,
  Modal,
  ModalBody,
  ModalContent,
  useDisclosure,
} from "@heroui/react";
import { ReservationForm } from "./reservation-form";

type Props = {
  isConfirmed?: boolean;
  rooms?: any[];
  reservationId: number;
  campusType: string;
  date: string;
  timeRange: string;
  roomName: string;
  roomId?: string;
  userName?: string;
};

export function EditReservationButton(props: Props) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <Button fullWidth onPress={onOpen} size="sm" variant="flat">
        編集
      </Button>
      {props.isConfirmed ? (
        <Modal
          isOpen={isOpen}
          onOpenChange={onOpenChange}
          placement="center"
          size="xs"
          closeButton={<></>}
        >
          <ModalContent>
            {(onClose) => (
              <>
                <ModalBody className="p-0 gap-0">
                  <div className="grid gap-4 py-6 text-center">
                    <p className="text-xl font-bold">予約を変更できません</p>
                    <p className="text-xs">
                      この予約は確定しているため、変更できません。
                    </p>
                  </div>
                  <div className="flex justify-center gap-6 border-t py-3">
                    <Button
                      className="w-32 font-bold border"
                      variant="bordered"
                      onPress={onClose}
                    >
                      閉じる
                    </Button>
                  </div>
                </ModalBody>
              </>
            )}
          </ModalContent>
        </Modal>
      ) : (
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
                    <h3 className="text-xl font-bold">予約変更</h3>
                    <p className="text-xs text-foreground-400">
                      希望するセクションを変更してください。
                    </p>
                  </div>
                  <ReservationForm 
                    type="update" 
                    rooms={props.rooms || []} 
                    reservationId={props.reservationId}
                    defaultCampus={props.campusType === "池袋" ? "ikebukuro" : "nakameguro"}
                    defaultDate={props.date}
                    defaultStartTime={props.timeRange.split(" 〜 ")[0]}
                    defaultEndTime={props.timeRange.split(" 〜 ")[1]}
                    defaultRoomName={props.roomId}
                    defaultUserName={props.userName}
                    onSuccess={onClose}
                  />
                </DrawerBody>
              </>
            )}
          </DrawerContent>
        </Drawer>
      )}
    </>
  );
}
