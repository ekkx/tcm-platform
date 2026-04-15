import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { authInterceptor } from "./interceptors/auth";
import { AuthService } from "./pb/auth/v1/auth_pb";
import { ReservationService } from "./pb/reservation/v1/reservation_pb";
import { RoomService } from "./pb/room/v1/room_pb";
import { SubscriptionService } from "./pb/subscription/v1/subscription_pb";
import { UserService } from "./pb/user/v1/user_pb";

const transport = createConnectTransport({
  baseUrl: import.meta.env.VITE_API_BASE_URL ?? "http://localhost:50051",
  interceptors: [authInterceptor],
});

export const authClient = createClient(AuthService, transport);
export const reservationClient = createClient(ReservationService, transport);
export const roomClient = createClient(RoomService, transport);
export const subscriptionClient = createClient(SubscriptionService, transport);
export const userClient = createClient(UserService, transport);
