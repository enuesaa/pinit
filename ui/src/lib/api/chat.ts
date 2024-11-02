import { mutatePost } from './base'

export const useChat = () => mutatePost<{ message: string }, { message: string }>(`/api/chat`, {
  invalidate: [],
})
