import { UUID } from 'crypto';
import { jwtDecode } from 'jwt-decode'
import { useState, useEffect } from 'react'

interface TokenPayload {
  user_id: UUID;
  username: string;
  exp: number;
  iat: number;
}

const decodeToken = (token: string): TokenPayload | null => {
  try {
    const decoded: TokenPayload = jwtDecode<TokenPayload>(token);
    return decoded;
  } catch (error) {
    console.error('Failed to decode token:', error);
    return null;
  }
};

export function useAuth() {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [isLoading, setIsLoading] = useState(true)
  const [user, setUser] = useState<string | null>(null)
  const [id, setId] = useState<UUID | null>(null)

  useEffect(() => {
    async function checkAuth() {
      try {
        const res = await fetch('/api/auth/status')
        const data = await res.json()
        setIsAuthenticated(data.isAuthenticated)
        if (data.isAuthenticated && data.token) {
          const decoded = decodeToken(data.token);
          setUser(decoded ? decoded.username : null); 
          setId(decoded ? decoded.user_id : null);
        }
      } catch (error) {
        console.error('Error checking auth status:', error)
      } finally {
        setIsLoading(false)
      }
    }
    checkAuth()
  }, [])

  return { isAuthenticated, isLoading, user, id }
}