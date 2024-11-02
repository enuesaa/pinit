import { Section } from '@radix-ui/themes'
import { useGetStory } from '@/lib/state/story'
import { HistoryDialog } from './HistoryDialog'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStoryTrash } from './BinderStoryTrash'
import { BinderStoryChatButton } from './BinderStoryChatButton'
import { BinderStorySaveButton } from './BinderStorySaveButton'
import { BinderStoryInput } from './BinderStoryInput'

export const BinderStory = () => {
  const story = useGetStory()

  if (story.binderName === '') {
    return <Section p='1' />
  }

  return (
    <Section p='1' style={{ position: 'relative' }}>
      <div>{story.binderName}</div>

      <BinderStoryRecorder />
      <BinderStoryTrash />
      <BinderStoryChatButton />
      <BinderStorySaveButton binderName={story.binderName} />
      <HistoryDialog name={story.binderName} />

      <BinderStoryInput />
    </Section>
  )
}
