import {
  addToast,
  Avatar,
  Button,
  Divider,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  useDisclosure,
} from "@heroui/react";
import { useState } from "react";
import { userClient } from "~/api";
import type { User } from "~/api/pb/user/v1/user_pb";
import { useAuth } from "~/providers/auth-provider";

export function UpdateProfileButton({ user }: { user?: User }) {
  const { setUser } = useAuth();
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  const [newUserName, setNewUserName] = useState("");
  const [isSaving, setIsSaving] = useState(false);

  const handleOpen = () => {
    setNewUserName(user?.displayName ?? "");
    onOpen();
  };

  const handleSave = async (onClose: () => void) => {
    if (!newUserName.trim()) {
      addToast({ title: "ユーザー名を入力してください", color: "warning" });
      return;
    }

    setIsSaving(true);
    try {
      await userClient.updateUser({ displayName: newUserName });
      setUser((prev) => (prev ? { ...prev, displayName: newUserName } : prev));
      onClose();
      addToast({ title: "ユーザー名を更新しました", color: "success" });
    } catch {
      addToast({
        title: "ユーザー名の更新に失敗しました",
        description: "もう一度お試しください。",
        color: "danger",
      });
    } finally {
      setIsSaving(false);
    }
  };

  return (
    <>
      <button
        type="button"
        className="flex items-center gap-3 w-full py-3.5 active:opacity-60 transition-opacity"
        onClick={handleOpen}
      >
        <div className="w-8 h-8 rounded-xl bg-primary/10 flex items-center justify-center flex-shrink-0">
          <svg
            className="w-4 h-4 text-primary"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <path
              fill="currentColor"
              d="M14 4h-4C6.229 4 4.343 4 3.172 5.172S2 8.229 2 12s0 5.657 1.172 6.828S6.229 20 10 20h4c3.771 0 5.657 0 6.828-1.172S22 15.771 22 12s0-5.657-1.172-6.828S17.771 4 14 4"
              opacity="0.5"
            />
            <path
              fill="currentColor"
              d="M13.25 9a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 0 1.5h-5a.75.75 0 0 1-.75-.75m1 3a.75.75 0 0 1 .75-.75h4a.75.75 0 0 1 0 1.5h-4a.75.75 0 0 1-.75-.75m1 3a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1-.75-.75M9 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 6c4 0 4-.895 4-2s-1.79-2-4-2s-4 .895-4 2s0 2 4 2"
            />
          </svg>
        </div>
        <span className="text-sm font-medium">プロフィール編集</span>
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

      {user && (
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
                <div className="grid gap-4 px-6 py-6">
                  <div className="flex flex-col items-center gap-2">
                    <Avatar size="lg" className="w-16 h-16" />
                    <p className="text-xs text-default-400">ID: {user.id}</p>
                  </div>
                  <Input
                    fullWidth
                    label="ユーザー名"
                    labelPlacement="outside"
                    placeholder="ユーザー名を入力"
                    classNames={{ label: "text-xs opacity-60" }}
                    value={newUserName}
                    onChange={(e) => setNewUserName(e.target.value)}
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
                    isLoading={isSaving}
                    onPress={() => handleSave(onClose)}
                  >
                    保存
                  </Button>
                </div>
              </ModalBody>
            )}
          </ModalContent>
        </Modal>
      )}
    </>
  );
}
