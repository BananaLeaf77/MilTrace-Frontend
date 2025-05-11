import axios from 'axios';
import { Device } from '../types/device';

const API_BASE_URL = 'http://your-backend-url/api';

export const fetchAllDevices = async (): Promise<Device[]> => {
  const response = await axios.get(`${API_BASE_URL}/devices`);
  return response.data;
};

export const fetchDevice = async (deviceId: string): Promise<Device> => {
  const response = await axios.get(`${API_BASE_URL}/devices/${deviceId}`);
  return response.data;
};

export const updateDeviceLocation = async (deviceId: string, lat: number, lng: number): Promise<void> => {
  await axios.put(`${API_BASE_URL}/devices/${deviceId}`, {
    latitude: lat,
    longitude: lng
  });
};

export const registerDevice = async (deviceId: string): Promise<void> => {
  await axios.post(`${API_BASE_URL}/devices`, {
    device_id: deviceId
  });
};

export const deleteDevice = async (deviceId: string): Promise<void> => {
  await axios.delete(`${API_BASE_URL}/devices/${deviceId}`);
};