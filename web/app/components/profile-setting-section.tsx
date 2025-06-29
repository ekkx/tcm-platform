import { Card, Divider } from "@heroui/react";
import { UpdatePasswordButton } from "./profile/update-password-button";
import { UpdateProfileButton } from "./profile/update-profile-button";

export function ProfileSettingSection() {
  return (
    <div className="grid gap-2">
      <h4 className="ml-2 text-xs text-foreground-400">設定</h4>
      <Card className="px-4">
        <UpdateProfileButton />
        <Divider />
        <UpdatePasswordButton />
      </Card>
    </div>
  );
}
