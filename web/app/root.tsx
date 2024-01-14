import { Links, LiveReload, Meta, Outlet, Scripts, ScrollRestoration } from '@remix-run/react'
import { Theme } from '@radix-ui/themes'
import '@radix-ui/themes/styles.css'
import { QueryClient, QueryClientProvider } from 'react-query'
import { StrictMode } from 'react'
// import { injectApiMock } from '@/lib/apimock'

export default function App() {
  const queryClient = new QueryClient()
  // if (process.env.NODE_ENV === 'development') {
  //   injectApiMock(queryClient)
  // }

  return (
    <html lang='ja'>
      <head>
        <meta charSet='utf-8' />
        <meta name='viewport' content='width=device-width, initial-scale=1' />
        <Meta />
        <Links />
      </head>
      <body className='dark'>
        <StrictMode>
          <QueryClientProvider client={queryClient}>
            <Theme appearance='dark' accentColor='amber'>
              <Outlet />
            </Theme>
          </QueryClientProvider>
          <ScrollRestoration />
          <Scripts />
          <LiveReload />
        </StrictMode>
      </body>
    </html>
  );
}

export function HydrateFallback() {
  return (
    <html lang='ja'>
      <head>
        <meta charSet='utf-8' />
        <meta name='viewport' content='width=device-width, initial-scale=1' />
        <Meta />
        <Links />
      </head>
      <body className='dark'>
        <Scripts />
        <LiveReload />
      </body>
    </html>
  )
}
