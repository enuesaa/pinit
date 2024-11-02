import { TextArea } from '@radix-ui/themes'
import { Note } from '@/lib/api/binders'

type Props = {
  note: Note
}
export const StoryInputReadonly = ({ note }: Props) => {
  return (
    <TextArea
      className='min-h-[80vh] p-[10px] outline-none'
      size='3'
      defaultValue={note.content}
      readOnly
      style={{ minHeight: '10vh', height: 'fit-content', margin: '10px' }}
    />
  )
}
