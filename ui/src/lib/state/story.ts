import { atom, useAtom } from 'jotai'

type Story = {
  binderName: string
  noteContent: string
}
const state = atom<Partial<Story>>({
  binderName: '',
  noteContent: '',
})

const useStory = (): [Story, (newone: Partial<Story>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<Story>) => setValue({ ...value, ...newone })

  const data = {
    binderName: value.binderName ?? '',
    noteContent: value.noteContent ?? '',
  }
  return [data, setter]
}

export const useGetStory = () => {
  const [data, _] = useStory()
  return data
}

export const useSetBinderName = () => {
  const [_, setter] = useStory()
  return (name: string) => setter({ binderName: name })
}

export const useSetStoryInput = () => {
  const [_, setter] = useStory()
  return (content: string) => setter({ noteContent: content })
}
