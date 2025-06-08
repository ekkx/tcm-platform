import { addToast, Button, Checkbox, Form, Input, Link } from "@heroui/react";
import { useEffect, useState } from "react";
import { createAuthorizationClient } from "~/api/grpc-client";
import type { AuthorizeReply } from "~/proto/v1/authorization/authorization.js";
import { Cookie } from "~/store/cookies";
import type { Route } from "./+types/login";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "練習室予約 ｜ 東京音楽大学" },
    {
      name: "description",
      content: "東京音楽大学の非公式練習室予約サイトです。",
    },
  ];
}

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const loginId = formData.get("login-id")?.toString();
  const password = formData.get("password")?.toString();
  const remember = formData.get("remember")?.toString();

  if (!loginId || !password) {
    return {
      error: "IDとパスワードを入力してください。",
    };
  }

  try {
    console.log("[Login Action] Creating gRPC client...");
    const authClient = createAuthorizationClient();

    console.log("[Login Action] Attempting to authenticate user:", loginId);
    const response = await new Promise<AuthorizeReply>((resolve, reject) => {
      authClient.authorize(
        {
          userId: loginId,
          password: password,
        },
        (error, response) => {
          if (error) {
            console.error("[Login Action] gRPC call failed:", error);
            reject(error);
          } else {
            console.log("[Login Action] gRPC call successful");
            resolve(response);
          }
        }
      );
    });

    if (
      response?.authorization?.accessToken &&
      response?.authorization?.refreshToken
    ) {
      console.log("[Login Action] Authentication successful");
      // Server-side cookie setting would be needed here in a real implementation
      // For now, we'll return the tokens and handle them on the client
      return {
        success: true,
        access_token: response.authorization.accessToken,
        refresh_token: response.authorization.refreshToken,
        remember: remember !== undefined,
      };
    }

    console.error("[Login Action] Invalid response structure:", response);
    return {
      error: "認証に失敗しました。",
    };
  } catch (error: any) {
    console.error("[Login Action] gRPC authentication error:", error);

    // Check if it's a connection error
    if (error.code === 14) {
      return {
        error:
          "サーバーに接続できません。しばらくしてからもう一度お試しください。",
      };
    }

    // Check if it's an authentication error
    if (error.code === 16 || error.code === 7) {
      return {
        error: "IDまたはパスワードが違います。",
      };
    }

    return {
      error: "認証中にエラーが発生しました。",
    };
  }
}

export default function LoginGrpc({ actionData }: Route.ComponentProps) {
  const [isVisible, setIsVisible] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

  const toggleVisibility = () => setIsVisible(!isVisible);

  useEffect(() => {
    if (Cookie.accessToken()) {
      window.location.href = "/";
      return;
    }
  }, []);

  useEffect(() => {
    if (actionData?.success) {
      // Handle successful login
      Cookie.setAccessToken(actionData.access_token);
      Cookie.setRefreshToken(actionData.refresh_token);
      window.location.href = "/";
    } else if (actionData?.error) {
      setIsLoading(false);
      addToast({
        title: "認証失敗",
        description: actionData.error,
        color: "danger",
      });
    }
  }, [actionData]);

  return (
    <div className="grid place-items-center w-dvw h-dvh">
      <div className="flex w-full max-w-sm flex-col gap-4 rounded-large px-8 pb-10 pt-6">
        <div className="grid gap-1">
          <p className="text-foreground-600">東京音楽大学</p>
          <p className="pb-4 text-left text-3xl font-semibold">
            練習室予約サイト
            <span aria-label="emoji" className="ml-2" role="img">
              🎹
            </span>
          </p>
        </div>
        <Form
          className="flex flex-col gap-4"
          validationBehavior="native"
          method="post"
          onSubmit={() => setIsLoading(true)}
        >
          <Input
            isRequired
            label="ログインID"
            labelPlacement="outside"
            name="login-id"
            placeholder="ログインID"
            type="text"
            variant="bordered"
            classNames={{
              label: "after:content-['']",
              input: "text-medium scale-[0.87] origin-left",
            }}
          />
          <Input
            isRequired
            endContent={
              <button type="button" onClick={toggleVisibility}>
                {isVisible ? (
                  <svg
                    className="pointer-events-none text-2xl text-default-400"
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                  >
                    <path
                      fill="currentColor"
                      d="M2.69 6.705a.75.75 0 0 0-1.38.59zm12.897 6.624l-.274-.698zm-6.546.409a.75.75 0 1 0-1.257-.818zm-2.67 1.353a.75.75 0 1 0 1.258.818zM22.69 7.295a.75.75 0 0 0-1.378-.59zM19 11.13l-.513-.547zm.97 2.03a.75.75 0 1 0 1.06-1.06zm-8.72 3.34a.75.75 0 0 0 1.5 0zm5.121-.591a.75.75 0 1 0 1.258-.818zm-10.84-4.25A.75.75 0 0 0 4.47 10.6zm-2.561.44a.75.75 0 0 0 1.06 1.06zM12 13.25c-3.224 0-5.539-1.605-7.075-3.26a13.6 13.6 0 0 1-1.702-2.28a12 12 0 0 1-.507-.946l-.022-.049l-.004-.01l-.001-.001L2 7l-.69.296h.001l.001.003l.003.006l.04.088q.039.088.117.243c.103.206.256.496.462.841c.41.69 1.035 1.61 1.891 2.533C5.54 12.855 8.224 14.75 12 14.75zm3.313-.62c-.97.383-2.071.62-3.313.62v1.5c1.438 0 2.725-.276 3.862-.723zm-7.529.29l-1.413 2.17l1.258.818l1.412-2.171zM22 7l-.69-.296h.001v.002l-.007.013l-.028.062a12 12 0 0 1-.64 1.162a13.3 13.3 0 0 1-2.15 2.639l1.027 1.094a14.8 14.8 0 0 0 3.122-4.26l.039-.085l.01-.024l.004-.007v-.003h.001v-.001zm-3.513 3.582c-.86.806-1.913 1.552-3.174 2.049l.549 1.396c1.473-.58 2.685-1.444 3.651-2.351zm-.017 1.077l1.5 1.5l1.06-1.06l-1.5-1.5zM11.25 14v2.5h1.5V14zm3.709-.262l1.412 2.171l1.258-.818l-1.413-2.171zm-10.49-3.14l-1.5 1.5L4.03 13.16l1.5-1.5z"
                    />
                  </svg>
                ) : (
                  <svg
                    className="pointer-events-none text-2xl text-default-400"
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                  >
                    <path
                      fill="currentColor"
                      d="M9.75 12a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0"
                    />
                    <path
                      fill="currentColor"
                      fillRule="evenodd"
                      d="M2 12c0 1.64.425 2.191 1.275 3.296C4.972 17.5 7.818 20 12 20s7.028-2.5 8.725-4.704C21.575 14.192 22 13.639 22 12c0-1.64-.425-2.191-1.275-3.296C19.028 6.5 16.182 4 12 4S4.972 6.5 3.275 8.704C2.425 9.81 2 10.361 2 12m10-3.75a3.75 3.75 0 1 0 0 7.5a3.75 3.75 0 0 0 0-7.5"
                      clipRule="evenodd"
                    />
                  </svg>
                )}
              </button>
            }
            label="パスワード"
            labelPlacement="outside"
            name="password"
            placeholder="パスワード"
            type={isVisible ? "text" : "password"}
            variant="bordered"
            classNames={{
              label: "after:content-['']",
              input: "text-medium scale-[0.87] origin-left",
            }}
          />
          <div className="flex w-full items-center justify-between px-1 py-2">
            <Checkbox defaultSelected name="remember" size="sm">
              <span className="text-foreground-600">
                このデバイスを記憶する
              </span>
            </Checkbox>
          </div>
          <Button
            className="w-full"
            color="primary"
            type="submit"
            isLoading={isLoading}
          >
            ログイン
          </Button>
        </Form>
        <p className="text-center text-small">
          <Link href="https://www.tokyo-ondai-career.jp/" size="sm" isExternal>
            公式サイトへ
          </Link>
        </p>
      </div>
    </div>
  );
}
