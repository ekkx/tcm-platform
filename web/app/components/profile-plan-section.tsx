import { Button, cn } from "@heroui/react";

const plans = [
  {
    name: "LITE",
    price: 1400,
    hours: 30,
    unitPrice: 47,
    priceId: import.meta.env.VITE_STRIPE_PRICE_LITE ?? "",
  },
  {
    name: "STANDARD",
    price: 2400,
    hours: 60,
    unitPrice: 40,
    popular: true,
    priceId: import.meta.env.VITE_STRIPE_PRICE_STANDARD ?? "",
  },
  {
    name: "PRO",
    price: 3300,
    hours: 90,
    unitPrice: 37,
    priceId: import.meta.env.VITE_STRIPE_PRICE_PRO ?? "",
  },
];

const baseUnitPrice = plans[0].unitPrice;

export function ProfilePlanSection({
  currentPlan,
  onSelectPlan,
  onChangePlan,
}: {
  currentPlan: string;
  onSelectPlan?: (priceId: string) => void;
  onChangePlan?: () => void;
}) {
  return (
    <>
      {/* セクションヘッダー */}
      <div className="flex items-center gap-4 px-2">
        <div className="flex-1 h-px bg-default-200" />
        <p className="text-[10px] text-default-400 tracking-[0.2em] font-semibold">
          プランを選択
        </p>
        <div className="flex-1 h-px bg-default-200" />
      </div>

      {/* プランカード */}
      {plans.map((plan) => {
        const isCurrent = plan.name === currentPlan;
        const isDark = isCurrent;
        const savings = Math.round(
          ((baseUnitPrice - plan.unitPrice) / baseUnitPrice) * 100
        );

        return (
          <div key={plan.name} className="relative">
            {/* おすすめバッジ */}
            {plan.popular && (
              <div className="absolute -top-3.5 left-1/2 -translate-x-1/2 z-10">
                <span className="bg-foreground text-background text-xs font-bold tracking-wider px-4 py-1 rounded-full">
                  おすすめ
                </span>
              </div>
            )}

            <div
              className={cn(
                "rounded-3xl p-6 bg-content1",
                plan.popular && "ring-2 ring-foreground"
              )}
            >
              {/* プラン名 */}
              <p
                className={cn(
                  "text-[10px] tracking-[0.15em] font-semibold mb-2 text-default-400"
                  // isDark ? "text-background/50" : "text-default-400"
                )}
              >
                {plan.name}
              </p>

              {/* 価格 */}
              <div className="flex items-baseline mb-4">
                <span className="text-3xl font-bold">
                  ¥{plan.price.toLocaleString()}
                </span>
                <span
                  className={cn(
                    "text-sm ml-0.5 text-default-400"
                    // isDark ? "text-background/50" : "text-default-400"
                  )}
                >
                  /月
                </span>
              </div>

              {/* 詳細 */}
              <div className="flex flex-col gap-2 mb-5">
                <div className="flex items-center gap-2">
                  <svg
                    className={cn(
                      "w-4 h-4 flex-shrink-0 text-foreground"
                      // isDark ? "text-background" : "text-foreground"
                    )}
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
                      strokeWidth="2"
                      d="m5 12l5 5L20 7"
                    />
                  </svg>
                  <span className="text-sm">
                    月{plan.hours}時間まで予約可能
                  </span>
                </div>
                <div className="flex items-center gap-2">
                  <svg
                    className={cn(
                      "w-4 h-4 flex-shrink-0 text-foreground"
                      // isDark ? "text-background" : "text-foreground"
                    )}
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
                      strokeWidth="2"
                      d="m5 12l5 5L20 7"
                    />
                  </svg>
                  <span className="text-sm">
                    1時間あたり約{plan.unitPrice}円
                  </span>
                  {savings > 0 && (
                    <span
                      className={cn(
                        "text-[10px] font-bold px-1.5 py-0.5 rounded bg-success/10 text-success"
                        // isDark
                        //   ? "bg-background/20 text-background"
                        //   : "bg-success/10 text-success"
                      )}
                    >
                      {savings}%おトク
                    </span>
                  )}
                </div>
              </div>

              {/* ボタン */}
              <Button
                fullWidth
                size="lg"
                radius="full"
                className={cn(
                  "font-semibold text-sm bg-default-100 text-foreground"
                )}
                isDisabled={isCurrent}
                onPress={() => {
                  if (isCurrent) return;
                  if (onSelectPlan) {
                    onSelectPlan(plan.priceId);
                  } else if (onChangePlan) {
                    onChangePlan();
                  }
                }}
              >
                {isCurrent
                  ? "現在のプラン"
                  : onChangePlan
                    ? "プランを変更"
                    : "このプランを選択"}
              </Button>
            </div>
          </div>
        );
      })}
    </>
  );
}
