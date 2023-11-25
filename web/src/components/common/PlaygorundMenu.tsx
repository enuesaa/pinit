import { HoverCard, Button } from '@radix-ui/themes'
import { SpeakToTextDialog } from '../audio/SpeakToTextDialog'
import { MouseEventHandler, useState } from 'react'

export const PlaygorundMenu = () => {
  const [open, setOpen] = useState<boolean>(false)
  const handleOpenSpeakToText: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    setOpen(true)
  }

  return (
    <>
      <HoverCard.Root>
        <HoverCard.Trigger>
          <Button variant='ghost' style={{ padding: '10px' }}>playground</Button>
        </HoverCard.Trigger>
        <HoverCard.Content>
          <Button variant='ghost' style={{cursor: 'pointer'}} onClick={handleOpenSpeakToText}>speak</Button>
        </HoverCard.Content>
      </HoverCard.Root>
      <SpeakToTextDialog open={open} setOpen={setOpen} />
    </>
  )
}