import { Button, Card, CardBody, Divider } from "@heroui/react";

export function ReservationListItem() {
  return (
    <Card>
      <CardBody className="flex flex-row items-center">
        <div className="flex flex-col items-center pl-3 pr-4">
          <span className="text-xs">水</span>
          <span className="text-2xl">28</span>
        </div>
        <Divider orientation="vertical" className="h-11" />
        <div className="mr-auto pl-4">
          <ul className="text-xs text-[10px] opacity-60">
            <li className="flex items-center gap-1.5">
              <svg
                className="w-3 h-3"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <defs>
                  <mask id="solarClockCircleBold0">
                    <g fill="none">
                      <path
                        fill="#fff"
                        d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10"
                      />
                      <path
                        fill="#000"
                        fill-rule="evenodd"
                        d="M12 7.25a.75.75 0 0 1 .75.75v3.69l2.28 2.28a.75.75 0 1 1-1.06 1.06l-2.5-2.5a.75.75 0 0 1-.22-.53V8a.75.75 0 0 1 .75-.75"
                        clip-rule="evenodd"
                      />
                    </g>
                  </mask>
                </defs>
                <path
                  fill="currentColor"
                  d="M0 0h24v24H0z"
                  mask="url(#solarClockCircleBold0)"
                />
              </svg>
              12:30 ~ 16:30
            </li>
            <li className="flex items-center gap-1.5">
              <svg
                className="w-3 h-3"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <path
                  fill="currentColor"
                  fill-rule="evenodd"
                  d="M12 2c-4.418 0-8 4.003-8 8.5c0 4.462 2.553 9.312 6.537 11.174a3.45 3.45 0 0 0 2.926 0C17.447 19.812 20 14.962 20 10.5C20 6.003 16.418 2 12 2m0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4"
                  clip-rule="evenodd"
                />
              </svg>
              中目黒・代官山
            </li>
            <li className="flex items-center gap-1.5">
              <svg
                className="w-3 h-3"
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
              >
                <path
                  fill="currentColor"
                  fill-rule="evenodd"
                  d="M22 8.293c0 3.476-2.83 6.294-6.32 6.294c-.636 0-2.086-.146-2.791-.732l-.882.878c-.519.517-.379.669-.148.919c.096.105.208.226.295.399c0 0 .735 1.024 0 2.049c-.441.585-1.676 1.404-3.086 0l-.294.292s.881 1.025.147 2.05c-.441.585-1.617 1.17-2.646.146l-1.028 1.024c-.706.703-1.568.293-1.91 0l-.883-.878c-.823-.82-.343-1.708 0-2.05l7.642-7.61s-.735-1.17-.735-2.78c0-3.476 2.83-6.294 6.32-6.294S22 4.818 22 8.293m-6.319 2.196a2.2 2.2 0 0 0 2.204-2.195a2.2 2.2 0 0 0-2.204-2.196a2.2 2.2 0 0 0-2.204 2.196a2.2 2.2 0 0 0 2.204 2.195"
                  clip-rule="evenodd"
                />
              </svg>
              P 445（G）
            </li>
          </ul>
        </div>
        <Button size="sm">削除</Button>
      </CardBody>
    </Card>
  );
}
