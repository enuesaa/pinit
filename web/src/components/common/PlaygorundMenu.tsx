import { Button, HoverCard } from '@radix-ui/themes'
import { MouseEventHandler, useState } from 'react'
import { SpeakToTextDialog } from '../audio/SpeakToTextDialog'
import { ChatDialog } from '../chat/ChatDialog'

export const PlaygorundMenu = () => {
  const [open, setOpen] = useState<boolean>(false)
  const [openChat, setOpenChat] = useState<boolean>(false)
  const handleOpenSpeakToText: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    setOpen(true)
  }
  const handleOpenChat: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    setOpenChat(true)
  }

  return (
    <>
      <HoverCard.Root>
        <HoverCard.Trigger>
          <Button variant='ghost' style={{ cursor: 'pointer', padding: '10px' }}>
            playground
          </Button>
        </HoverCard.Trigger>
        <HoverCard.Content>
        <Button variant='ghost' style={{ cursor: 'pointer' }} onClick={handleOpenSpeakToText}>
            speak
          </Button>
          <Button variant='ghost' style={{ cursor: 'pointer' }} onClick={handleOpenChat}>
            chat
          </Button>
        </HoverCard.Content>
      </HoverCard.Root>
      <ChatDialog open={openChat} setOpen={setOpenChat} />
      <SpeakToTextDialog open={open} setOpen={setOpen} />
    </>
  )
}
