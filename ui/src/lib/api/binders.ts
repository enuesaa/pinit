import { useMutation, useQueryClient } from 'react-query'
import { mutateDelete, mutatePost, mutatePut, queryGet } from './base'

const apiBaseUrl = import.meta.env.VITE_API_BASE ?? ''

export type Binder = {
  name: string
}
export type Note = {
  name: string
  binderName: string
  content: string
}

export const useListBinders = () => queryGet<{items: Binder[]}>('/api/binders')

export const useListBinderNotes = (name: string) => queryGet<{items: Note[]}>(`/api/binders/${name}/notes`)

export const useCreateBinder = () => mutatePost<{}, {name: string}>(`/api/binders`, {
  invalidate: [],
})

export const useCreateBinderNote = (binderName: string) => mutatePost<{}, {name: string}>(`/api/binders/${binderName}/notes`, {
  invalidate: [],
})

export const useUpdateBinderNote = (binderName: string, noteName: string) => mutatePut<{content: string}, {}>(`/api/binders/${binderName}/notes/${noteName}`, {
  invalidate: [],
})

export const useDeleteBinder = (name: string) => mutateDelete<void, void>(`/api/binders/${name}`, {
  invalidate: ['/api/binders'],
})
