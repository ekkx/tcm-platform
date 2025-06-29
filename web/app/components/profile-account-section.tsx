import { Card, Divider } from "@heroui/react";
import { CreateSlaveAccountButton } from "./profile/create-slave-account-button";
import { SlaveAccountListButton } from "./profile/slave-account-list-button";

export function ProfileAccountSection() {
  return (
    <div className="grid gap-2">
      <h4 className="ml-2 text-xs text-foreground-400">共用アカウント</h4>
      <Card className="px-4">
        <CreateSlaveAccountButton />
        <Divider />
        <SlaveAccountListButton />
      </Card>
    </div>
  );
}
