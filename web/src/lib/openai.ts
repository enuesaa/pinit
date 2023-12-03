import { useMutation } from 'react-query'

export const useChat = (apiKey: string) =>
  useMutation({
    mutationFn: async (message: string): Promise<string> => {
      const res = await fetch('/api/chat', {
        method: 'POST',
        body: JSON.stringify({
          message,
        })
      })
      const body = await res.json()
      return body?.massage ?? ''
    },
  })
