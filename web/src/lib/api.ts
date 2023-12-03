import { useQuery } from 'react-query'
import { type Binder, type Note, type Config } from './schema';

export const useGetConfig = () => 
  useQuery('getConfig', async (): Promise<Config> => {
    const res = await fetch(`/api/config`)
    const body = await res.json()
    return body
  })

export const useListBinders = () => 
  useQuery('listBinders', async(): Promise<Binder[]> => {
    const res = await fetch(`/api/binders`)
    const body = await res.json()
    return body?.items
  })

export const useListBinderNotes = (id: string) => 
  useQuery(`listBinderNotes-${id}`, async(): Promise<Note[]> => {
    const res = await fetch(`/api/binders/${id}/notes`)
    const body = await res.json()
    return body
  })

