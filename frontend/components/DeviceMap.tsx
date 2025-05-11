import React, { useState, useEffect } from 'react';
import Map, { Marker, Popup } from 'react-map-gl';
import 'mapbox-gl/dist/mapbox-gl.css';
import { Device } from '../types/device';
import { fetchAllDevices } from '../api/deviceService';
import StatusIndicator from './StatusIndicator';

const MAPBOX_TOKEN = 'your-mapbox-token';

const DeviceMap: React.FC = () => {
  const [devices, setDevices] = useState<Device[]>([]);
  const [selectedDevice, setSelectedDevice] = useState<Device | null>(null);
  const [viewport, setViewport] = useState({
    latitude: -6.1754,
    longitude: 106.8272,
    zoom: 12
  });

  useEffect(() => {
    const loadDevices = async () => {
      const data = await fetchAllDevices();
      setDevices(data);
    };
    
    loadDevices();
    const interval = setInterval(loadDevices, 30000); // Refresh every 30 seconds
    
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="h-full w-full rounded-lg overflow-hidden shadow-lg">
      <Map
        {...viewport}
        mapStyle="mapbox://styles/mapbox/streets-v11"
        mapboxAccessToken={MAPBOX_TOKEN}
        onMove={evt => setViewport(evt.viewState)}
      >
        {devices.map(device => (
          <Marker
            key={device.device_id}
            latitude={device.latitude}
            longitude={device.longitude}
            onClick={e => {
              e.originalEvent.stopPropagation();
              setSelectedDevice(device);
            }}
          >
            <div className="cursor-pointer">
              <StatusIndicator updatedAt={device.updated_at} />
            </div>
          </Marker>
        ))}

        {selectedDevice && (
          <Popup
            latitude={selectedDevice.latitude}
            longitude={selectedDevice.longitude}
            onClose={() => setSelectedDevice(null)}
            closeButton={false}
            anchor="bottom"
          >
            <div className="p-2">
              <h3 className="font-bold">{selectedDevice.device_id}</h3>
              <p>Lat: {selectedDevice.latitude.toFixed(6)}</p>
              <p>Lng: {selectedDevice.longitude.toFixed(6)}</p>
              <p>Last updated: {new Date(selectedDevice.updated_at).toLocaleString()}</p>
              <a 
                href={`/devices/${selectedDevice.device_id}`}
                className="text-blue-500 hover:underline"
              >
                View details
              </a>
            </div>
          </Popup>
        )}
      </Map>
    </div>
  );
};

export default DeviceMap;