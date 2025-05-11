import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { fetchDevice, updateDeviceLocation } from '../api/deviceService';
import { Device } from '../types/device';
import Map from 'react-map-gl';
import 'mapbox-gl/dist/mapbox-gl.css';

const DeviceDetail: React.FC = () => {
  const { deviceId } = useParams<{ deviceId: string }>();
  const [device, setDevice] = useState<Device | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const loadDevice = async () => {
      try {
        if (!deviceId) return;
        const data = await fetchDevice(deviceId);
        setDevice(data);
        setLoading(false);
      } catch (err) {
        setError('Failed to load device');
        setLoading(false);
      }
    };
    
    loadDevice();
    const interval = setInterval(loadDevice, 30000); // Refresh every 30 seconds
    
    return () => clearInterval(interval);
  }, [deviceId]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div className="text-red-500">{error}</div>;
  if (!device) return <div>Device not found</div>;

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="container mx-auto px-4 py-8">
        <div className="bg-white p-6 rounded-lg shadow">
          <h1 className="text-2xl font-bold mb-4">Device: {device.device_id}</h1>
          
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div>
              <h2 className="text-xl font-semibold mb-2">Details</h2>
              <div className="space-y-2">
                <p><span className="font-medium">Device ID:</span> {device.device_id}</p>
                <p><span className="font-medium">Latitude:</span> {device.latitude}</p>
                <p><span className="font-medium">Longitude:</span> {device.longitude}</p>
                <p><span className="font-medium">Last Updated:</span> {new Date(device.updated_at).toLocaleString()}</p>
              </div>
            </div>
            
            <div className="h-64">
              <h2 className="text-xl font-semibold mb-2">Location</h2>
              <Map
                initialViewState={{
                  latitude: device.latitude,
                  longitude: device.longitude,
                  zoom: 14
                }}
                style={{ width: '100%', height: '100%' }}
                mapStyle="mapbox://styles/mapbox/streets-v11"
                mapboxAccessToken="your-mapbox-token"
              >
                <Marker
                  latitude={device.latitude}
                  longitude={device.longitude}
                >
                  <div className="w-6 h-6 bg-blue-500 rounded-full border-2 border-white"></div>
                </Marker>
              </Map>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default DeviceDetail;