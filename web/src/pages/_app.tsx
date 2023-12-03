import { Theme } from '@radix-ui/themes'
import '@radix-ui/themes/styles.css'
import type { AppProps } from 'next/app'
import { QueryClient, QueryClientProvider } from 'react-query'

export default function App({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient()
  // see https://oita.oika.me/2021/09/19/react-query-stub
  // mock
  queryClient.setQueryData('getConfig', {
    token: 'aa',
  })

  return (
    <QueryClientProvider client={queryClient}>
      <Theme appearance='dark' accentColor='amber'>
        <Component {...pageProps} />
      </Theme>
    </QueryClientProvider>
  )
}
