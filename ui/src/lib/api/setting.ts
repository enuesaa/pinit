import { mutatePut, queryGet } from './base'

export type AppConfig = {
  openaiToken: string
}
export const useGetSetting = () => queryGet<AppConfig>(`/api/setting`)

export const useUpdateSetting = () => mutatePut<AppConfig, {}>(`/api/setting`, {
  invalidate: [],
})
