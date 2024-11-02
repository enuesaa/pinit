import { Section } from '@radix-ui/themes'
import { useGetStory } from '@/lib/state/story'
import { HistoryDialog } from './HistoryDialog'
import { StoryRecorder } from './StoryRecorder'
import { StoryTrash } from './StoryTrash'
import { StoryChatButton } from './StoryChatButton'
import { StorySaveButton } from './StorySaveButton'
import { StoryInput } from './StoryInput'

export const Story = () => {
  const story = useGetStory()

  if (story.binderName === '') {
    return <Section p='1' />
  }

  return (
    <Section p='1' style={{ position: 'relative' }}>
      <div>{story.binderName}</div>

      <StoryRecorder />
      <StoryTrash />
      <StoryChatButton />
      <StorySaveButton binderName={story.binderName} />
      <HistoryDialog name={story.binderName} />

      <StoryInput />
    </Section>
  )
}
