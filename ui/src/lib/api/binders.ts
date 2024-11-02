import { mutateDelete, mutatePost, mutatePut, queryGet } from './base'

export type Binder = {
  name: string
}

type ListResponse = {
  items: Binder[]
}
export const useListBinders = () => queryGet<ListResponse>('/api/binders')


type GetResponse = Binder & {
  content: string
}
export const useGetBinder = (name: string) => queryGet<GetResponse>(`/api/binders/${name}`)


type CreateResponse = {
  name: string
}
export const useCreateBinder = () => mutatePost<{}, CreateResponse>(`/api/binders`, {
  invalidate: [],
})


type UpdateRequest = {
  name: string
}
type UpdateResponse = {
  name: string
}
export const useUpdateBinder = (binderName: string) => mutatePut<UpdateRequest, UpdateResponse>(`/api/binders/${binderName}`, {
  invalidate: [],
})


export const useDeleteBinder = (name: string) => mutateDelete<void, void>(`/api/binders/${name}`, {
  invalidate: ['/api/binders'],
})
