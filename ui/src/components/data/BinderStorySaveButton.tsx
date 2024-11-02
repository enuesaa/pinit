import { useCreateBinderNote } from '@/lib/api/binders'
import { useGetStory, useSetNote } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler, useEffect } from 'react'
import { FaCheck } from 'react-icons/fa'
import styles from './BinderStoryChatButton.css'

type Props = {
  binderName: string
  name: string
}
export const BinderStorySaveButton = ({ binderName, name }: Props) => {
  const story = useGetStory()
  const setNote = useSetNote()
  const createNote = useCreateBinderNote(binderName)

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const content = story.noteContent
    await createNote.mutateAsync({ content })
  }

  useEffect(() => {
    if (createNote.data?.name !== undefined) {
      setNote(createNote.data.name, '')
    }
  }, [createNote.data?.name])

  return (
    <Button m='2' className={styles.main} onClick={handleClick} style={{ fontSize: '20px' }}>
      <FaCheck />
    </Button>
  )
}
