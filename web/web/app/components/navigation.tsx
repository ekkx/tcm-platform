import { Tab, Tabs } from "@heroui/react";
import { CreateReservationButton } from "./create-reservation-button";

export function Navigation() {
  return (
    <div className="fixed bottom-7 inset-x-0">
      <div className="flex items-center justify-center gap-4">
        <Tabs
          color="primary"
          radius="full"
          classNames={{
            tabList:
              "p-2 justify-between gap-6 bg-foreground/10 backdrop-blur-xl border border-[0.5px] border-default-300",
            tab: "w-10 h-10",
          }}
        >
          <Tab
            key="home"
            title={
              <svg
                className="w-5 h-5 text-foreground"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <path
                  fill="currentColor"
                  fill-rule="evenodd"
                  d="M2.52 7.823C2 8.77 2 9.915 2 12.203v1.522c0 3.9 0 5.851 1.172 7.063S6.229 22 10 22h4c3.771 0 5.657 0 6.828-1.212S22 17.626 22 13.725v-1.521c0-2.289 0-3.433-.52-4.381c-.518-.949-1.467-1.537-3.364-2.715l-2-1.241C14.111 2.622 13.108 2 12 2s-2.11.622-4.116 1.867l-2 1.241C3.987 6.286 3.038 6.874 2.519 7.823M11.25 18a.75.75 0 0 0 1.5 0v-3a.75.75 0 0 0-1.5 0z"
                  clip-rule="evenodd"
                />
              </svg>
            }
          />
          <Tab
            key="history"
            title={
              <svg
                className="w-5 h-5 text-foreground"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <g fill="none" stroke="currentColor" stroke-width="1.5">
                  <circle cx="12" cy="12" r="10" />
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M12 8v4l2.5 2.5"
                  />
                </g>
              </svg>
            }
          />
          <Tab
            key="profile"
            title={
              <svg
                className="w-5 h-5 text-foreground"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <g fill="none" stroke="currentColor" stroke-width="1.5">
                  <circle cx="12" cy="6" r="4" />
                  <ellipse cx="12" cy="17" rx="7" ry="4" />
                </g>
              </svg>
            }
          />
        </Tabs>
        <CreateReservationButton />
      </div>
    </div>
  );
}
