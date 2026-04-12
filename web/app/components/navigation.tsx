import { Tab, Tabs } from "@heroui/react";
import { useLocation } from "react-router";
import { CreateReservationButton } from "./create-reservation-button";

export function Navigation() {
  const { pathname } = useLocation();

  return (
    <div className="fixed bottom-7 inset-x-0">
      <div className="flex items-center justify-center gap-4">
        <Tabs
          selectedKey={pathname}
          color="primary"
          radius="full"
          classNames={{
            tabList:
              "p-2 justify-between gap-6 bg-foreground/10 backdrop-blur-xl border border-[0.5px] border-default-300",
            tab: "w-12 h-12",
          }}
        >
          <Tab
            key="/plans"
            href="/plans"
            title={
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <path
                  fill="currentColor"
                  fill-rule="evenodd"
                  d="M12 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14m0-10c-.284 0-.474.34-.854 1.023l-.098.176c-.108.194-.162.29-.246.354c-.085.064-.19.088-.4.135l-.19.044c-.738.167-1.107.25-1.195.532s.164.577.667 1.165l.13.152c.143.167.215.25.247.354s.021.215 0 .438l-.02.203c-.076.785-.114 1.178.115 1.352c.23.174.576.015 1.267-.303l.178-.082c.197-.09.295-.135.399-.135s.202.045.399.135l.178.082c.691.319 1.037.477 1.267.303s.191-.567.115-1.352l-.02-.203c-.021-.223-.032-.334 0-.438s.104-.187.247-.354l.13-.152c.503-.588.755-.882.667-1.165c-.088-.282-.457-.365-1.195-.532l-.19-.044c-.21-.047-.315-.07-.4-.135c-.084-.064-.138-.16-.246-.354l-.098-.176C12.474 6.34 12.284 6 12 6"
                  clip-rule="evenodd"
                />
                <path
                  fill="currentColor"
                  d="m7.093 15.941l-.379 1.382c-.628 2.292-.942 3.438-.523 4.065c.147.22.344.396.573.513c.652.332 1.66-.193 3.675-1.243c.67-.35 1.006-.524 1.362-.562a2 2 0 0 1 .398 0c.356.038.691.213 1.362.562c2.015 1.05 3.023 1.575 3.675 1.243c.229-.117.426-.293.573-.513c.42-.627.105-1.773-.523-4.065l-.379-1.382A8.46 8.46 0 0 1 12 17.5a8.46 8.46 0 0 1-4.907-1.559"
                />
              </svg>
            }
          />
          <Tab
            key="/home"
            href="/home"
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
          {/* <Tab
            key="/history"
            href="/history"
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
          /> */}
          <Tab
            key="/profile"
            href="/profile"
            title={
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <circle cx="12" cy="6" r="4" fill="currentColor" />
                <path
                  fill="currentColor"
                  d="M20 17.5c0 2.485 0 4.5-8 4.5s-8-2.015-8-4.5S7.582 13 12 13s8 2.015 8 4.5"
                />
              </svg>
            }
          />
        </Tabs>
        <CreateReservationButton />
      </div>
    </div>
  );
}
