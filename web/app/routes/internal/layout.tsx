import { useState } from "react";
import { Outlet } from "react-router";
import { Header, type FilterValues } from "~/components/header";
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
  const [filters, setFilters] = useState<FilterValues>({
    dateRange: null,
    campusType: null,
    pianoType: null,
  });

  return (
    <AuthProvider>
      <Header onFilter={setFilters} />
      <main className="w-dvw h-dvh pt-[88px] pb-28 overflow-y-auto">
        <Outlet context={{ filters }} />
        <Navigation />
      </main>
    </AuthProvider>
  );
}
