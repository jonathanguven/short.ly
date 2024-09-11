import { NextApiRequest, NextApiResponse } from 'next'

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  const token = req.cookies.token
  if (token) {
    res.status(200).json({ isAuthenticated: true, token })
  } else {
    res.status(200).json({ isAuthenticated: false, token })
  }
}