import { atom, useAtom } from 'jotai'

type Story = {
  binderName: string
  noteName: string
  noteContent: string
}
const state = atom<Partial<Story>>({
  binderName: '',
  noteName: '',
  noteContent: '',
})

const useStory = (): [Story, (newone: Partial<Story>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<Story>) => setValue({ ...value, ...newone })

  const data = {
    binderName: value.binderName ?? '',
    noteName: value.noteName ?? '',
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

export const useSetNote = () => {
  const [_, setter] = useStory()
  return (name: string, content: string) => setter({ noteName: name, noteContent: content })
}

export const useSetStoryInput = () => {
  const [_, setter] = useStory()
  return (content: string) => setter({ noteContent: content })
}
