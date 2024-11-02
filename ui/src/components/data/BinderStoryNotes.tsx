import { useListBinderNotes } from '@/lib/api/binders'
import { BinderStoryInput } from './BinderStoryInput'
import { Box } from '@radix-ui/themes'
import { BinderStoryTrash } from './BinderStoryTrash'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStoryChatButton } from './BinderStoryChatButton'
import { BinderStorySaveButton } from './BinderStorySaveButton'
import { BinderStoryInputReadonly } from './BinderStoryInputReadonly'

type Props = {
  name: string
  selectedNoteName: string
}
export const BinderStoryNotes = ({ name, selectedNoteName }: Props) => {
  const notes = useListBinderNotes(name)

  return (
    <>
      {notes.data?.items?.map(v => {
        if (v.name !== selectedNoteName) {
          return <BinderStoryInputReadonly note={v} key={v.name} />
        }

        return (
          <Box key={v.name} style={{ position: 'relative' }}>
            <div style={{ position: 'absolute', bottom: '10px', right: '10px' }}>
              <BinderStoryRecorder />
              <BinderStoryTrash />
              <BinderStoryChatButton />
              <BinderStorySaveButton binderName={name} name={selectedNoteName} />
            </div>
            <BinderStoryInput />
          </Box>
        )
      })}
    </>
  )
}
