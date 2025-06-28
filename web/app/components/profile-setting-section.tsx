import {
  Accordion,
  AccordionItem,
  Avatar,
  Button,
  Card,
  Divider,
  Input,
  Snippet,
} from "@heroui/react";

export function ProfileSettingSection() {
  return (
    <div className="grid gap-2">
      <h4 className="ml-2 text-xs text-foreground-400">設定</h4>
      <Card>
        <Accordion
          isCompact
          variant="shadow"
          className="bg-content1"
          itemClasses={{
            trigger: "py-3",
            title: "text-sm",
            indicator:
              "-rotate-180 data-[open=true]:-rotate-90 rtl:rotate-180 rtl:data-[open=true]:rotate-90",
            content: "px-4 pt-4 pb-10",
          }}
        >
          <AccordionItem
            title="プロフィール"
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
          >
            <div className="grid gap-2">
              <div>
                <span className="text-xs text-foreground-400">ユーザーID</span>
                <Snippet
                  className="border-none p-0"
                  variant="bordered"
                  symbol={<></>}
                >
                  01JYVFTZR1TMGQXX1MZQ9TGDTG
                </Snippet>
              </div>
              <div className="grid gap-2">
                <span className="text-xs text-foreground-400">
                  ユーザーネーム
                </span>
                <div className="flex gap-2">
                  <Input
                    size="sm"
                    name="username"
                    placeholder="ユーザーネームを入力"
                    classNames={{ input: "text-xs" }}
                  />
                  <Button size="sm" variant="solid">
                    保存
                  </Button>
                </div>
              </div>
            </div>
          </AccordionItem>
          <AccordionItem
            title="セキュリティ"
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
                  d="M3 10.417c0-3.198 0-4.797.378-5.335c.377-.537 1.88-1.052 4.887-2.081l.573-.196C10.405 2.268 11.188 2 12 2s1.595.268 3.162.805l.573.196c3.007 1.029 4.51 1.544 4.887 2.081C21 5.62 21 7.22 21 10.417v1.574c0 5.638-4.239 8.375-6.899 9.536C13.38 21.842 13.02 22 12 22s-1.38-.158-2.101-.473C7.239 20.365 3 17.63 3 11.991z"
                  opacity="0.5"
                />
                <path
                  fill="currentColor"
                  d="M14 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 8c4 0 4-.895 4-2s-1.79-2-4-2s-4 .895-4 2s0 2 4 2"
                />
              </svg>
            }
          >
            <div className="grid gap-2 pt-1">
              <div className="grid gap-2">
                <span className="text-xs text-foreground-400">
                  新しいパスワード
                </span>
                <Input
                  size="sm"
                  name="username"
                  placeholder="パスワードを入力"
                  classNames={{ input: "text-xs" }}
                />
              </div>
              <div className="grid gap-2">
                <span className="text-xs text-foreground-400">
                  確認用パスワード
                </span>
                <Input
                  size="sm"
                  name="username"
                  placeholder="パスワードを再入力"
                  classNames={{ input: "text-xs" }}
                />
              </div>
              <Button fullWidth size="sm" className="mt-2">
                保存
              </Button>
            </div>
          </AccordionItem>
          <AccordionItem
            // isDisabled
            title="アカウント管理"
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
                  d="M15.5 7.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0"
                />
                <path
                  fill="currentColor"
                  d="M19.5 7.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m-15 0a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0"
                  opacity="0.4"
                />
                <path
                  fill="currentColor"
                  d="M18 16.5c0 1.933-2.686 3.5-6 3.5s-6-1.567-6-3.5S8.686 13 12 13s6 1.567 6 3.5"
                />
                <path
                  fill="currentColor"
                  d="M22 16.5c0 1.38-1.79 2.5-4 2.5s-4-1.12-4-2.5s1.79-2.5 4-2.5s4 1.12 4 2.5m-20 0C2 17.88 3.79 19 6 19s4-1.12 4-2.5S8.21 14 6 14s-4 1.12-4 2.5"
                  opacity="0.4"
                />
              </svg>
            }
          >
            <div className="grid gap-2 pt-1">
              <div className="grid gap-2">
                <span className="text-xs text-foreground-400">
                  共用アカウント
                </span>
                <div className="bg-content2 rounded-2xl">
                  <div className="flex items-center gap-2 p-2">
                    <svg
                      className="w-8 h-8"
                      xmlns="http://www.w3.org/2000/svg"
                      width="24"
                      height="24"
                      viewBox="0 0 24 24"
                    >
                      <path
                        fill="currentColor"
                        d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"
                        opacity="0.5"
                      />
                      <path
                        fill="currentColor"
                        d="M12.75 9a.75.75 0 0 0-1.5 0v2.25H9a.75.75 0 0 0 0 1.5h2.25V15a.75.75 0 0 0 1.5 0v-2.25H15a.75.75 0 0 0 0-1.5h-2.25z"
                      />
                    </svg>
                    <div>
                      <span className="text-sm">新規</span>
                    </div>
                    <Button size="sm" variant="flat" className="ml-auto">
                      追加
                    </Button>
                  </div>
                  <Divider />
                  <div className="flex items-center gap-2 p-2">
                    <Avatar size="sm" />
                    <div className="flex flex-col">
                      <span className="text-sm">ゆうすけ</span>
                      <span className="text-xs text-foreground-400">
                        作成日: 2025/5/28
                      </span>
                    </div>
                    {/* TODO: ドロップダウンメニューに変更 */}
                    <Button
                      size="sm"
                      color="danger"
                      variant="flat"
                      className="ml-auto"
                    >
                      削除
                    </Button>
                  </div>
                  <Divider />
                  <div className="flex items-center gap-2 p-2">
                    <Avatar size="sm" />
                    <div className="flex flex-col">
                      <span className="text-sm">間宮</span>
                      <span className="text-xs text-foreground-400">
                        作成日: 2025/5/28
                      </span>
                    </div>
                    {/* TODO: ドロップダウンメニューに変更 */}
                    <Button
                      size="sm"
                      color="danger"
                      variant="flat"
                      className="ml-auto"
                    >
                      削除
                    </Button>
                  </div>
                </div>
              </div>
            </div>
          </AccordionItem>
        </Accordion>
      </Card>
    </div>
  );
}
