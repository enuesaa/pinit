import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from 'react-query'
import '@radix-ui/themes/styles.css'
import { Theme } from '@radix-ui/themes'
import TopPage from './pages/index'
import BindersPage from './pages/binders'

export const App = () => {
  const queryClient = new QueryClient()

  // if (process.env.NODE_ENV === 'development') {
  //   injectApiMock(queryClient)
  // }

  return (
    <QueryClientProvider client={queryClient}>
      <Theme appearance='dark' accentColor='purple'>
        <BrowserRouter>
          <Routes>
            <Route path='/' element={<TopPage />} />
            <Route path='/settings' element={<BindersPage />} />
          </Routes>
        </BrowserRouter>
      </Theme>
    </QueryClientProvider>
  )
}