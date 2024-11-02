import { useListBinderNotes } from '@/lib/api/binders'
import { BinderStoryInput } from './BinderStoryInput'
import { Box } from '@radix-ui/themes'
import { BinderStoryTrash } from './BinderStoryTrash'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStoryChatButton } from './BinderStoryChatButton'
import { BinderStorySaveButton } from './BinderStorySaveButton'

type Props = {
  name: string
  selectedNoteName: string
}
export const BinderStoryNote = ({ name, selectedNoteName }: Props) => {
  const notes = useListBinderNotes(name)

  if (notes.data?.items.length === 0) {
    return <></>
  }

  return (
    <Box style={{ position: 'relative' }}>
      <div style={{ position: 'absolute', bottom: '10px', right: '10px' }}>
        <BinderStoryRecorder />
        <BinderStoryTrash />
        <BinderStoryChatButton />
        <BinderStorySaveButton binderName={name} name={selectedNoteName} />
      </div>
      <BinderStoryInput />
    </Box>
  )
}
