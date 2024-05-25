import { atom, useAtomValue, useSetAtom } from 'jotai'

type Story = {
  input: string
  output: string
}
const storyAtom = atom<Story>({ input: '', output: '' })
export const useGetStory = () => useAtomValue(storyAtom)
export const useSetStoryInput = () => {
  const setStory = useSetAtom(storyAtom)
  const setter = (input: string) => setStory((state) => ({ ...state, input }))
  return setter
}
export const useSetStoryOuptut = () => {
  const setStory = useSetAtom(storyAtom)
  const setter = (output: string) => setStory((state) => ({ ...state, output }))
  return setter
}
