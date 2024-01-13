import { Box, Flex, Section, Text, TextArea } from '@radix-ui/themes'
import { BinderStoryActions } from './BinderStoryActions'

export const BinderStory = () => {
  return (
    <Section p='1'>
      <Flex width='100%' gap='5'>
        <Box grow='1'>
          <BinderStoryActions />
          <TextArea>aa</TextArea>
        </Box>
        <Box grow='1'>
          <Text>aa</Text>          
        </Box>
      </Flex>
    </Section>
  )
}