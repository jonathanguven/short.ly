'use client';

import { useParams } from 'next/navigation';

const UserProfile = () => {
  const params = useParams(); 
  const user = params?.user as string | undefined;

  return (
    <div>
      <h1>Profile of {user}</h1>
    </div>
  );
};

export default UserProfile;
