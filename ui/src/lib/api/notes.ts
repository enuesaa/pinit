import { queryGet } from './base'

export type Note = {
  name: string
  binderName: string
  content: string
}

type ListResponse = {
  items: Note[]
}
export const useListBinderNotes = (name: string) => queryGet<ListResponse>(`/api/binders/${name}/notes`)
