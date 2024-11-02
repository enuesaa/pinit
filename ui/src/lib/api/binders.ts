import { mutateDelete, mutatePost, mutatePut, queryGet } from './base'

export type Binder = {
  name: string
}
export type Note = {
  name: string
  binderName: string
  content: string
}
export type BinderWithContent = Binder & {
  content: string
}

export const useListBinders = () => queryGet<{items: Binder[]}>('/api/binders')

export const useGetBinder = (name: string) => queryGet<BinderWithContent>(`/api/binders/${name}`)

export const useCreateBinder = () => mutatePost<{}, {name: string}>(`/api/binders`, {
  invalidate: [],
})

export const useUpdateBinder = (binderName: string) => mutatePut<{content: string}, {name: string}>(`/api/binders/${binderName}`, {
  invalidate: [],
})

export const useDeleteBinder = (name: string) => mutateDelete<void, void>(`/api/binders/${name}`, {
  invalidate: ['/api/binders'],
})

export const useListBinderNotes = (name: string) => queryGet<{items: Note[]}>(`/api/binders/${name}/notes`)
