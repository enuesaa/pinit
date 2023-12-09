import { useMutation, useQuery } from 'react-query'
import { type Binder, type Note, type Config } from './schema';

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


export const useRecog = () =>
  useMutation({
    mutationFn: async (blobUrl: string): Promise<string> => {
      const response = await fetch(blobUrl);
      const blob = await response.blob()
      const file = new File([blob], 'hello', { type: blob.type })
      const res = await fetch('/api/recog', {
        method: 'POST',
        headers: {
          'Content-Type': 'audio/wav',
        },
        body: file,
      })
      const resbody = await res.json()
      return resbody?.text ?? ''
    },
  })
