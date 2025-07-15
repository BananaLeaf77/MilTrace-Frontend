import React, { useState, useEffect } from 'react';
import { MapContainer, TileLayer, Marker, Popup, Polyline } from 'react-leaflet';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';

// Fix for default marker icons
delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon-2x.png',
  iconUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon.png',
  shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png',
});

// Helper to create faded marker icons
const createFadedMarkerIcon = (opacity) => {
  return new L.Icon({
    iconUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-icon.png',
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [1, -34],
    shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png',
    shadowSize: [41, 41],
    className: `faded-marker-${opacity}`,
    opacity: opacity,
  });
};

const DeviceTracker = ({ deviceId = "esp32-001" }) => {
  const [deviceData, setDeviceData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch device data from API
  const fetchDeviceData = async () => {
    try {
      const response = await fetch(`http://localhost:8080/device/get/${deviceId}`);
      const data = await response.json();
      if (data.success) {
        setDeviceData(data.data);
      } else {
        setError(data.message || "Failed to load device data");
      }
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  // Initial fetch and set up refresh interval
  useEffect(() => {
    fetchDeviceData();
    const interval = setInterval(fetchDeviceData, 1000);
    return () => clearInterval(interval);
  }, [deviceId]);

  if (loading) return <div className="loading">Loading device data...</div>;
  if (error) return <div className="error">Error: {error}</div>;
  if (!deviceData) return <div className="error">No device data found</div>;

  const currentPosition = [deviceData.latitude, deviceData.longitude];
  const pathHistory = deviceData.locations.map(loc => [loc.latitude, loc.longitude]);

  // Calculate time range for opacity calculation
  const newestDate = new Date(deviceData.updated_at);
  const oldestDate = new Date(deviceData.locations[deviceData.locations.length - 1]?.created_at || deviceData.updated_at);
  const totalTimeRange = newestDate - oldestDate || 1; // Prevent division by zero

  return (
    <div style={{ fontFamily: 'Arial, sans-serif', maxWidth: '1200px', margin: '0 auto', padding: '20px' }}>
      <h1 style={{ textAlign: 'center', color: '#2c3e50' }}>GPS Device Tracker: {deviceId}</h1>
      
      <div style={{ 
        display: 'grid', 
        gridTemplateColumns: '1fr 300px', 
        gap: '20px',
        height: '80vh'
      }}>
        {/* Map Section */}
        <div style={{ borderRadius: '8px', overflow: 'hidden', boxShadow: '0 2px 10px rgba(0,0,0,0.1)' }}>
          <MapContainer 
            center={currentPosition} 
            zoom={13} 
            style={{ height: '100%', width: '100%' }}
          >
            <TileLayer
              url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
            />

            {/* Current position (full opacity) */}
            <Marker position={currentPosition}>
              <Popup>
                <div>
                  <h3>Current Position</h3>
                  <p>Device: {deviceData.device_id}</p>
                  <p>Lat: {deviceData.latitude.toFixed(6)}</p>
                  <p>Lng: {deviceData.longitude.toFixed(6)}</p>
                  <p>Last update: {new Date(deviceData.updated_at).toLocaleString()}</p>
                </div>
              </Popup>
            </Marker>

            {/* Historical markers with fading effect */}
            {deviceData.locations.map((location, index) => {
              const locationDate = new Date(location.created_at);
              const ageRatio = (newestDate - locationDate) / totalTimeRange;
              const opacity = Math.max(0.2, 1 - ageRatio * 0.8); // Range: 0.2-1.0

              return (
                <Marker
                  key={location.id}
                  position={[location.latitude, location.longitude]}
                  icon={createFadedMarkerIcon(opacity)}
                  opacity={opacity}
                >
                  <Popup>
                    <div>
                      <h3>Historical Position</h3>
                      <p>Recorded: {locationDate.toLocaleString()}</p>
                      <p>Age: {Math.round(ageRatio * 100)}% old</p>
                    </div>
                  </Popup>
                </Marker>
              );
            })}

            <Polyline 
              positions={pathHistory} 
              color="blue" 
              weight={3} 
              opacity={0.5}
            />
          </MapContainer>
        </div>
        
        {/* Status Panel */}
        <div style={{ 
          backgroundColor: '#f8f9fa', 
          padding: '20px', 
          borderRadius: '8px',
          boxShadow: '0 2px 10px rgba(0,0,0,0.1)',
          overflowY: 'auto'
        }}>
          <h2 style={{ marginTop: 0, color: '#2c3e50' }}>Device Status</h2>
          
          <div style={{ marginBottom: '20px' }}>
            <h3>Current Position</h3>
            <p><strong>Latitude:</strong> {deviceData.latitude.toFixed(6)}</p>
            <p><strong>Longitude:</strong> {deviceData.longitude.toFixed(6)}</p>
            <p><strong>Last Updated:</strong> {new Date(deviceData.updated_at).toLocaleString()}</p>
          </div>
          
          <div>
            <h3>Location History</h3>
            <div style={{ maxHeight: '300px', overflowY: 'auto' }}>
              {deviceData.locations.map((location, index) => (
                <div key={location.id} style={{ 
                  padding: '10px', 
                  marginBottom: '10px', 
                  backgroundColor: index === 0 ? '#e6f7ff' : '#fff',
                  border: '1px solid #eee',
                  borderRadius: '4px'
                }}>
                  <p><strong>#{deviceData.locations.length - index}</strong></p>
                  <p>Lat: {location.latitude.toFixed(6)}</p>
                  <p>Lng: {location.longitude.toFixed(6)}</p>
                  <p>Time: {new Date(location.created_at).toLocaleTimeString()}</p>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default DeviceTracker;