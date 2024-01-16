import { useGetAction } from '@/lib/state'
import { Button, Text, TextArea } from '@radix-ui/themes'
import { MouseEventHandler, useEffect, useRef } from 'react'
import styles from './BinderStoryInput.css'

export const BinderStoryInput = () => {
  const action = useGetAction()
  const textareaRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    if (textareaRef.current === null) {
      return
    }
    if (textareaRef.current.value === '') {
      textareaRef.current.defaultValue = action.template
      return
    }
    const oldValue = textareaRef.current.value
    textareaRef.current.value = `${action.template} \n\n`
    textareaRef.current.value += oldValue
  }, [action])

  return (
    <TextArea className={styles.textarea} ref={textareaRef} size='3' />
  )
}
