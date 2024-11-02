import { useGetStory } from '@/lib/state/story'
import { Section } from '@radix-ui/themes'
import { HistoryDialog } from './HistoryDialog'
import { StoryChatButton } from './StoryChatButton'
import { StoryInput } from './StoryInput'
import { StoryRecorder } from './StoryRecorder'
import { StoryTrash } from './StoryTrash'
import { StoryName } from './StoryName'

export const Story = () => {
  const story = useGetStory()

  if (story.binderName === '') {
    return <Section p='1' />
  }

  return (
    <Section p='1' className='relative'>
      <StoryName name={story.binderName} />

      <StoryRecorder />
      <StoryTrash />
      <StoryChatButton />
      <HistoryDialog name={story.binderName} />

      <StoryInput name={story.binderName} />
    </Section>
  )
}
