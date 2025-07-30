import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { authInterceptor } from "./interceptors/auth";
import { AuthService } from "./pb/auth/v1/auth_pb";
import { ReservationService } from "./pb/reservation/v1/reservation_pb";
import { RoomService } from "./pb/room/v1/room_pb";
import { UserService } from "./pb/user/v1/user_pb";

const transport = createConnectTransport({
  baseUrl: "http://localhost:50051",
  interceptors: [authInterceptor],
  // useBinaryFormat: true,
});

export const authClient = createClient(AuthService, transport);
export const reservationClient = createClient(ReservationService, transport);
export const roomClient = createClient(RoomService, transport);
export const userClient = createClient(UserService, transport);
