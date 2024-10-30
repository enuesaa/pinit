import { Box, Flex } from '@radix-ui/themes'
import { ReactNode } from 'react'
import { ListBinders } from '../data/ListBinders'
import { Header } from './Header'

type Props = {
  children: ReactNode
}
export const WithSidebar = ({ children }: Props) => {
  return (
    <Flex gap='5'>
      <Box flexGrow='0' flexShrink='0' style={{ width: '500px' }}>
        <Header />
        create
        <ListBinders />
      </Box>
      <Box flexGrow='1' flexShrink='1'>
        {children}
      </Box>
    </Flex>
  )
}
