import { mutatePost } from './base'

type CreateRequest = {
  newName: string
}
export const useRenameBinder = (name: string) => mutatePost<CreateRequest, {}>(`/api/binders/${name}/rename`, {
  invalidate: [],
})
