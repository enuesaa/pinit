import { useListBinderNotes } from '@/lib/api/binders'
import { Button, Dialog } from '@radix-ui/themes'
import { BinderStoryInputReadonly } from './BinderStoryInputReadonly'

type Props = {
  name: string
}
export const HistoryDialog = ({ name }: Props) => {
  const notes = useListBinderNotes(name)

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='surface'>History</Button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth='90%'>
        <Dialog.Title>History</Dialog.Title>

        {notes?.data?.items.map(v => (
          <BinderStoryInputReadonly note={v} key={v.name} />
        ))}

        <Dialog.Close>
          <Button variant='surface'>Cancel</Button>
        </Dialog.Close>
      </Dialog.Content>
    </Dialog.Root>
  )
}
