import { Button, TextArea } from '@radix-ui/themes'
import { FaCaretRight, FaMicrophone, FaTrash } from 'react-icons/fa'
import styles from './BinderStoryInputArea.css'
import { useGetAction } from '@/lib/state'
import { MouseEventHandler, useEffect, useRef } from 'react'
import { useChat } from '@/lib/api'

export const BinderStoryInputArea = () => {
  const action = useGetAction()
  const textareaRef = useRef<HTMLTextAreaElement>(null)
  const chat = useChat()

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

  const handleTrash: MouseEventHandler<HTMLSpanElement> = (e) => {
    e.preventDefault()
    if (textareaRef.current === null) {
      return
    }
    textareaRef.current.value = ''
  }

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await chat.mutateAsync(textareaRef.current?.value ?? '')
  }

  return (
    <div className={styles.main}>
      <TextArea className={styles.textarea} ref={textareaRef} />
      <span className={styles.speakButton}><FaMicrophone /></span>
      <span className={styles.trashButton} onClick={handleTrash}><FaTrash /></span>
      <Button onClick={handleClick} variant='solid' m='2' style={{ cursor: 'pointer', fontSize: '29px', lineHeight: '15px' }}>
        <FaCaretRight />
      </Button>
    </div>
  )
}
