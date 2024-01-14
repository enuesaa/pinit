import { useListActions } from '@/lib/api'
import { Box, Flex, Section } from '@radix-ui/themes'
import { BinderStoryAction } from './BinderStoryAction'
import { BinderStoryInputArea } from './BinderStoryInputArea'
import { BinderStoryOutputArea } from './BinderStoryOutputArea'

export const BinderStory = () => {
  const { data: actions } = useListActions()

  return (
    <Section p='1'>
      <Flex width='100%' gap='5'>
        <Box style={{ flexBasis: '50%' }}>
          {actions && actions.map((a, i) => <BinderStoryAction key={i} action={a} />)}
          <BinderStoryInputArea />
        </Box>
        <Box style={{ flexBasis: '50%' }}>
          <BinderStoryOutputArea />
        </Box>
      </Flex>
    </Section>
  )
}
