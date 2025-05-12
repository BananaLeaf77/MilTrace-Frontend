import React, { useState, useEffect } from 'react';
import { Device } from '../types/device';
import { DeviceService } from '../api/deviceService'; // Correct import
import StatusIndicator from './StatusIndicator';
import { useNavigate } from 'react-router-dom';
import { formatDistanceToNow } from 'date-fns'; // Missing import

const DeviceList: React.FC = () => {
  const [devices, setDevices] = useState<Device[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const loadDevices = async () => {
      try {
        const data = await DeviceService.getAllDevices(); // Correct function call
        setDevices(data);
        setLoading(false);
      } catch (err) {
        setError('Failed to load devices');
        setLoading(false);
        console.error(err);
      }
    };
    
    loadDevices();
    const interval = setInterval(loadDevices, 30000); // Refresh every 30 seconds
    
    return () => clearInterval(interval);
  }, []);

  const handleDelete = async (deviceId: string) => {
    if (window.confirm(`Are you sure you want to delete device ${deviceId}?`)) {
      try {
        await DeviceService.deleteDevice(deviceId); // Correct function call
        setDevices(devices.filter(d => d.device_id !== deviceId));
      } catch (err) {
        setError('Failed to delete device');
        console.error(err);
      }
    }
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div className="text-red-500">{error}</div>;

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white rounded-lg overflow-hidden">
        <thead className="bg-gray-100">
          <tr>
            <th className="py-3 px-4 text-left">Device ID</th>
            <th className="py-3 px-4 text-left">Location</th>
            <th className="py-3 px-4 text-left">Last Updated</th>
            <th className="py-3 px-4 text-left">Status</th>
            <th className="py-3 px-4 text-left">Actions</th>
          </tr>
        </thead>
        <tbody className="divide-y divide-gray-200">
          {devices.map(device => (
            <tr key={device.device_id} className="hover:bg-gray-50">
              <td className="py-3 px-4">
                <button 
                  onClick={() => navigate(`/devices/${device.device_id}`)}
                  className="text-blue-500 hover:underline"
                >
                  {device.device_id}
                </button>
              </td>
              <td className="py-3 px-4">
                {device.latitude?.toFixed(6)}, {device.longitude?.toFixed(6)}
              </td>
              <td className="py-3 px-4">
                {formatDistanceToNow(new Date(device.updated_at))} ago
              </td>
              <td className="py-3 px-4">
                <StatusIndicator updatedAt={device.updated_at} />
              </td>
              <td className="py-3 px-4 space-x-2">
                <button 
                  onClick={() => navigate(`/devices/${device.device_id}`)}
                  className="text-blue-500 hover:text-blue-700"
                >
                  View
                </button>
                <button 
                  onClick={() => handleDelete(device.device_id)}
                  className="text-red-500 hover:text-red-700"
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default DeviceList;