import { Theme } from '@radix-ui/themes'
import '@radix-ui/themes/styles.css'
import { QueryClient, QueryClientProvider } from 'react-query'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import BindersPage from './pages/binders'
import TopPage from './pages/index'

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
