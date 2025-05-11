import React from 'react';
import { formatDistanceToNow } from 'date-fns';

interface StatusIndicatorProps {
  updatedAt: string;
}

const StatusIndicator: React.FC<StatusIndicatorProps> = ({ updatedAt }) => {
  const lastUpdated = new Date(updatedAt);
  const minutesAgo = (Date.now() - lastUpdated.getTime()) / (1000 * 60);
  
  const status = minutesAgo < 5 ? 'online' : minutesAgo < 30 ? 'idle' : 'offline';

  return (
    <div className="flex items-center">
      <span className={`inline-block w-3 h-3 rounded-full mr-2 ${
        status === 'online' ? 'bg-green-500' :
        status === 'idle' ? 'bg-yellow-500' : 'bg-gray-500'
      }`}></span>
      <span className="text-sm">
        {status === 'online' ? 'Online' : 
         status === 'idle' ? 'Idle' : 'Offline'}
      </span>
    </div>
  );
};

export default StatusIndicator;