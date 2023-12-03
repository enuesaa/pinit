import { useChat } from '@/lib/openai'
import { Button, TextArea } from '@radix-ui/themes'
import { MouseEventHandler, useRef, useState } from 'react'
import { SpeakToTextCopyButton } from '../audio/SpeakToTextCopyButton'

type Props = {
  apiKey: string
}
export const Chat = ({ apiKey }: Props) => {
  const { mutateAsync: chat } = useChat(apiKey)
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
