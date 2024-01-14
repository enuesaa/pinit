import { Theme } from '@radix-ui/themes'
import '@radix-ui/themes/styles.css'
import type { AppProps } from 'next/app'
import { QueryClient, QueryClientProvider } from 'react-query'
import { injectApiMock } from '@/lib/apimock'

export default function App({ Component, pageProps }: AppProps) {
  const queryClient = new QueryClient()
  if (process.env.NODE_ENV === 'development') {
    injectApiMock(queryClient)
  }

  return (
    <QueryClientProvider client={queryClient}>
      <Theme appearance='dark' accentColor='amber'>
        <Component {...pageProps} />
      </Theme>
    </QueryClientProvider>
  )
}
