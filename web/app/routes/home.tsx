import { CreateReservationButton } from "~/components/create-reservation-button";
import { FilterReservationsButton } from "~/components/filter-reservations-button";
import { LogoutButton } from "~/components/logout-button";
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
    <div>
      {/* <div className="sticky top-0 h-14 flex items-center justify-center bg-white/20 backdrop-blur-lg border-b z-50">
        <span className="text-xl">練習室予約</span>
      </div> */}
      <div className="px-4 py-4">
        <div className="flex justify-center items-center">
          <span className="text-xl">予約一覧</span>
        </div>
        <div className="mt-4">
          <div className="grid gap-4">
            <div className="grid">
              <span>本日</span>
              <div className="flex gap-4 py-3 overflow-x-auto">
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
                  isConfirmed
                  campusName="中目黒・代官山"
                  date="2025年04月29日(火)"
                  timeRange="10:00 〜 11:30"
                  userName="田中太郎"
                  roomName="A101"
                  pianoType="グランドピアノ"
                />
              </div>
            </div>
            <div className="grid">
              <span>明日</span>
              <div className="flex gap-4 py-3 overflow-x-auto">
                <ReservationItem
                  isConfirmed
                  campusName="中目黒・代官山"
                  date="2025年04月30日(水)"
                  timeRange="10:00 〜 11:30"
                  userName="田中太郎"
                  roomName="A101"
                  pianoType="グランドピアノ"
                />
              </div>
            </div>
            <div className="grid">
              <span>2025年05月2日(金)</span>
              <div className="flex gap-4 py-3 overflow-x-auto">
                <ReservationItem
                  campusName="池袋"
                  date="2025年05月2日(金)"
                  timeRange="17:00 〜 22:30"
                  roomName="A408"
                  pianoType="アップライトピアノ"
                />
                <ReservationItem
                  campusName="中目黒・代官山"
                  date="2025年05月2日(金)"
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
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="sticky bottom-0 border-t backdrop-blur-xl">
        <div className="flex justify-between items-center h-16 px-14 max-w-[500px] mx-auto">
          <FilterReservationsButton />
          <CreateReservationButton />
          <LogoutButton />
        </div>
      </div>
    </div>
  );
}
