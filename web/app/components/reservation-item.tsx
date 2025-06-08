import { Card, CardBody, Chip } from "@heroui/react";
import { CancelReservationButton } from "./cancel-reservation-button";
import { EditReservationButton } from "./edit-reservation-button";

type Props = {
  isConfirmed?: boolean;
  campusName: string;
  date: string;
  timeRange: string;
  userName?: string;
  roomName: string;
  pianoType: string;
  reservationId: number;
  onDelete?: () => void;
  rooms?: any[];
};

export function ReservationItem(props: Props) {
  return (
    <Card shadow="none" className="flex-shrink-0 border w-[300px]">
      <CardBody className="grid gap-2">
        <div className="flex">
          <span className="flex items-center gap-1 font-bold">
            <svg
              className="w-5 h-5 text-foreground-400"
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
            >
              <path
                fill="currentColor"
                fillRule="evenodd"
                d="M12 2c-4.418 0-8 4.003-8 8.5c0 4.462 2.553 9.312 6.537 11.174a3.45 3.45 0 0 0 2.926 0C17.447 19.812 20 14.962 20 10.5C20 6.003 16.418 2 12 2m0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4"
                clipRule="evenodd"
              />
            </svg>
            {props.campusName}
          </span>
          {props.isConfirmed ? (
            <Chip
              className="ml-auto"
              size="sm"
              color="success"
              variant="light"
              startContent={
                <svg
                  className="w-4 h-4"
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                >
                  <path
                    fill="currentColor"
                    fillRule="evenodd"
                    d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-5.97-3.03a.75.75 0 0 1 0 1.06l-5 5a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l1.47 1.47l2.235-2.235L14.97 8.97a.75.75 0 0 1 1.06 0"
                    clipRule="evenodd"
                  />
                </svg>
              }
            >
              <span className="font-bold">確定</span>
            </Chip>
          ) : (
            <Chip
              className="ml-auto"
              size="sm"
              variant="light"
              startContent={
                <svg
                  className="w-4 h-4"
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                >
                  <g fill="none" stroke="currentColor" strokeWidth="1.5">
                    <circle cx="12" cy="12" r="10" />
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M12 8v4l2.5 2.5"
                    />
                  </g>
                </svg>
              }
            >
              <span className="font-bold">予約済み</span>
            </Chip>
          )}
        </div>
        <div className="grid gap-2">
          <div className="grid">
            <p className="font-medium">{props.date}</p>
            <p className="flex items-center gap-1.5 text-sm font-medium">
              {props.timeRange}
            </p>
          </div>
          <div className="grid gap-1.5">
            {props.userName && (
              <p className="flex items-center gap-1.5 text-sm font-medium">
                {/* <svg
                  className="w-5 h-5 text-foreground-400"
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                >
                  <circle cx="12" cy="6" r="4" fill="currentColor" />
                  <ellipse cx="12" cy="17" fill="currentColor" rx="7" ry="4" />
                </svg> */}
                {props.userName}
              </p>
            )}
            <p className="flex items-center gap-3 text-sm font-medium">
              <span className="flex items-center gap-1.5">
                <svg
                  className="w-5 h-5 text-foreground-400"
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                >
                  <path
                    fill="currentColor"
                    fillRule="evenodd"
                    d="M2.123 12.816c.287 1.003 1.06 1.775 2.605 3.32l1.83 1.83C9.248 20.657 10.592 22 12.262 22c1.671 0 3.015-1.344 5.704-4.033c2.69-2.69 4.034-4.034 4.034-5.705c0-1.67-1.344-3.015-4.033-5.704l-1.83-1.83c-1.546-1.545-2.318-2.318-3.321-2.605c-1.003-.288-2.068-.042-4.197.45l-1.228.283c-1.792.413-2.688.62-3.302 1.233S3.27 5.6 2.856 7.391l-.284 1.228c-.491 2.13-.737 3.194-.45 4.197m8-5.545a2.017 2.017 0 1 1-2.852 2.852a2.017 2.017 0 0 1 2.852-2.852m8.928 4.78l-6.979 6.98a.75.75 0 0 1-1.06-1.061l6.978-6.98a.75.75 0 0 1 1.061 1.061"
                    clipRule="evenodd"
                  />
                </svg>
                {props.roomName}
              </span>
              <span className="flex items-center gap-1.5">
                <svg
                  className="w-5 h-5 text-foreground-400"
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                >
                  <path
                    fill="currentColor"
                    d="m10.09 11.963l9.274-3.332v5.54a3.8 3.8 0 0 0-1.91-.501c-1.958 0-3.545 1.426-3.545 3.185s1.587 3.185 3.545 3.185c1.959 0 3.546-1.426 3.546-3.185V7.492c0-1.12 0-2.059-.088-2.807a7 7 0 0 0-.043-.31c-.084-.51-.234-.988-.522-1.386a2.2 2.2 0 0 0-.676-.617l-.009-.005c-.771-.461-1.639-.428-2.532-.224c-.864.198-1.936.6-3.25 1.095l-2.284.859c-.615.231-1.137.427-1.547.63c-.435.216-.81.471-1.092.851c-.281.38-.398.79-.452 1.234c-.05.418-.05.926-.05 1.525v7.794a3.8 3.8 0 0 0-1.91-.501C4.587 15.63 3 17.056 3 18.815S4.587 22 6.545 22c1.959 0 3.546-1.426 3.546-3.185z"
                  />
                </svg>
                {props.pianoType}
              </span>
            </p>
          </div>
        </div>
        <div className="flex w-full gap-2 mt-2">
          <EditReservationButton isConfirmed={props.isConfirmed} rooms={props.rooms} />
          <CancelReservationButton
            reservationId={props.reservationId}
            onDelete={props.onDelete}
          />
        </div>
      </CardBody>
    </Card>
  );
}
