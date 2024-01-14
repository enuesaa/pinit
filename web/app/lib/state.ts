import { atom, useAtomValue, useSetAtom } from 'jotai'
import { Action } from './api'

const actionAtom = atom<Action>({id: 0, name: '', template: ''})
export const useGetAction = () => useAtomValue(actionAtom)
export const useChooseAction = () => {
  const setAction = useSetAtom(actionAtom)
  const setter = (action: Action) => setAction(action)
  return setter
}