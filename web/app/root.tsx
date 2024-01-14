import { Links, LiveReload, Meta, Outlet, Scripts, ScrollRestoration } from '@remix-run/react'
import { Theme } from '@radix-ui/themes'
import '@radix-ui/themes/styles.css'
import { QueryClient, QueryClientProvider } from 'react-query'
import { injectApiMock } from '@/lib/apimock'

export default function App() {
  const queryClient = new QueryClient()
  if (process.env.NODE_ENV === 'development') {
    injectApiMock(queryClient)
  }
  return (
    <html lang='ja'>
      <head>
        <meta charSet='utf-8' />
        <meta name='viewport' content='width=device-width, initial-scale=1' />
        <Meta />
        <Links />
      </head>
      <body>
        <QueryClientProvider client={queryClient}>
          <Theme appearance='dark' accentColor='amber'>
            <Outlet />
            <ScrollRestoration />
            <Scripts />
            <LiveReload />
          </Theme>
        </QueryClientProvider>
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
      <body>
        <Scripts />
        <LiveReload />
      </body>
    </html>
  )
}
