import { useMutation } from 'react-query'

export const useChat = () =>
  useMutation({
    mutationFn: async (message: string): Promise<string> => {
      const res = await fetch('/api/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          message,
        })
      })
      const body = await res.json()
      return body?.message ?? ''
    },
  })
