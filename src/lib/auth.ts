import jwt from 'jsonwebtoken';
import { UUID } from 'crypto';

interface TokenPayload {
  user_id: UUID;
  username: string;
  exp: number;
  iat: number;
}

interface AuthResult {
  isAuthenticated: boolean;
  user: string | null;
  user_id: UUID | null;
}

export function validateToken(token: string | undefined): AuthResult {
  const jwtSecret = process.env.JWT_SECRET;

  if (!jwtSecret) {
    throw new Error('JWT secret is not defined');
  }

  if (!token) {
    return {
      isAuthenticated: false,
      user: null,
      user_id: null,
    };
  }

  try {
    const decoded = jwt.verify(token, jwtSecret) as TokenPayload;
    return {
      isAuthenticated: true,
      user: decoded.username,
      user_id: decoded.user_id,
    };
  } catch (error) {
    console.error('Failed to decode token:', error);
    return {
      isAuthenticated: false,
      user: null,
      user_id: null,
    };
  }
}