'use client'

import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { UUID } from 'crypto';

interface AuthContextType {
  isAuthenticated: boolean;
  user: string | null;
  id: UUID | null;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<string | null>(null);
  const [id, setId] = useState<UUID | null>(null);

  useEffect(() => {
    async function checkAuth() {
      try {
        const res = await fetch('/api/auth/status', {
          method: 'GET',
          credentials: 'same-origin',
        });
        const data = await res.json();
        setIsAuthenticated(data.isAuthenticated);
        setUser(data.user);
        setId(data.user_id);
      } catch (error) {
        console.error('Error checking auth status:', error);
      }
    }
    checkAuth();
  }, []);

  const authContextValue = {
    isAuthenticated,
    user,
    id,
  };

  return (
    <AuthContext.Provider value={authContextValue}>
      {children}
    </AuthContext.Provider>
  );
}
