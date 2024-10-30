import { Section } from '@radix-ui/themes'
import { BinderStoryName } from './BinderStoryName'
import { useGetStory } from '@/lib/state/story'
import { useListBinderNotes } from '@/lib/api/binders'
import { BinderStoryInput } from './BinderStoryInput'
import { Box } from '@radix-ui/themes'
import { BinderStoryTrash } from './BinderStoryTrash'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStoryChatButton } from './BinderStoryChatButton'
import { BinderStorySaveButton } from './BinderStorySaveButton'
import { BinderStoryReadonly } from './BinderStoryInputReadonly'
import { BinderStoryCreateNoteButton } from './BinderStoryCreateNoteButton'

export const BinderStory = () => {
  const story = useGetStory()
  const notes = useListBinderNotes(story.binderName)

  return (
    <Section p='1'>
      <BinderStoryName name={story.binderName} />
      {notes.data?.items?.map((v,i) => {
        if (v.name !== story.noteName) {
          return <BinderStoryReadonly note={v} key={i} />
        }

        return (
          <Box key={i}>
            <BinderStoryRecorder />
            <BinderStoryTrash />
            <BinderStoryInput note={v} />
            <BinderStoryChatButton />
            <BinderStorySaveButton />
          </Box>
        )
      })}
      <BinderStoryCreateNoteButton name={story.binderName} />
    </Section>
  )
}
