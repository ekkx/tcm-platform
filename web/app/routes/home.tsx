import { Card, CardBody, CardHeader } from "@heroui/react";
import { Tab, Tabs } from "@heroui/tabs";
import { ReservationForm } from "~/components/reservation-form";
import { ReservationItem } from "~/components/reservation-item";
import type { Route } from "./+types/home";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "練習室予約 ｜ 東京音楽大学" },
    {
      name: "description",
      content: "東京音楽大学の非公式練習室予約サイトです。",
    },
  ];
}

export default function Home() {
  return (
    <div className="px-4 py-8">
      <h1 className="text-center mb-8 text-3xl font-bold">練習室予約</h1>
      <Tabs fullWidth size="lg">
        <Tab key="new-reservation" title="新規予約">
          <Card>
            <CardHeader className="grid gap-1">
              <h3 className="text-xl font-bold">新規予約</h3>
              <p className="text-xs text-foreground-400">
                希望の日時と部屋を選択して予約してください。
              </p>
            </CardHeader>
            <CardBody>
              <ReservationForm type="create" />
            </CardBody>
          </Card>
        </Tab>
        <Tab key="my-reservation" title="予約一覧">
          <Card>
            <CardHeader className="grid gap-1">
              <h3 className="text-xl font-bold">予約一覧</h3>
              <p className="text-xs text-foreground-400">
                あなたの予約一覧です。予約日2日前は編集できません。
              </p>
            </CardHeader>
            <CardBody className="grid gap-4">
              <ReservationItem
                isConfirmed
                campusName="中目黒・代官山"
                date="2025年04月29日(火)"
                timeRange="10:00 〜 11:30"
                userName="田中太郎"
                roomName="A101"
                pianoType="グランドピアノ"
              />
              <ReservationItem
                campusName="池袋"
                date="2025年05月2日(金)"
                timeRange="17:00 〜 22:30"
                roomName="A408"
                pianoType="アップライトピアノ"
              />
            </CardBody>
          </Card>
        </Tab>
      </Tabs>
    </div>
  );
}
