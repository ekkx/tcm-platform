import { Avatar } from "@heroui/react";
import { ProfileLoginSection } from "~/components/profile-login-section";
import { ProfileSettingSection } from "~/components/profile-setting-section";

export default function Profile() {
  return (
    <div className="grid gap-6 px-4 pt-6">
      <div className="flex flex-col gap-3 items-center mx-auto pt-8">
        <Avatar size="lg" />
        <div className="flex flex-col gap-1 items-center">
          <strong>ねこぱんち</strong>
          <small className="text-foreground-400">
            <span>ID: </span>01JYVFTZR1TMGQXX1MZQ9TGDTG
          </small>
        </div>
      </div>
      <ProfileSettingSection />
      <ProfileLoginSection />
    </div>
  );
}
