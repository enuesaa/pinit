import { QueryClient } from 'react-query'
import { Binder, Config } from './schema'

// see https://oita.oika.me/2021/09/19/react-query-stub
export const injectApiMock = (query: QueryClient) => {
  query.setQueryData<Config>('getConfig', {
    token: 'aa',
  })
  query.setQueryData<Binder[]>('listBinders', [])
}
