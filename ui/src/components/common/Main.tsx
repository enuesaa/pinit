import { Box, Flex } from '@radix-ui/themes'
import { ReactNode } from 'react'
import { SidebarBinders } from './SidebarBinders'
import { SidebarCreateButton } from './SidebarCreateButton'
import { Header } from './Header'

type Props = {
  children: ReactNode
}
export const Main = ({ children }: Props) => {
  return (
    <Flex gap='5'>
      <Box flexGrow='0' flexShrink='0' style={{ width: '500px' }}>
        <Header />
        <SidebarCreateButton />
        <SidebarBinders />
      </Box>
      <Box flexGrow='1' flexShrink='1'>
        {children}
      </Box>
    </Flex>
  )
}
