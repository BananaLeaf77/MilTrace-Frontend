import React from 'react';
import DeviceMap from '../components/DeviceMap';
import DeviceList from '../components/DeviceList';
import Navbar from '../components/Navbar';

const Dashboard: React.FC = () => {
  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      <div className="container mx-auto px-4 py-8">
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div className="lg:col-span-2">
            <div className="bg-white p-6 rounded-lg shadow mb-8">
              <h2 className="text-xl font-bold mb-4">Device Locations</h2>
              <div className="h-96">
                <DeviceMap />
              </div>
            </div>
          </div>
          <div className="lg:col-span-1">
            <div className="bg-white p-6 rounded-lg shadow">
              <h2 className="text-xl font-bold mb-4">Recent Activity</h2>
              <div className="h-96 overflow-y-auto">
                <DeviceList />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;