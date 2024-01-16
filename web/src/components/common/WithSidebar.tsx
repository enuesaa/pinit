import { Box, Flex } from '@radix-ui/themes'
import { Header } from './Header'
import { ReactNode } from 'react'
import { ListBinders } from '../data/ListBinders'

type Props = {
  children: ReactNode
}
export const WithSidebar = ({ children }: Props) => {
  return (
    <Flex gap='5'>
      <Box grow='0' shrink='0' style={{width: '500px'}}>
        <Header />
        <ListBinders />
      </Box>
      <Box grow='1' shrink='1'>
        {children}
      </Box>
    </Flex>
  )
}