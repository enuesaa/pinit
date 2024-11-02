import { useListBinderNotes } from '@/lib/api/notes'
import { Button, Dialog } from '@radix-ui/themes'
import { BsThreeDots } from 'react-icons/bs'
import { StoryInputReadonly } from './StoryInputReadonly'

type Props = {
  name: string
}
export const HistoryDialog = ({ name }: Props) => {
  const notes = useListBinderNotes(name)

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='soft' m='2'>
          <BsThreeDots />
        </Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='90%'>
        <Dialog.Title>History</Dialog.Title>

        {notes?.data?.items.map((v) => <StoryInputReadonly content={v.content} key={v.name} />)}

        <Dialog.Close>
          <Button variant='surface' m='2'>
            Cancel
          </Button>
        </Dialog.Close>
      </Dialog.Content>
    </Dialog.Root>
  )
}
