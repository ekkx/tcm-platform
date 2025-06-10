import { createAuthorizationClient } from './grpc-client';
import { Cookie } from '~/store/cookies';

export async function refreshAccessToken(refreshToken: string): Promise<{ accessToken: string; refreshToken: string } | null> {
  try {
    const authClient = createAuthorizationClient();
    
    const result = await new Promise<any>((resolve, reject) => {
      authClient.reauthorize(
        { refreshToken },
        (error, response) => {
          if (error) {
            console.error('[Refresh Token] Failed:', error);
            reject(error);
          } else {
            resolve(response);
          }
        }
      );
    });

    if (result?.authorization?.accessToken && result?.authorization?.refreshToken) {
      return {
        accessToken: result.authorization.accessToken,
        refreshToken: result.authorization.refreshToken
      };
    }
    
    return null;
  } catch (error) {
    console.error('[Refresh Token] Error:', error);
    return null;
  }
}

export async function getAuthTokensFromRequest(request: Request): Promise<{ accessToken: string; refreshToken: string } | null> {
  const cookieHeader = request.headers.get('Cookie');
  const cookies = new Map<string, string>();
  
  if (cookieHeader) {
    cookieHeader.split(';').forEach(cookie => {
      const [key, value] = cookie.split('=').map(s => s.trim());
      if (key && value) {
        cookies.set(key, value);
      }
    });
  }

  const accessToken = cookies.get('access-token');
  const refreshToken = cookies.get('refresh-token');

  if (!accessToken || !refreshToken) {
    return null;
  }

  return { accessToken, refreshToken };
}

export async function tryAuthenticatedRequest<T>(
  request: Request,
  makeRequest: (accessToken: string) => Promise<T>
): Promise<{ data?: T; redirect?: string; newTokens?: { accessToken: string; refreshToken: string } }> {
  const tokens = await getAuthTokensFromRequest(request);
  
  if (!tokens) {
    return { redirect: '/login' };
  }

  try {
    // Try with current access token
    const data = await makeRequest(tokens.accessToken);
    return { data };
  } catch (error: any) {
    // Check if it's an authentication error (code 16 = UNAUTHENTICATED)
    if (error.code === 16) {
      console.log('[Auth] Access token expired, attempting refresh...');
      
      // Try to refresh the token
      const newTokens = await refreshAccessToken(tokens.refreshToken);
      
      if (newTokens) {
        // Set cookies will be handled in the loader/action response
        try {
          const data = await makeRequest(newTokens.accessToken);
          return { 
            data,
            // Pass new tokens back to be set in cookies
            newTokens
          };
        } catch (retryError) {
          console.error('[Auth] Request failed after token refresh:', retryError);
          return { redirect: '/login' };
        }
      } else {
        console.error('[Auth] Token refresh failed');
        return { redirect: '/login' };
      }
    }
    
    // For other errors, throw them to be handled by the caller
    throw error;
  }
}