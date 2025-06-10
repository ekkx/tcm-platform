import {
  addToast,
  Button,
  Modal,
  ModalBody,
  ModalContent,
  useDisclosure,
} from "@heroui/react";
import { Form, useNavigation, useActionData } from "react-router";
import { useEffect } from "react";

type Props = {
  reservationId: number;
  onDelete?: () => void;
};

export function CancelReservationButton(props: Props) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const navigation = useNavigation();
  const actionData = useActionData() as any;
  const isDeleting = navigation.state === "submitting";

  useEffect(() => {
    if (actionData?.deleted) {
      addToast({
        title: "予約を削除しました",
        color: "success",
      });
      props.onDelete?.();
      onOpenChange();
      // Reload the page after successful deletion
      window.location.reload();
    } else if (actionData?.error) {
      addToast({
        title: actionData.error,
        color: "danger",
      });
    }
  }, [actionData, props, onOpenChange]);

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
                <Form method="post">
                  <input type="hidden" name="intent" value="delete-reservation" />
                  <input type="hidden" name="reservation_id" value={props.reservationId} />
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
                      type="submit"
                      isLoading={isDeleting}
                    >
                      削除
                    </Button>
                  </div>
                </Form>
              </ModalBody>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}
