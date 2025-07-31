import { Code, ConnectError, type Interceptor } from "@connectrpc/connect";
import { Cookie } from "~/store/cookies";
import { authClient } from "..";
import { AuthService } from "../pb/auth/v1/auth_pb";

export const authInterceptor: Interceptor = (next) => async (req) => {
  // Skip the interceptor for AuthService requests to avoid circular calls
  if (req.service.typeName === AuthService.typeName) {
    return next(req);
  }

  const token = Cookie.accessToken();

  if (token) {
    req.header.set("Authorization", `Bearer ${token}`);
  }

  try {
    return await next(req);
  } catch (err) {
    if (
      err instanceof ConnectError &&
      err.code === Code.Unauthenticated &&
      err.message === "[unauthenticated] expired token"
    ) {
      console.warn("Token expired, attempting to refresh...");

      try {
        await refreshToken();
        const newToken = Cookie.accessToken();
        if (newToken) {
          req.header.set("Authorization", `Bearer ${newToken}`);
        }
        return await next(req);
      } catch (refreshErr) {
        console.error("Failed to refresh token:", refreshErr);
        throw refreshErr;
      }
    }

    throw err;
  }
};

async function refreshToken() {
  const refreshToken = Cookie.refreshToken();
  if (!refreshToken) {
    throw new Error("No refresh token available");
  }

  const response = await authClient.reauthorize({ refreshToken });
  if (!response.auth) {
    throw new Error("Failed to refresh token");
  }

  Cookie.setAccessToken(response.auth.accessToken);
  Cookie.setRefreshToken(response.auth.refreshToken);
  response.auth.user && Cookie.setUserId(response.auth.user.id);
}
