import { useChat } from '@/lib/api'
import { Button, TextArea } from '@radix-ui/themes'
import { MouseEventHandler, useRef, useState } from 'react'
import { SpeakToTextCopyButton } from '../audio/SpeakToTextCopyButton'

export const Chat = () => {
  const { mutateAsync: chat } = useChat()
  const textareaRef = useRef<HTMLTextAreaElement>(null)
  const [res, setRes] = useState('')

  const handleChat: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    if (textareaRef.current === null) {
      return
    }
    const message = textareaRef.current.value
    const response = await chat(message)
    setRes(response)
  }

  return (
    <>
      <TextArea size='3' ref={textareaRef} style={{ height: '300px' }} />
      <Button onClick={handleChat}>chat</Button>
      <TextArea value={res} readOnly size='3'style={{ height: '300px' }} />
      <SpeakToTextCopyButton text={res ?? ''} />
    </>
  )
}
