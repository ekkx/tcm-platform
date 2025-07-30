import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from "react";
import { userClient } from "~/api";
import { type User } from "~/api/pb/user/v1/user_pb";
import { Cookie } from "~/store/cookies";

const AuthContext = createContext<User | null>(null);

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
    } catch (error) {
      if (error instanceof Error) {
        window.location.href = "/";
        return;
      }
    }
  }, []);

  useEffect(() => {
    authenticate();
  }, [authenticate]);

  return <AuthContext.Provider value={user}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  return useContext(AuthContext);
}
