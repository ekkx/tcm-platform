import {
  Button,
  Modal,
  ModalBody,
  ModalContent,
  useDisclosure,
} from "@heroui/react";

export function CancelReservationButton() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <Button
        onPress={onOpen}
        fullWidth
        size="sm"
        color="danger"
        variant="flat"
      >
        削除
      </Button>
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
                  <p className="text-xl font-bold">予約を削除しますか？</p>
                  <p className="text-xs">
                    この予約を削除してもよろしいですか？
                  </p>
                </div>
                <div className="flex justify-center gap-6 border-t py-3">
                  <Button
                    className="w-32 font-bold border"
                    variant="bordered"
                    onPress={onClose}
                  >
                    キャンセル
                  </Button>
                  <Button
                    className="w-32 font-bold"
                    color="danger"
                    onPress={onClose}
                  >
                    削除
                  </Button>
                </div>
              </ModalBody>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}
