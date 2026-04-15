import { addToast } from "@heroui/react";
import { ConnectError } from "@connectrpc/connect";
import { useEffect, useState } from "react";
import { subscriptionClient } from "~/api";
import {
  PlanType,
  type Subscription,
} from "~/api/pb/subscription/v1/subscription_pb";
import { ProfilePlanSection } from "~/components/profile-plan-section";

const planLabel: Record<number, string> = {
  [PlanType.UNLIMITED]: "UNLIMITED",
  [PlanType.LITE]: "LITE",
  [PlanType.STANDARD]: "STANDARD",
  [PlanType.PRO]: "PRO",
};

export default function Plans() {
  const [subscription, setSubscription] = useState<Subscription | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    (async () => {
      try {
        const res = await subscriptionClient.getSubscription({});
        setSubscription(res.subscription ?? null);
      } catch {
        // サブスクリプション未作成
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  const currentPlan = subscription
    ? (planLabel[subscription.plan] ?? "FREE")
    : "FREE";
  const isUnlimited = subscription?.plan === PlanType.UNLIMITED;
  const hasActiveSubscription =
    subscription != null && subscription.status === "active";

  const handleSelectPlan = async (priceId: string) => {
    try {
      const res = await subscriptionClient.createCheckoutSession({ priceId });
      if (res.checkoutUrl) {
        window.location.href = res.checkoutUrl;
      }
    } catch (error) {
      // 既に契約中の場合は Customer Portal へリダイレクト
      if (
        error instanceof ConnectError &&
        error.message.includes("already subscribed")
      ) {
        handleOpenPortal();
        return;
      }
      addToast({
        title: "チェックアウトの作成に失敗しました",
        color: "danger",
      });
    }
  };

  const handleOpenPortal = async () => {
    try {
      const res = await subscriptionClient.createPortalSession({});
      window.location.href = res.portalUrl;
    } catch {
      addToast({
        title: "ポータルの表示に失敗しました",
        color: "danger",
      });
    }
  };

  if (loading) {
    return <div className="flex flex-col gap-8 p-6 pb-10" />;
  }

  if (isUnlimited) {
    return (
      <div className="flex flex-col gap-8 p-6 pb-10">
        <div className="text-center py-12">
          <p className="text-lg font-bold mb-2">UNLIMITED</p>
          <p className="text-sm text-default-400">
            無制限プランをご利用中です
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-8 p-6 pb-10">
      <ProfilePlanSection
        currentPlan={currentPlan}
        onSelectPlan={hasActiveSubscription ? undefined : handleSelectPlan}
        onChangePlan={hasActiveSubscription ? handleOpenPortal : undefined}
      />
    </div>
  );
}
