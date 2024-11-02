import { mutatePost } from './base'

type Request = {
  message: string
}
type Response = {
  message: string
}
export const useChat = () => mutatePost<Request, Response>(`/api/chat`, {
  invalidate: [],
})
