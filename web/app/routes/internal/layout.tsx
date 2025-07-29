import { Outlet } from "react-router";
import { Navigation } from "~/components/navigation";
import { AuthProvider } from "~/providers/auth-provider";

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
