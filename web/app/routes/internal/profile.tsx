import {
  addToast,
  Avatar,
  Button,
  Chip,
  Divider,
  Modal,
  ModalBody,
  ModalContent,
  Progress,
  useDisclosure,
} from "@heroui/react";
import { UpdateProfileButton } from "~/components/profile/update-profile-button";
import { UpdatePasswordButton } from "~/components/profile/update-password-button";
import { SlaveAccountListButton } from "~/components/profile/slave-account-list-button";
import { useAuth } from "~/providers/auth-provider";
import { Cookie } from "~/store/cookies";

export default function Profile() {
  const { user } = useAuth();
  // TODO: APIから取得する
  const currentPlan = "STANDARD";
  const usedHours = 12.5;
  const totalHours = 60;
  const remainingHours = totalHours - usedHours;
  const usagePercent = Math.round((usedHours / totalHours) * 100);

  const handleCopyId = async () => {
    if (!user?.id) return;
    await navigator.clipboard.writeText(user.id);
    addToast({ title: "IDをコピーしました", color: "success" });
  };

  const handleLogout = () => {
    Cookie.destroy();
    window.location.href = "/";
  };

  return (
    <div className="flex flex-col gap-6 p-6 pb-10">
      {/* プロフィールヘッダー */}
      <div className="flex flex-col items-center gap-3 pt-4 pb-2">
        <Avatar size="lg" className="w-20 h-20" />
        <p className="text-xl font-bold">{user?.displayName}</p>
        <Chip
          size="sm"
          variant="solid"
          color="primary"
          classNames={{
            content:
              "flex items-center gap-1 text-[10px] font-bold tracking-wider px-1",
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            className="w-3 h-3"
          >
            <path
              fill="currentColor"
              fill-rule="evenodd"
              d="M12 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14m0-10c-.284 0-.474.34-.854 1.023l-.098.176c-.108.194-.162.29-.246.354c-.085.064-.19.088-.4.135l-.19.044c-.738.167-1.107.25-1.195.532s.164.577.667 1.165l.13.152c.143.167.215.25.247.354s.021.215 0 .438l-.02.203c-.076.785-.114 1.178.115 1.352c.23.174.576.015 1.267-.303l.178-.082c.197-.09.295-.135.399-.135s.202.045.399.135l.178.082c.691.319 1.037.477 1.267.303s.191-.567.115-1.352l-.02-.203c-.021-.223-.032-.334 0-.438s.104-.187.247-.354l.13-.152c.503-.588.755-.882.667-1.165c-.088-.282-.457-.365-1.195-.532l-.19-.044c-.21-.047-.315-.07-.4-.135c-.084-.064-.138-.16-.246-.354l-.098-.176C12.474 6.34 12.284 6 12 6"
              clip-rule="evenodd"
            />
            <path
              fill="currentColor"
              d="m7.093 15.941l-.379 1.382c-.628 2.292-.942 3.438-.523 4.065c.147.22.344.396.573.513c.652.332 1.66-.193 3.675-1.243c.67-.35 1.006-.524 1.362-.562a2 2 0 0 1 .398 0c.356.038.691.213 1.362.562c2.015 1.05 3.023 1.575 3.675 1.243c.229-.117.426-.293.573-.513c.42-.627.105-1.773-.523-4.065l-.379-1.382A8.46 8.46 0 0 1 12 17.5a8.46 8.46 0 0 1-4.907-1.559"
            />
          </svg>
          <span>{currentPlan}</span>
        </Chip>
        <button
          type="button"
          className="flex items-center gap-1 active:opacity-60 transition-opacity"
          onClick={handleCopyId}
        >
          <span className="text-xs text-default-400">ID: {user?.id}</span>
          <svg
            className="w-3 h-3 text-default-400"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <g fill="none" stroke="currentColor" strokeWidth="1.5">
              <path d="M6 11c0-2.828 0-4.243.879-5.121C7.757 5 9.172 5 12 5h3c2.828 0 4.243 0 5.121.879C21 6.757 21 8.172 21 11v5c0 2.828 0 4.243-.879 5.121C19.243 22 17.828 22 15 22h-3c-2.828 0-4.243 0-5.121-.879C6 20.243 6 18.828 6 16z" />
              <path d="M6 19a3 3 0 0 1-3-3v-6c0-3.771 0-5.657 1.172-6.828S7.229 2 11 2h4a3 3 0 0 1 3 3" />
            </g>
          </svg>
        </button>
      </div>

      {/* 現在の使用状況 */}
      <div className="bg-content1 rounded-3xl p-6">
        <p className="text-[10px] text-default-400 tracking-widest font-semibold mb-4">
          今月の利用状況
        </p>
        <div className="flex items-end justify-between mb-4">
          <div>
            <span className="text-3xl font-bold">{usedHours}</span>
            <span className="text-base text-default-400 font-medium">
              /{totalHours}h
            </span>
          </div>
          <div className="text-right">
            <p className="text-[10px] text-default-400 font-semibold">残り</p>
            <p className="text-xl font-bold">{remainingHours}h</p>
          </div>
        </div>
        <Progress
          size="sm"
          value={usagePercent}
          color={usagePercent >= 80 ? "danger" : "default"}
          classNames={{
            track: "h-2.5 bg-default-100",
            indicator: "bg-foreground",
          }}
        />
        <p className="text-[11px] text-default-400 mt-3 flex items-center gap-1">
          <svg
            className="w-3 h-3"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <path
              fill="currentColor"
              d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2"
              opacity="0.5"
            />
            <path
              fill="currentColor"
              d="M12 7.25a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-1.5 0V8a.75.75 0 0 1 .75-.75M12 16a1 1 0 1 0 0-2a1 1 0 0 0 0 2"
            />
          </svg>
          次の更新日は2026年5月1日です
        </p>
      </div>

      {/* メニュー */}
      <div className="flex flex-col">
        <UpdateProfileButton user={user || undefined} />
        <Divider />
        <UpdatePasswordButton />
        <Divider />
        <SlaveAccountListButton />
        <Divider />
        <MenuLink
          label="サブスクリプション管理"
          color="bg-danger/10"
          iconColor="text-danger"
          icon={
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
            >
              <path
                fill="currentColor"
                fill-rule="evenodd"
                d="M12 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14m0-10c-.284 0-.474.34-.854 1.023l-.098.176c-.108.194-.162.29-.246.354c-.085.064-.19.088-.4.135l-.19.044c-.738.167-1.107.25-1.195.532s.164.577.667 1.165l.13.152c.143.167.215.25.247.354s.021.215 0 .438l-.02.203c-.076.785-.114 1.178.115 1.352c.23.174.576.015 1.267-.303l.178-.082c.197-.09.295-.135.399-.135s.202.045.399.135l.178.082c.691.319 1.037.477 1.267.303s.191-.567.115-1.352l-.02-.203c-.021-.223-.032-.334 0-.438s.104-.187.247-.354l.13-.152c.503-.588.755-.882.667-1.165c-.088-.282-.457-.365-1.195-.532l-.19-.044c-.21-.047-.315-.07-.4-.135c-.084-.064-.138-.16-.246-.354l-.098-.176C12.474 6.34 12.284 6 12 6"
                clip-rule="evenodd"
              />
              <path
                fill="currentColor"
                d="m7.093 15.941l-.379 1.382c-.628 2.292-.942 3.438-.523 4.065c.147.22.344.396.573.513c.652.332 1.66-.193 3.675-1.243c.67-.35 1.006-.524 1.362-.562a2 2 0 0 1 .398 0c.356.038.691.213 1.362.562c2.015 1.05 3.023 1.575 3.675 1.243c.229-.117.426-.293.573-.513c.42-.627.105-1.773-.523-4.065l-.379-1.382A8.46 8.46 0 0 1 12 17.5a8.46 8.46 0 0 1-4.907-1.559"
              />
            </svg>
          }
          onPress={() => {}}
        />
      </div>

      {/* ログアウト */}
      <LogoutButton onLogout={handleLogout} />

      <p className="text-[10px] text-default-300 text-center pb-4">
        VERSION 1.0.0
      </p>
    </div>
  );
}

function LogoutButton({ onLogout }: { onLogout: () => void }) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <button
        type="button"
        className="w-full py-3.5 rounded-full border border-danger text-danger text-sm font-medium active:opacity-60 transition-opacity"
        onClick={onOpen}
      >
        ログアウト
      </button>
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
              <div className="grid gap-2 px-3 py-6 text-center">
                <p className="text-xl font-bold">ログアウトしますか？</p>
                <p className="text-xs text-default-500">
                  再度ログインが必要になります
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
                  onPress={onLogout}
                >
                  ログアウト
                </Button>
              </div>
            </ModalBody>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}

function MenuLink({
  label,
  color,
  iconColor,
  icon,
  bgIcon,
  onPress,
}: {
  label: string;
  color: string;
  iconColor: string;
  icon: React.ReactNode;
  bgIcon?: React.ReactNode;
  onPress: () => void;
}) {
  return (
    <button
      type="button"
      className="flex items-center gap-3 w-full py-3.5 active:opacity-60 transition-opacity"
      onClick={onPress}
    >
      <div
        className={`w-8 h-8 rounded-xl ${color} flex items-center justify-center flex-shrink-0`}
      >
        <svg
          className={`w-4 h-4 ${iconColor}`}
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          viewBox="0 0 24 24"
        >
          {bgIcon}
          {icon}
        </svg>
      </div>
      <span className="text-sm font-medium">{label}</span>
      <svg
        className="w-4 h-4 text-default-300 ml-auto"
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
          d="m9 5l6 7l-6 7"
        />
      </svg>
    </button>
  );
}
