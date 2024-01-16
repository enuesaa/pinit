import { atom, useAtomValue, useSetAtom } from 'jotai'
import { Action } from './api'

const actionAtom = atom<Action>({ id: 0, name: '', template: '' })
export const useGetAction = () => useAtomValue(actionAtom)
export const useChooseAction = () => {
  const setAction = useSetAtom(actionAtom)
  const setter = (action: Action) => setAction(action)
  return setter
}

type Story = {
  input: string
  output: string
}
const storyAtom = atom<Story>({input: '', output: ''})
export const useGetStory = () => useAtomValue(storyAtom)
export const useSetStoryInput = () => {
  const setStory = useSetAtom(storyAtom)
  const setter = (input: string) => setStory(state => {
    state.input = input
    return state
  })
  return setter
}
export const useSetStoryOuptut = () => {
  const setStory = useSetAtom(storyAtom)
  const setter = (output: string) => setStory(state => {
    state.output = output
    return state
  })
  return setter
}
