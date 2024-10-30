import { TextArea } from '@radix-ui/themes'
import styles from './BinderStoryInput.css'
import { Note } from '@/lib/api/binders'
import { useSetNote } from '@/lib/state/story'
import { MouseEventHandler } from 'react'

type Props = {
  note: Note
}
export const BinderStoryReadonly = ({ note }: Props) => {
  const setNote = useSetNote()

  const handleClick: MouseEventHandler<HTMLTextAreaElement> = (e) => {
    e.preventDefault()
    setNote(note.name, note.content)
  }

  return (
    <TextArea
      onClick={handleClick}
      className={styles.textarea}
      size='3'
      defaultValue={note.content}
      readOnly
    />
  )
}
