import { useMutation, useQueryClient } from 'react-query'

const baseUrl = import.meta.env.VITE_API_BASE ?? ''

export type MutateConfig = {
  invalidate: string[]
}
export const mutate = <R, T>(method: string, path: string, { invalidate }: MutateConfig) => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (data: R): Promise<T> => {
      const res = await fetch(`${baseUrl}${path}`, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      })
      const body = await res.json()
      return body ?? {}
    },
    onSuccess: () => queryClient.invalidateQueries(invalidate),
  })
}
