import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // Your Go backend URL

// Type definitions matching your Go struct
export interface Device {
  device_id: string;
  latitude: number;
  longitude: number;
  updated_at: string;
}

export interface ApiResponse<T> {
  message: string;
  data?: T;
  error?: string;
}

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Utility function for error handling
const handleApiError = (error: unknown): never => {
  if (axios.isAxiosError(error)) {
    throw new Error(error.response?.data?.error || error.message);
  }
  throw new Error('An unexpected error occurred');
};

export const DeviceService = {
  ping: async (): Promise<string> => {
    try {
      const response = await api.get('/ping');
      return response.data.message;
    } catch (error) {
      return handleApiError(error);
    }
  },

  registerDevice: async (deviceId: string): Promise<Device> => {
    try {
      const response = await api.post<ApiResponse<Device>>('/device', { 
        device_id: deviceId 
      });
      return response.data.data!;
    } catch (error) {
      return handleApiError(error);
    }
  },

  updateDevice: async (device: Partial<Device>): Promise<Device> => {
    try {
      const response = await api.put<ApiResponse<Device>>('/device/update', device);
      return response.data.data!;
    } catch (error) {
      return handleApiError(error);
    }
  },

  deleteDevice: async (deviceId: string): Promise<void> => {
    try {
      await api.delete(`/device/delete/${deviceId}`);
    } catch (error) {
      handleApiError(error);
    }
  },

  getAllDevices: async (): Promise<Device[]> => {
    try {
      const response = await api.get<ApiResponse<Device[]>>('/device/all');
      return response.data.data || [];
    } catch (error) {
      return handleApiError(error);
    }
  },

  getDevice: async (deviceId: string): Promise<Device> => {
    try {
      const response = await api.get<ApiResponse<Device>>(`/device/get/${deviceId}`);
      return response.data.data!;
    } catch (error) {
      return handleApiError(error);
    }
  },

  updateLocation: async (payload: {
    device_id: string;
    latitude: number;
    longitude: number;
  }): Promise<Device> => {
    try {
      const response = await api.put<ApiResponse<Device>>(
        '/device/receiveLocation',
        payload
      );
      return response.data.data!;
    } catch (error) {
      return handleApiError(error);
    }
  },
};