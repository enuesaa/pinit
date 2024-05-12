import { useMutation, useQuery, useQueryClient } from 'react-query'
import { Chat, CreateBinder, CreateNote, DeleteBinder, ListActios, ListBinders, ListNotes, Recog } from '../../wailsjs/go/main/App'
import { usecase } from '../../wailsjs/go/models'

export const useListBinders = () =>
  useQuery('listBinders', async (): Promise<usecase.ListBindersItem[]> => {
    return await ListBinders()
  })

export const useListBinderNotes = (id: number) =>
  useQuery(`listBinderNotes-${id}`, async (): Promise<usecase.ListNotesItem[]> => {
    return await ListNotes(id)
  })

export const useListActions = () =>
  useQuery('listActions', async (): Promise<usecase.ListActionsItem[]> => {
    return await ListActios()
  })

export const useRecog = () =>
  useMutation({
    mutationFn: async (blobUrl: string): Promise<string> => {
      const response = await fetch(blobUrl)
      const blob = await response.blob()
      const text = await blob.text()
      console.log(text)
      const res = await Recog(text)
      return res.text ?? ''
    },
  })

export const useChat = () =>
  useMutation({
    mutationKey: 'chat',
    mutationFn: async (message: string): Promise<string> => {
      const res = await Chat({message})
      return res.message
    },
  })

// todo refactor
export const useCreateBinder = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationKey: 'createBinder',
    mutationFn: async ({ name, content }: { name: string; content: string }) => {
      const res = await CreateBinder({name})
      const binderId = res.id as number
      await CreateNote({
        binderId,
        content,
      })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey:'listBinders'})
    },
  })
}

export const useDeleteBinder = () => {
  const queryClient = useQueryClient()

  return useMutation({
    mutationKey: 'deleteBinder',
    mutationFn: async (id: number) => {
      await DeleteBinder(id)
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey:'listBinders'})
    },
  })
}
