import { Avatar } from "@heroui/react";
import { ProfileAccountSection } from "~/components/profile-account-section";
import { ProfileLoginSection } from "~/components/profile-login-section";
import { ProfileSettingSection } from "~/components/profile-setting-section";
import { useAuth } from "~/providers/auth-provider";

export default function Profile() {
  const { user } = useAuth();

  return (
    <div className="grid gap-6 px-4 pt-6">
      <div className="flex flex-col gap-3 items-center mx-auto pt-8">
        <Avatar size="lg" />
        <div className="flex flex-col gap-1 items-center">
          <strong>{user?.displayName}</strong>
          <small className="text-foreground-400">
            <span>ID: </span>
            {user?.id}
          </small>
        </div>
      </div>
      <ProfileSettingSection user={user || undefined} />
      <ProfileAccountSection />
      <ProfileLoginSection />
    </div>
  );
}
