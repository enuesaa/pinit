import { atom, useAtom } from 'jotai'

type Story = {
  binderName: string
  noteName: string
  noteInput: string
  noteOutput: string
}
const state = atom<Partial<Story>>({
  binderName: '',
  noteName: '',
  noteInput: '',
  noteOutput: '',
})

const useStory = (): [Story, (newone: Partial<Story>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<Story>) => setValue({ ...value, ...newone })

  const data = {
    binderName: value.binderName ?? '',
    noteName: value.noteName ?? '',
    noteInput: value.noteInput ?? '',
    noteOutput: value.noteOutput ?? '',
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
  return (name: string, input: string) => setter({ noteName: name, noteInput: input })
}

export const useSetStoryInput = () => {
  const [_, setter] = useStory()
  return (input: string) => setter({ noteInput: input })
}
