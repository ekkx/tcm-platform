import {
  addToast,
  Button,
  Chip,
  Divider,
  Modal,
  ModalBody,
  ModalContent,
  Textarea,
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
  const configMap: Partial<
    Record<ReservationStatus, { color: string; label: string }>
  > = {
    [ReservationStatus.PENDING]: { color: "bg-warning", label: "予約待ち" },
    [ReservationStatus.SUCCESS]: { color: "bg-success", label: "予約確定" },
    [ReservationStatus.FAILED]: { color: "bg-danger", label: "予約失敗" },
  };
  const config = configMap[status];

  if (!config) return null;

  return (
    <Chip
      size="sm"
      classNames={{
        base: "bg-foreground-100 gap-1 px-2",
        content: "text-[10px] font-semibold px-0",
      }}
      startContent={
        <span className={`inline-block w-2 h-2 rounded-full ${config.color}`} />
      }
    >
      {config.label}
    </Chip>
  );
}

export function ReservationListItem({
  reservation,
  onDelete,
  onNoteUpdated,
}: {
  reservation: Reservation;
  onDelete?: (reservationId: string) => void;
  onNoteUpdated?: (reservationId: string, note: string | undefined) => void;
}) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const {
    isOpen: isNoteOpen,
    onOpen: onNoteOpen,
    onOpenChange: onNoteOpenChange,
  } = useDisclosure();
  const [isDeleting, setIsDeleting] = useState(false);
  const [noteText, setNoteText] = useState(reservation.note ?? "");
  const [isSavingNote, setIsSavingNote] = useState(false);

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

  const handleSaveNote = async (onClose: () => void) => {
    setIsSavingNote(true);
    try {
      const trimmed = noteText.trim();
      await reservationClient.updateReservationNote({
        reservationId: reservation.id,
        note: trimmed || undefined,
      });
      addToast({
        title: "メモを保存しました",
        color: "success",
      });
      onNoteUpdated?.(reservation.id, trimmed || undefined);
      onClose();
    } catch {
      addToast({
        title: "メモの保存に失敗しました",
        description: "もう一度お試しください。",
        color: "danger",
      });
    } finally {
      setIsSavingNote(false);
    }
  };

  const handleDelete = async (onClose: () => void) => {
    setIsDeleting(true);
    try {
      await reservationClient.deleteReservation({
        reservationId: reservation.id,
      });
      addToast({
        title: "予約をキャンセルしました",
        color: "success",
      });
      onDelete?.(reservation.id);
    } catch (error) {
      addToast({
        title: "予約のキャンセルに失敗しました",
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
      <div className="grid gap-6 rounded-2xl">
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

        {/* Note */}
        <div className="flex items-start justify-between rounded-xl px-3 pt-2 pb-3 bg-foreground-100">
          <div className="flex-1 min-w-0">
            <p className="text-[10px] text-default-400 mb-1">NOTE</p>
            <p className="text-xs text-default-500 whitespace-pre-wrap">
              {reservation.note || "メモなし"}
            </p>
          </div>
          <Button
            size="sm"
            variant="light"
            className="text-[10px] min-w-0 h-6 px-1 py-2 underline underline-offset-2 text-foreground-600"
            onPress={() => {
              setNoteText(reservation.note ?? "");
              onNoteOpen();
            }}
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
          isDisabled={reservation.status === ReservationStatus.SUCCESS}
          onPress={onOpen}
        >
          キャンセル
        </Button>

        <Modal
          isOpen={isNoteOpen}
          onOpenChange={onNoteOpenChange}
          placement="center"
          size="xs"
          closeButton={<></>}
        >
          <ModalContent>
            {(onClose) => (
              <ModalBody className="p-0 gap-0">
                <div className="grid gap-4 px-3 py-6">
                  <p className="text-xl font-bold text-center">メモを編集</p>
                  <Textarea
                    placeholder="練習内容やメモを入力"
                    value={noteText}
                    onValueChange={setNoteText}
                    maxRows={5}
                    classNames={{ input: "text-base" }}
                  />
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
                    color="primary"
                    variant="flat"
                    isLoading={isSavingNote}
                    onPress={() => handleSaveNote(onClose)}
                  >
                    保存
                  </Button>
                </div>
              </ModalBody>
            )}
          </ModalContent>
        </Modal>

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
                  <p className="text-xl font-bold">
                    予約をキャンセルしますか？
                  </p>
                  <p className="text-xs">
                    この予約をキャンセルしてもよろしいですか？
                  </p>
                </div>
                <Divider />
                <div className="flex justify-center gap-6 py-3">
                  <Button
                    className="w-32 font-bold"
                    variant="flat"
                    onPress={onClose}
                  >
                    閉じる
                  </Button>
                  <Button
                    className="w-32 font-bold"
                    color="danger"
                    variant="flat"
                    isLoading={isDeleting}
                    onPress={() => handleDelete(onClose)}
                  >
                    キャンセル
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
