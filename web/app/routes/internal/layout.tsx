import { Outlet } from "react-router";
import { Navigation } from "~/components/navigation";
import { AuthProvider } from "~/providers/auth-provider";
import type { Route } from "./+types/layout";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "練習室予約 ｜ 東京音楽大学" },
    {
      name: "description",
      content: "東京音楽大学の非公式練習室予約サイトです。",
    },
  ];
}

export default function Layout() {
  return (
    <AuthProvider>
      <main className="w-dvw h-dvh pb-28 overflow-y-auto">
        <Outlet />
        <Navigation />
      </main>
    </AuthProvider>
  );
}
