import { Code, ConnectError } from "@connectrpc/connect";
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
  type Dispatch,
  type SetStateAction,
} from "react";
import { userClient } from "~/api";
import { type User } from "~/api/pb/user/v1/user_pb";
import { Cookie } from "~/store/cookies";

type AuthContextValue = {
  user: User | null;
  setUser: Dispatch<SetStateAction<User | null>>;
};

const AuthContext = createContext<AuthContextValue | null>(null);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);

  const authenticate = useCallback(async () => {
    try {
      const response = await userClient.getUser(
        {
          userId: Cookie.userId() ?? "",
        },
        {
          headers: {
            Authorization: `Bearer ${Cookie.accessToken() ?? ""}`,
          },
        }
      );

      if (!response.user) {
        window.location.href = "/";
        return;
      }

      setUser(response.user);
    } catch (err) {
      if (
        !(err instanceof ConnectError) ||
        err.code !== Code.Unauthenticated ||
        err.message !== "[unauthenticated] expired token"
      ) {
        // トークンの再発行に関するエラーじゃなさそうなら、無限リロードを避けるためクッキーを削除
        Cookie.destroy();
      }
      window.location.href = "/";
      return;
    }
  }, []);

  useEffect(() => {
    authenticate();
  }, [authenticate]);

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
