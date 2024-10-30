import { TextArea } from '@radix-ui/themes'
import styles from './BinderStoryInput.css'
import { Note } from '@/lib/api/binders'

type Props = {
  note: Note
}
export const BinderStoryReadonly = ({ note }: Props) => {
  return (
    <TextArea
      className={styles.textarea}
      size='3'
      defaultValue={note.content}
      readOnly
    />
  )
}
