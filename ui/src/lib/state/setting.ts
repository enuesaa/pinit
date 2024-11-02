import { atom, useAtom } from 'jotai'

type AppSetting = {
  openaiToken: string
}
const state = atom<Partial<AppSetting>>({
  openaiToken: '',
})

export const useAppSetting = (): [AppSetting, (newone: Partial<AppSetting>) => void] => {
  const [value, setValue] = useAtom(state)
  const setter = (newone: Partial<AppSetting>) => setValue({ ...value, ...newone })

  const data = {
    openaiToken: value.openaiToken ?? '',
  }
  return [data, setter]
}
