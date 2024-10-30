import { useQuery } from 'react-query'

const baseUrl = import.meta.env.VITE_API_BASE ?? ''

type QueryConfig = {}

export const query = <T>(method: string, path: string, config: QueryConfig) =>
  useQuery(path, async (): Promise<T> => {
    const res = await fetch(`${baseUrl}${path}`, {
      method,
      headers: {
        Accept: 'application/json',
      },
    })
    const body = await res.json()
    return body ?? {}
  })
