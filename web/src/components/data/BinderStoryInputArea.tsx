import { TextArea } from '@radix-ui/themes'
import { FaMicrophone } from 'react-icons/fa'
import styles from './BinderStoryInputArea.css'
import { useGetAction } from '@/lib/state'
import { useEffect, useRef } from 'react'

export const BinderStoryInputArea = () => {
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
    <div className={styles.main}>
      <TextArea className={styles.textarea} ref={textareaRef} />
      <span className={styles.speakButton}><FaMicrophone /></span>
    </div>
  )
}
