export interface Device {
  device_id: string;
  latitude: number;
  longitude: number;
  updated_at: string;
}

export interface DeviceStatus {
  online: boolean;
  lastUpdated: Date;
}