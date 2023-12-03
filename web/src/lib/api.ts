import { useQuery } from 'react-query'
import { type Binder, type Note } from './schema';

type ConfigResponse = {
  token: string;
}
export const useGetConfig = () => 
  useQuery('getConfig', async (): Promise<ConfigResponse> => {
    const res = await fetch(`/api/config`)
    const body = await res.json()
    return body
  })

export const useListBinders = () => 
  useQuery('listBinders', async(): Promise<Binder[]> => {
    const res = await fetch(`/api/binders`)
    const body = await res.json()
    return body
  })

export const useListBinderNotes = (id: string) => 
  useQuery(`listBinderNotes-${id}`, async(): Promise<Note[]> => {
    const res = await fetch(`/api/binders/${id}/notes`)
    const body = await res.json()
    return body
  })

