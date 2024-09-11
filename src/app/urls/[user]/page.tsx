'use client';

import { useParams } from 'next/navigation';

const UserProfile = () => {
  const { user } = useParams(); 
  return (
    <div>
      <h1>Profile of {user}</h1>
    </div>
  );
};

export default UserProfile;
