import {
  Button,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  useDisclosure,
} from "@heroui/react";

export function CreateSlaveAccountButton() {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();

  return (
    <>
      <button
        type="button"
        className="flex items-center gap-3 w-full py-3 opacity-40"
        disabled
      >
        <div className="w-8 h-8 rounded-xl bg-success/10 flex items-center justify-center flex-shrink-0">
          <svg
            className="w-4 h-4 text-success"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
          >
            <circle cx="12" cy="6" r="4" fill="currentColor" />
            <path
              fill="currentColor"
              d="M18.095 15.031C17.67 15 17.149 15 16.5 15c-1.65 0-2.475 0-2.987.513C13 16.025 13 16.85 13 18.5c0 1.166 0 1.92.181 2.443Q12.605 21 12 21c-3.866 0-7-1.79-7-4s3.134-4 7-4c2.613 0 4.892.818 6.095 2.031"
              opacity="0.5"
            />
            <path
              fill="currentColor"
              fillRule="evenodd"
              d="M16.5 22c-1.65 0-2.475 0-2.987-.513C13 20.975 13 20.15 13 18.5s0-2.475.513-2.987C14.025 15 14.85 15 16.5 15s2.475 0 2.987.513C20 16.025 20 16.85 20 18.5s0 2.475-.513 2.987C18.975 22 18.15 22 16.5 22m.583-5.056a.583.583 0 1 0-1.166 0v.973h-.973a.583.583 0 1 0 0 1.166h.973v.973a.583.583 0 1 0 1.166 0v-.973h.973a.583.583 0 1 0 0-1.166h-.973z"
              clipRule="evenodd"
            />
          </svg>
        </div>
        <span className="text-sm font-medium">新規追加</span>
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
      <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader>共用アカウントを追加</ModalHeader>
              <ModalBody>
                <Input
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
                  label="ユーザー名"
                  labelPlacement="outside"
                  placeholder="ユーザー名を入力"
                  classNames={{ label: "text-xs opacity-60" }}
                />
                <Input
                  label="パスワード"
                  labelPlacement="outside"
                  placeholder="パスワードを入力"
                  type="password"
                  classNames={{ label: "text-xs opacity-60" }}
                  endContent={
                    <svg
                      className="w-5 h-5 text-default-400 pointer-events-none flex-shrink-0"
                      xmlns="http://www.w3.org/2000/svg"
                      width="24"
                      height="24"
                      viewBox="0 0 24 24"
                    >
                      <path
                        fill="currentColor"
                        fillRule="evenodd"
                        d="M5.25 10.055V8a6.75 6.75 0 0 1 13.5 0v2.055c1.115.083 1.84.293 2.371.824C22 11.757 22 13.172 22 16s0 4.243-.879 5.121C20.243 22 18.828 22 16 22H8c-2.828 0-4.243 0-5.121-.879C2 20.243 2 18.828 2 16s0-4.243.879-5.121c.53-.531 1.256-.741 2.371-.824M6.75 8a5.25 5.25 0 0 1 10.5 0v2.004Q16.676 9.999 16 10H8q-.677-.001-1.25.004zM14 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0"
                        clipRule="evenodd"
                      />
                    </svg>
                  }
                />
              </ModalBody>
              <ModalFooter>
                <Button fullWidth color="primary" onPress={onClose}>
                  追加する
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
}
