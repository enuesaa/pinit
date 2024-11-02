import { TextArea } from '@radix-ui/themes'
import styles from './BinderStoryInput.css'
import { Note } from '@/lib/api/binders'

type Props = {
  note: Note
}
export const BinderStoryInputReadonly = ({ note }: Props) => {
  return (
    <TextArea
      className={styles.textarea}
      size='3'
      defaultValue={note.content}
      readOnly
      style={{ minHeight: '10vh', height: 'fit-content', margin: '10px' }}
    />
  )
}
