import { useGetStory, useSetStoryInput } from '@/lib/state'
import { TextArea } from '@radix-ui/themes'
import { KeyboardEventHandler, useEffect, useRef } from 'react'
import styles from './BinderStoryInput.css'

export const BinderStoryInput = () => {
  const story = useGetStory()
  const setStoryInput = useSetStoryInput()
  const textareaRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    if (textareaRef.current === null) {
      return
    }
    textareaRef.current.value = story.input
  }, [story.input])

  const handleKeyDown: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    setStoryInput(e.currentTarget.value)
  }

  return (
    <TextArea className={styles.textarea} ref={textareaRef} size='3' onKeyDown={handleKeyDown} />
  )
}
