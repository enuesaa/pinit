import { useGetStory } from '@/lib/state'
import { Box, Link } from '@radix-ui/themes'
import Markdown from 'react-markdown'

export const BinderStoryOutput = () => {
  const story = useGetStory()

  return (
    <Box mt='6' ml='3'>
      <Markdown components={{ a: LinkBlank }}>{story.output}</Markdown>
    </Box>
  )
}

const LinkBlank = ({ href, children }: any) => {
  return (
    <Link href={href} target='_blank'>
      {children}
    </Link>
  )
}
