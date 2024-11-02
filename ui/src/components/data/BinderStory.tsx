import { Section } from '@radix-ui/themes'
import { BinderStoryName } from './BinderStoryName'
import { useGetStory } from '@/lib/state/story'
import { BinderStoryNote } from './BinderStoryNote'
import { HistoryDialog } from './HistoryDialog'

export const BinderStory = () => {
  const story = useGetStory()

  if (story.binderName === '') {
    return <Section p='1' />
  }

  return (
    <Section p='1'>
      <HistoryDialog name={story.binderName} />
      <BinderStoryName name={story.binderName} />
      <BinderStoryNote name={story.binderName} />
    </Section>
  )
}
