import { useListBinderNotes } from '@/lib/api/binders'
import { BinderStoryInput } from './BinderStoryInput'
import { Box } from '@radix-ui/themes'
import { BinderStoryTrash } from './BinderStoryTrash'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStoryChatButton } from './BinderStoryChatButton'
import { BinderStorySaveButton } from './BinderStorySaveButton'
import { BinderStoryReadonly } from './BinderStoryInputReadonly'

type Props = {
  name: string
  selectedNoteName: string
}
export const BinderStoryNotes = ({ name, selectedNoteName }: Props) => {
  const notes = useListBinderNotes(name)

  return (
    <>
      {notes.data?.items?.map((v,i) => {
        if (v.name !== selectedNoteName) {
          return <BinderStoryReadonly note={v} key={i} />
        }

        return (
          <Box key={i}>
            <BinderStoryRecorder />
            <BinderStoryTrash />
            <BinderStoryChatButton />
            <BinderStorySaveButton binderName={name} name={selectedNoteName} />
            <BinderStoryInput />
          </Box>
        )
      })}
    </>
  )
}
