import { useMutation, useQueryClient } from 'react-query'
import { mutateDelete, mutatePost, queryGet } from './api/base'

const apiBaseUrl = import.meta.env.VITE_API_BASE ?? ''

export type Binder = {
  name: string
}

export type Note = {
  name: number
  binderName: number
  content: string
}

export type Action = {
  id: number
  name: string
  template: string
}

export const useListBinders = () => queryGet<{items: Binder[]}>('/api/binders')

export const useListBinderNotes = (name: string) => queryGet<{items: Note[]}>(`/api/binders/${name}/notes`)

export const useRecog = () =>
  useMutation({
    mutationFn: async (blobUrl: string): Promise<string> => {
      const response = await fetch(blobUrl)
      const blob = await response.blob()
      const file = new File([blob], 'hello', { type: blob.type })
      const res = await fetch(`${apiBaseUrl}/api/recog`, {
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

export const useChat = () => mutatePost<{message: string}, {message: string}>(`/api/chat`, {
  invalidate: [],
})

export const useCreateBinder = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async ({ content }: { content: string }) => {
      const res = await fetch(`${apiBaseUrl}/api/binders`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: '{}',
      })
      const body = await res.json()
      const binderName = body.name
      await fetch(`/api/binders/${binderName}/notes`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content }),
      })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: 'listBinders' })
    },
  })
}

export const useDeleteBinder = (name: string) => mutateDelete<void, void>(`/api/binders/${name}`, {
  invalidate: ['/api/binders'],
})
