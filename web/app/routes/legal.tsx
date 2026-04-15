import { Card, CardBody, Chip, Divider, Link } from "@heroui/react";

const plans = [
  { name: "LITE", price: "¥1,400", hours: "30時間" },
  { name: "STANDARD", price: "¥2,400", hours: "60時間" },
  { name: "PRO", price: "¥3,300", hours: "90時間" },
];

export default function Legal() {
  return (
    <div className="fixed inset-0 overflow-y-auto bg-background">
      <div className="w-full max-w-lg mx-auto px-6 py-12 pb-20">
        {/* ヘッダー */}
        <div className="mb-10">
          <p className="text-[10px] tracking-[0.2em] text-default-400 font-semibold mb-2">
            COMMERCE DISCLOSURE
          </p>
          <h1 className="text-2xl font-bold">特定商取引法に基づく表記</h1>
        </div>

        {/* 事業者情報 */}
        <section className="mb-8">
          <h2 className="text-xs tracking-[0.15em] text-default-400 font-semibold mb-4">
            事業者情報
          </h2>
          <Card className="bg-content1">
            <CardBody className="gap-4 p-5">
              <Row label="販売事業者" value="個人運営" />
              <Divider />
              <Row
                label="運営責任者"
                value="請求があった場合に遅滞なく開示いたします"
              />
              <Divider />
              <Row
                label="所在地"
                value="請求があった場合に遅滞なく開示いたします"
              />
              <Divider />
              <Row label="お問い合わせ">
                <Link href="mailto:nikola.desuga@gmail.com" size="sm">
                  nikola.desuga@gmail.com
                </Link>
              </Row>
            </CardBody>
          </Card>
        </section>

        {/* サービス内容 */}
        <section className="mb-8">
          <h2 className="text-xs tracking-[0.15em] text-default-400 font-semibold mb-4">
            サービス内容
          </h2>
          <Card className="bg-content1">
            <CardBody className="p-5">
              <p className="text-sm leading-relaxed">
                東京音楽大学の学生向け練習室予約サービス（非公式）。月額サブスクリプションにより、練習室の予約枠を提供します。
              </p>
            </CardBody>
          </Card>
        </section>

        {/* 料金プラン */}
        <section className="mb-8">
          <h2 className="text-xs tracking-[0.15em] text-default-400 font-semibold mb-4">
            販売価格（税込）
          </h2>
          <div className="grid gap-3">
            {plans.map((plan) => (
              <Card key={plan.name} className="bg-content1">
                <CardBody className="flex-row items-center justify-between p-4">
                  <div className="flex items-center gap-3">
                    <Chip
                      size="sm"
                      variant="flat"
                      className="font-bold text-[10px] tracking-wider"
                    >
                      {plan.name}
                    </Chip>
                    <span className="text-sm text-default-500">
                      月{plan.hours}まで
                    </span>
                  </div>
                  <span className="text-sm font-bold">{plan.price}/月</span>
                </CardBody>
              </Card>
            ))}
          </div>
        </section>

        {/* 支払い・提供 */}
        <section className="mb-8">
          <h2 className="text-xs tracking-[0.15em] text-default-400 font-semibold mb-4">
            支払い・提供について
          </h2>
          <Card className="bg-content1">
            <CardBody className="gap-4 p-5">
              <Row label="支払い方法" value="クレジットカード（Stripe経由）" />
              <Divider />
              <Row
                label="支払い時期"
                value="契約時に初回決済が行われ、以降は毎月自動で課金されます"
              />
              <Divider />
              <Row
                label="サービス提供時期"
                value="決済完了後、直ちにご利用いただけます"
              />
            </CardBody>
          </Card>
        </section>

        {/* キャンセル・返金 */}
        <section className="mb-8">
          <h2 className="text-xs tracking-[0.15em] text-default-400 font-semibold mb-4">
            キャンセル・返金
          </h2>
          <Card className="bg-content1">
            <CardBody className="gap-4 p-5">
              <div>
                <p className="text-xs text-default-400 font-semibold mb-1.5">
                  キャンセル・解約
                </p>
                <p className="text-sm leading-relaxed">
                  サブスクリプションはいつでもキャンセル可能です。キャンセル後も現在の請求期間の終了まではサービスをご利用いただけます。日割り返金には対応しておりません。
                </p>
              </div>
              <Divider />
              <div>
                <p className="text-xs text-default-400 font-semibold mb-1.5">
                  返金ポリシー
                </p>
                <p className="text-sm leading-relaxed">
                  デジタルサービスの性質上、原則として返金には対応しておりません。ただし、サービスの重大な不具合等があった場合は個別にご相談ください。
                </p>
              </div>
            </CardBody>
          </Card>
        </section>

        {/* フッター */}
        <div className="text-center pt-2">
          <Link href="/" size="sm">
            ログインページに戻る
          </Link>
        </div>
      </div>
    </div>
  );
}

function Row({
  label,
  value,
  children,
}: {
  label: string;
  value?: string;
  children?: React.ReactNode;
}) {
  return (
    <div>
      <p className="text-xs text-default-400 font-semibold mb-1">{label}</p>
      {children ?? <p className="text-sm">{value}</p>}
    </div>
  );
}
