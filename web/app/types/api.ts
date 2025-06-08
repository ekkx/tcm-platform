// Shared types for client-server communication
// These types mirror the gRPC types but don't depend on gRPC libraries

export enum CampusType {
  CAMPUS_UNSPECIFIED = 0,
  NAKAMEGURO = 1,
  IKEBUKURO = 2,
}

export enum PianoType {
  PIANO_TYPE_UNSPECIFIED = 0,
  GRAND = 1,
  UPRIGHT = 2,
  NONE = 3,
}

export interface Room {
  id: string;
  name: string;
  pianoType: PianoType;
  pianoNumber: number;
  isClassroom: boolean;
  isBasement: boolean;
  campusType: CampusType;
  floor: number;
}

export interface Reservation {
  id: number;
  externalId?: string;
  campusType: CampusType;
  date: Date | undefined;
  roomId: string;
  fromHour: number;
  fromMinute: number;
  toHour: number;
  toMinute: number;
  bookerName?: string;
  createdAt: Date | undefined;
}

// Loader data types
export interface HomeLoaderData {
  authenticated: boolean;
  rooms: Room[];
  reservations: Reservation[];
}

// Convert gRPC Room to component Room type
export function convertRoomToComponent(room: Room): any {
  return {
    id: room.id,
    name: room.name,
    piano_type: room.pianoType === PianoType.GRAND ? "grand" : 
                room.pianoType === PianoType.UPRIGHT ? "upright" : "none",
    piano_number: room.pianoNumber,
    is_classroom: room.isClassroom,
    is_basement: room.isBasement,
    campus_code: room.campusType === CampusType.IKEBUKURO ? "2" : "1",
    floor: room.floor
  };
}