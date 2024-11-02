import { atom, useAtom } from 'jotai'

type Story = {
  binderName: string
  content: string
}
const state = atom<Partial<Story>>({
  binderName: '',
  content: '',
})

const useStory = (): [Story, (newone: Partial<Story>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<Story>) => setValue({ ...value, ...newone })

  const data = {
    binderName: value.binderName ?? '',
    content: value.content ?? '',
  }
  return [data, setter]
}

export const useGetStory = () => {
  const [data, _] = useStory()
  return data
}

export const useSetStory = () => {
  const [_, setter] = useStory()
  return (binderName: string, content: string) => setter({ binderName, content })
}

export const useSetStoryInput = () => {
  const [_, setter] = useStory()
  return (content: string) => setter({ content })
}
