import { useMutation, useQueryClient } from 'react-query'
import { mutateDelete, queryGet } from './base'

const apiBaseUrl = import.meta.env.VITE_API_BASE ?? ''

export type Binder = {
  name: string
}
export type Note = {
  name: number
  binderName: number
  content: string
}

export const useListBinders = () => queryGet<{items: Binder[]}>('/api/binders')

export const useListBinderNotes = (name: string) => queryGet<{items: Note[]}>(`/api/binders/${name}/notes`)

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
