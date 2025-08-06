import { Card, Divider } from "@heroui/react";
import type { User } from "~/api/pb/user/v1/user_pb";
import { UpdatePasswordButton } from "./profile/update-password-button";
import { UpdateProfileButton } from "./profile/update-profile-button";

export function ProfileSettingSection({ user }: { user?: User }) {
  return (
    <div className="grid gap-2">
      <h4 className="ml-2 text-xs text-foreground-400">設定</h4>
      <Card className="px-4">
        <UpdateProfileButton user={user} />
        <Divider />
        <UpdatePasswordButton />
      </Card>
    </div>
  );
}
