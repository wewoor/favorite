// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

type Data = {
  name: string
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {

  const fetchRes = await fetch(`${process.env.REMOTE_HOST}/containers`)
  console.log('res->', fetchRes)
  res.status(200).json(await fetchRes.json())
}
