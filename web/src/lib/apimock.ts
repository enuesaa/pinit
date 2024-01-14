import { QueryClient } from 'react-query'
import { Binder } from './api'

export const injectApiMock = (query: QueryClient) => {
  query.setQueryData<Binder[]>('listBinders', [])
}
