import { useMutation, useQuery, useQueryClient } from 'react-query'

export type Binder = {
  id: number
  name: string
}

export type Note = {
  id: number
  binderId: number
  content: string
}

export type Action = {
  id: number
  name: string
  template: string
}

export const useListBinders = () =>
  useQuery('listBinders', async (): Promise<Binder[]> => {
    const res = await fetch(`/api/binders`)
    const body = await res.json()
    return body?.items
  })

export const useListBinderNotes = (id: number) =>
  useQuery(`listBinderNotes-${id}`, async (): Promise<Note[]> => {
    const res = await fetch(`/api/binders/${id}/notes`)
    const body = await res.json()
    return body?.items
  })

export const useListActions = () =>
  useQuery('listActions', async (): Promise<Action[]> => {
    const res = await fetch(`/api/actions`)
    const body = await res.json()
    return body?.items
  })

export const useRecog = () =>
  useMutation({
    mutationFn: async (blobUrl: string): Promise<string> => {
      const response = await fetch(blobUrl)
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

export const useChat = () =>
  useMutation({
    mutationKey: 'chat',
    mutationFn: async (message: string): Promise<string> => {
      const res = await fetch('/api/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          message,
        }),
      })
      const body = await res.json()
      return body?.message ?? ''
    },
  })

// todo refactor
export const useCreateBinder = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationKey: 'createBinder',
    mutationFn: async ({ name, content }: { name: string; content: string }) => {
      const res = await fetch('/api/binders', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name,
        }),
      })
      const body = await res.json()
      const binderId = body.id as number
      await fetch('/api/notes', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          binderId,
          content,
        }),
      })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey:'listBinders'})
    },
  })
}

export const useDeleteBinder = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationKey: 'deleteBinder',
    mutationFn: async (id: number) => {
      const res = await fetch(`/api/binders/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      })
      await res.json()
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey:'listBinders'})
    },
  })
}
