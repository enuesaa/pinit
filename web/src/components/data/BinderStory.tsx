import { Box, Flex, Section, Text, TextArea } from '@radix-ui/themes'
import { BinderStoryActions } from './BinderStoryActions'
import { FaMicrophone } from 'react-icons/fa'
import styles from './BinderStory.css'

export const BinderStory = () => {
  return (
    <Section p='1'>
      <Flex width='100%' gap='5'>
        <Box grow='1' position='relative'>
          <BinderStoryActions />
          <TextArea className={styles.textarea} />
          <span className={styles.speakButton}><FaMicrophone /></span>
        </Box>
        <Box grow='1'>
          <Text>aa</Text>          
        </Box>
      </Flex>
    </Section>
  )
}