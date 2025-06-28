import { Outlet } from "react-router";
import { Navigation } from "~/components/navigation";

export default function Layout() {
  return (
    <main className="w-dvw h-dvh pb-28 overflow-y-auto">
      <Outlet />
      <Navigation />
    </main>
  );
}
