import { atom, useAtom } from 'jotai'

type Story = {
  binderName: string
}
const state = atom<Partial<Story>>({
  binderName: '',
})

const useStory = (): [Story, (newone: Partial<Story>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<Story>) => setValue({ ...value, ...newone })

  const data = {
    binderName: value.binderName ?? '',
  }
  return [data, setter]
}

export const useGetStory = () => {
  const [data, _] = useStory()
  return data
}
