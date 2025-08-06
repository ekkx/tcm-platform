import {
  addToast,
  Button,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Snippet,
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

  const handleUpdateProfile = async (onClose: () => void) => {
    if (!newUserName.trim()) {
      onClose();
      addToast({
        title: "ユーザー名を入力してください",
        color: "warning",
      });
      return; // ユーザー名が空の場合は何もしない
    }

    try {
      await userClient.updateUser({ displayName: newUserName });
      setUser((prev) => (prev ? { ...prev, displayName: newUserName } : prev));
      onClose();
      addToast({
        title: "ユーザー名を更新しました",
        color: "success",
      });
    } catch (error) {
      addToast({
        title: "ユーザー名の更新に失敗しました",
        description: "もう一度お試しください。",
        color: "danger",
      });
    }
  };

  return (
    <>
      <Button
        disableRipple
        onPress={onOpen}
        className="px-0 gap-3 justify-start bg-transparent"
        startContent={
          <svg
            className="w-5 h-5"
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
        }
        endContent={
          <svg
            className="w-4 h-4 text-foreground-400 ml-auto"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <path
              fill="none"
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="1.5"
              d="m9 5l6 7l-6 7"
            />
          </svg>
        }
      >
        プロフィール
      </Button>
      {user && (
        <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
          <ModalContent>
            {(onClose) => (
              <>
                <ModalHeader>プロフィールを編集</ModalHeader>
                <ModalBody>
                  {/* <div className="grid gap-2"> */}
                  <div className="grid">
                    <span className="text-xs text-foreground-400">
                      ユーザーID
                    </span>
                    <Snippet
                      className="border-none p-0"
                      variant="bordered"
                      symbol={<></>}
                    >
                      {user.id}
                    </Snippet>
                  </div>
                  <Input
                    label="新しいユーザー名"
                    labelPlacement="outside"
                    placeholder="新しいユーザー名を入力"
                    classNames={{ label: "text-xs opacity-60" }}
                    value={newUserName}
                    onChange={(e) => setNewUserName(e.target.value)}
                    endContent={
                      <svg
                        className="w-5 h-5 text-default-400 pointer-events-none flex-shrink-0"
                        xmlns="http://www.w3.org/2000/svg"
                        width="24"
                        height="24"
                        viewBox="0 0 24 24"
                      >
                        <circle cx="12" cy="6" r="4" fill="currentColor" />
                        <path
                          fill="currentColor"
                          d="M20 17.5c0 2.485 0 4.5-8 4.5s-8-2.015-8-4.5S7.582 13 12 13s8 2.015 8 4.5"
                        />
                      </svg>
                    }
                  />
                </ModalBody>
                <ModalFooter>
                  <Button
                    fullWidth
                    color="primary"
                    onPress={() => handleUpdateProfile(onClose)}
                  >
                    保存する
                  </Button>
                </ModalFooter>
              </>
            )}
          </ModalContent>
        </Modal>
      )}
    </>
  );
}
