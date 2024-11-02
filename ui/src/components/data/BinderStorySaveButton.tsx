import { Note, useUpdateBinderNote } from '@/lib/api/binders'
import { useGetStory } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaCheck } from 'react-icons/fa'
import styles from './BinderStoryChatButton.css'

type Props = {
  binderName: string
  name: string
}
export const BinderStorySaveButton = ({ binderName, name }: Props) => {
  const story = useGetStory()
  const updateNote = useUpdateBinderNote(binderName, name)

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const content = story.noteContent
    await updateNote.mutateAsync({ content })
  }

  return (
    <Button m='2' className={styles.main} onClick={handleClick} style={{ fontSize: '20px' }}>
      <FaCheck />
    </Button>
  )
}
