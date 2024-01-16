import { useChat } from '@/lib/api'
import { useGetAction } from '@/lib/state'
import { Button, Text, TextArea } from '@radix-ui/themes'
import { MouseEventHandler, useEffect, useRef } from 'react'
import { FaCaretRight } from 'react-icons/fa'
import styles from './BinderStoryInputArea.css'

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

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await chat.mutateAsync(textareaRef.current?.value ?? '')
  }

  return (
    <div className={styles.main}>
      <TextArea className={styles.textarea} ref={textareaRef} size='3' />
      <Button
        onClick={handleClick}
        variant='solid'
        m='2'
        style={{ cursor: 'pointer', fontSize: '29px', lineHeight: '15px' }}
      >
        <FaCaretRight />
      </Button>

      <Text as='p' m='6'>
        {chat.data}
      </Text>
    </div>
  )
}
