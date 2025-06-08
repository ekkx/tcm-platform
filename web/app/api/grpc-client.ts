import grpc from '@grpc/grpc-js';
import { AuthorizationServiceClient } from '../proto/v1/authorization/authorization.js';
import { ReservationServiceClient } from '../proto/v1/reservation/reservation.js';
import { RoomServiceClient } from '../proto/v1/room/room.js';

// In server-side code (React Router actions/loaders), we need to use the full URL
// For local development with Docker, use 127.0.0.1:50051 to force IPv4
// For production, set GRPC_SERVER_URL environment variable
const GRPC_SERVER_URL = process.env.GRPC_SERVER_URL || '127.0.0.1:50051';

// Log the gRPC server URL for debugging
if (typeof window === 'undefined') {
  // Only log on server-side
  console.log('[gRPC Client] Server URL:', GRPC_SERVER_URL);
  console.log('[gRPC Client] Environment:', {
    GRPC_SERVER_URL: process.env.GRPC_SERVER_URL,
    NODE_ENV: process.env.NODE_ENV
  });
}

// Removed unused function

export function createAuthorizationClient(): AuthorizationServiceClient {
  console.log('[gRPC Client] Creating authorization client for:', GRPC_SERVER_URL);
  
  const client = new AuthorizationServiceClient(
    GRPC_SERVER_URL,
    grpc.credentials.createInsecure()
  );
  
  console.log('[gRPC Client] Authorization client created');
  return client;
}

export function createReservationClient(): ReservationServiceClient {
  return new ReservationServiceClient(
    GRPC_SERVER_URL,
    grpc.credentials.createInsecure()
  );
}

export function createRoomClient(): RoomServiceClient {
  return new RoomServiceClient(
    GRPC_SERVER_URL,
    grpc.credentials.createInsecure()
  );
}

export function createAuthenticatedClient<T>(
  ClientConstructor: new (address: string, credentials: grpc.ChannelCredentials, options?: object) => T,
  accessToken: string
): T {
  const metadata = new grpc.Metadata();
  metadata.add('authorization', `Bearer ${accessToken}`);
  
  const callCredentials = grpc.credentials.createFromMetadataGenerator(
    (params, callback) => {
      callback(null, metadata);
    }
  );
  
  const channelCredentials = grpc.credentials.createInsecure();
  const combinedCredentials = grpc.credentials.combineChannelCredentials(
    channelCredentials,
    callCredentials
  );
  
  return new ClientConstructor(
    GRPC_SERVER_URL,
    combinedCredentials
  );
}