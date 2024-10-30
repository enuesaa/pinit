import { useGetStory, useSetStoryInput } from '@/lib/state/story'
import { TextArea } from '@radix-ui/themes'
import { KeyboardEventHandler, useEffect, useRef } from 'react'
import styles from './BinderStoryInput.css'

export const BinderStoryInput = () => {
  const textareaRef = useRef<HTMLTextAreaElement>(null)
  const story = useGetStory()
  const setInput = useSetStoryInput()

  useEffect(() => {
    if (textareaRef.current === null) {
      return
    }
    textareaRef.current.value = story.noteContent
  }, [story.noteContent])

  const handleKeyUp: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    setInput(e.currentTarget.value)
  }

  return (
    <TextArea
      className={styles.textarea}
      ref={textareaRef}
      size='3'
      defaultValue={story.noteContent}
      onKeyUp={handleKeyUp}
    />
  )
}
