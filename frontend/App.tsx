import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import DeviceDetail from './pages/DeviceDetail';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/devices/:deviceId" element={<DeviceDetail />} />
      </Routes>
    </Router>
  );
};

export default App;