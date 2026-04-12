import { ProfilePlanSection } from "~/components/profile-plan-section";

export default function Plans() {
  // TODO: APIから取得する
  const currentPlan = "STANDARD";

  return (
    <div className="flex flex-col gap-8 p-6 pb-10">
      <ProfilePlanSection currentPlan={currentPlan} />
    </div>
  );
}
