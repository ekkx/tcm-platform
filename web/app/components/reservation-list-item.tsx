import {
  addToast,
  Button,
  Card,
  CardBody,
  Chip,
  Divider,
  Modal,
  ModalBody,
  ModalContent,
  useDisclosure,
} from "@heroui/react";
import { useState } from "react";
import { reservationClient } from "~/api";
import {
  ReservationStatus,
  type Reservation,
} from "~/api/pb/reservation/v1/reservation_pb";
import { CampusType, PianoType } from "~/api/pb/room/v1/room_pb";

function formatTime(hour: number, minute: number) {
  return `${hour}:${String(minute).padStart(2, "0")}`;
}

function formatDate(dateStr: string) {
  const date = new Date(`${dateStr}T00:00:00+09:00`);
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const daysOfWeek = ["日", "月", "火", "水", "木", "金", "土"];
  const weekday = daysOfWeek[date.getDay()];
  return `${month}月${day}日（${weekday}）`;
}

function getCampusLabel(campusType: CampusType) {
  switch (campusType) {
    case CampusType.NAKAMEGURO:
      return "中目黒・代官山キャンパス";
    case CampusType.IKEBUKURO:
      return "池袋キャンパス";
    default:
      return "Unknown Campus";
  }
}

function getPianoLabel(pianoType: PianoType) {
  switch (pianoType) {
    case PianoType.GRAND:
      return "グランドピアノ";
    case PianoType.UPRIGHT:
      return "アップライトピアノ";
    case PianoType.NONE:
      return "ピアノ無";
    default:
      return "";
  }
}

function StatusChip({ status }: { status: ReservationStatus }) {
  switch (status) {
    case ReservationStatus.PENDING:
      return (
        <Chip
          size="sm"
          variant="flat"
          color="warning"
          classNames={{ content: "text-[10px] font-semibold" }}
        >
          予約待ち
        </Chip>
      );
    case ReservationStatus.SUCCESS:
      return (
        <Chip
          size="sm"
          variant="flat"
          color="success"
          classNames={{ content: "text-[10px] font-semibold" }}
        >
          予約確定
        </Chip>
      );
    case ReservationStatus.FAILED:
      return (
        <Chip
          size="sm"
          variant="flat"
          color="danger"
          classNames={{ content: "text-[10px] font-semibold" }}
        >
          予約失敗
        </Chip>
      );
    default:
      return null;
  }
}

export function ReservationListItem({
  reservation,
  onDelete,
}: {
  reservation: Reservation;
  onDelete?: (reservationId: string) => void;
}) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const [isDeleting, setIsDeleting] = useState(false);

  const timeRange = `${formatTime(
    reservation.fromHour,
    reservation.fromMinute
  )} - ${formatTime(reservation.toHour, reservation.toMinute)}`;
  const dateLabel = formatDate(reservation.date);
  const campusLabel = getCampusLabel(reservation.campusType);
  const roomName = reservation.room?.name ?? "未設定";
  const pianoLabel = reservation.room
    ? getPianoLabel(reservation.room.pianoType)
    : "";

  const handleDelete = async (onClose: () => void) => {
    setIsDeleting(true);
    try {
      await reservationClient.deleteReservation({
        reservationId: reservation.id,
      });
      addToast({
        title: "予約を削除しました",
        color: "success",
      });
      onDelete?.(reservation.id);
    } catch (error) {
      addToast({
        title: "予約の削除に失敗しました",
        description: "もう一度お試しください。",
        color: "danger",
      });
    } finally {
      setIsDeleting(false);
      onClose();
    }
  };

  return (
    <div className="bg-content1 rounded-3xl p-6">
      <div className="grid gap-3 rounded-2xl">
        {/* Header: Date + Time + Status */}
        <div className="flex items-start justify-between">
          <div className="grid gap-1">
            <p className="text-xs text-default-400">{dateLabel}</p>
            <p className="text-xl font-bold leading-tight">{timeRange}</p>
          </div>
          <StatusChip status={reservation.status} />
        </div>

        {/* Location + Instrument */}
        <div className="flex gap-4 items-center">
          <div className="w-1/2">
            <p className="text-[10px] text-default-400 mb-1">LOCATION</p>
            <p className="text-base font-semibold">{roomName}</p>
            <p className="text-[10px] text-default-500">{campusLabel}</p>
          </div>
          <Divider orientation="vertical" className="h-auto self-stretch" />
          <div className="w-1/2">
            <p className="text-[10px] text-default-400 mb-1">INSTRUMENT</p>
            <p className="text-base font-semibold">{pianoLabel || "-"}</p>
          </div>
        </div>

        {/* Memo */}
        <div className="flex items-start justify-between rounded-xl px-3 pt-2 pb-3 bg-foreground-100">
          <div className="flex-1 min-w-0">
            <p className="text-[10px] text-default-400 mb-1">NOTE</p>
            <p className="text-xs text-default-500">メモなし</p>
          </div>
          <Button
            size="sm"
            variant="light"
            className="text-[10px] min-w-0 h-6 px-1 py-2 underline underline-offset-2"
            isDisabled
          >
            メモを編集
          </Button>
        </div>

        {/* Delete Button */}
        <Button
          fullWidth
          variant="bordered"
          color="default"
          size="lg"
          className="border-1 rounded-3xl text-sm"
          onPress={onOpen}
        >
          キャンセル
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
              <ModalBody className="p-0 gap-0">
                <div className="grid gap-4 px-3 py-6 text-center">
                  <p className="text-xl font-bold">予約を削除しますか？</p>
                  <p className="text-xs">
                    この予約を削除してもよろしいですか？
                  </p>
                </div>
                <Divider />
                <div className="flex justify-center gap-6 py-3">
                  <Button
                    className="w-32 font-bold"
                    variant="flat"
                    onPress={onClose}
                  >
                    キャンセル
                  </Button>
                  <Button
                    className="w-32 font-bold"
                    color="danger"
                    variant="flat"
                    isLoading={isDeleting}
                    onPress={() => handleDelete(onClose)}
                  >
                    削除
                  </Button>
                </div>
              </ModalBody>
            )}
          </ModalContent>
        </Modal>
      </div>
    </div>
  );
}
