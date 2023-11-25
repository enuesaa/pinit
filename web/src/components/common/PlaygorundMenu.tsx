import { HoverCard, Button } from '@radix-ui/themes'
import { SpeakToTextDialog } from '../audio/SpeakToTextDialog'

export const PlaygorundMenu = () => {
  return (
    <HoverCard.Root>
      <HoverCard.Trigger>
        <Button variant='ghost' style={{ padding: '10px'}}>playground</Button>
      </HoverCard.Trigger>
      <HoverCard.Content>
        <SpeakToTextDialog />
      </HoverCard.Content>
    </HoverCard.Root>
  )
}