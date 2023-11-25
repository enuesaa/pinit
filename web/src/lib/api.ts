import { useQuery } from 'react-query'

type ConfigResponse = {
  token: string;
}
export const useGetConfig = () => 
  useQuery('getConfig', async (): Promise<ConfigResponse> => {
    const res = await fetch(`/api/config`)
    const body = await res.json()
    return body
  })