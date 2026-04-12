import { useEffect, useRef, useState } from "react";
import { Outlet, useLocation } from "react-router";
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

  const mainRef = useRef<HTMLElement>(null);
  const { pathname } = useLocation();

  useEffect(() => {
    mainRef.current?.scrollTo({ top: 0 });
  }, [pathname]);

  const isProfile = pathname === "/profile" || pathname === "/plans";

  return (
    <AuthProvider>
      {!isProfile && <Header onFilter={setFilters} />}
      <main ref={mainRef} className={`w-dvw h-dvh pb-28 overflow-y-auto ${isProfile ? "" : "pt-[88px]"}`}>
        <Outlet context={{ filters }} />
        <Navigation />
      </main>
    </AuthProvider>
  );
}
