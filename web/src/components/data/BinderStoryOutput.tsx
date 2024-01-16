import { useGetStory } from '@/lib/state'
import { Text } from '@radix-ui/themes'

export const BinderStoryOutput = () => {
  const story = useGetStory()

  return (
    <Text as='p' m='6'>
      {story.output}
    </Text>
  )
}
