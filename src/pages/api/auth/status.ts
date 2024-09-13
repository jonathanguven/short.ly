import { NextApiRequest, NextApiResponse } from 'next';
import jwt from 'jsonwebtoken';
import { UUID } from 'crypto';

interface TokenPayload {
  user_id: UUID; 
  username: string;
  exp: number;
  iat: number;
}

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  const token = req.cookies.token;
  const jwtSecret = process.env.JWT_SECRET;

  if (!jwtSecret) {
    throw new Error('JWT secret is not defined');
  }

  if (!token) {
    return res.status(200).json({
      isAuthenticated: false,
      user: null,
      user_id: null,
    });
  }

  try {
    const decoded = jwt.verify(token, jwtSecret) as TokenPayload;
    const data = {
      isAuthenticated: true,
      user: decoded.username,
      user_id: decoded.user_id,
    }
    return res.status(200).json(data);
  } catch (error) {
    console.error('Failed to decode token:', error);

    return res.status(200).json({
      isAuthenticated: false,
      user: null,
      user_id: null,
    });
  }
}
