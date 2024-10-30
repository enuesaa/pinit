import { useGetStory, useSetStoryInput } from '@/lib/state/story'
import { TextArea } from '@radix-ui/themes'
import { KeyboardEventHandler, useEffect, useRef } from 'react'
import styles from './BinderStoryInput.css'
import { Note } from '@/lib/api/binders'

type Props = {
  note: Note
}
export const BinderStoryInput = ({ note }: Props) => {
  const textareaRef = useRef<HTMLTextAreaElement>(null)
  const story = useGetStory()
  const setInput = useSetStoryInput()

  useEffect(() => {
    if (textareaRef.current === null) {
      return
    }
    textareaRef.current.value = story.noteInput
  }, [story.noteInput])

  const handleKeyDown: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    setInput(e.currentTarget.value)
  }

  return (
    <TextArea
      className={styles.textarea}
      ref={textareaRef}
      size='3'
      defaultValue={note.content}
      onKeyDown={handleKeyDown}
    />
  )
}
