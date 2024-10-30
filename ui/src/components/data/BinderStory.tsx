import { Section } from '@radix-ui/themes'
import { BinderStoryName } from './BinderStoryName'
import { useGetStory } from '@/lib/state/story'
import { BinderStoryCreateNoteButton } from './BinderStoryCreateNoteButton'
import { BinderStoryNotes } from './BinderStoryNotes'

export const BinderStory = () => {
  const story = useGetStory()

  if (story.binderName === '') {
    return <Section p='1' />
  }

  return (
    <Section p='1'>
      <BinderStoryName name={story.binderName} />
      <BinderStoryNotes name={story.binderName} selectedNoteName={story.noteName} />
      <BinderStoryCreateNoteButton name={story.binderName} />
    </Section>
  )
}
