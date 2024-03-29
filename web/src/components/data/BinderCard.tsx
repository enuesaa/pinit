import { Binder, useListBinderNotes } from '@/lib/api'
import { Card, Text } from '@radix-ui/themes'
import { BinderCardDeleteButton } from './BinderCardDeleteButton'

type Props = {
  binder: Binder
}
export const BinderCard = ({ binder }: Props) => {
  const { data: notes } = useListBinderNotes(binder.id)
  const latestNote = (notes !== undefined ? notes.at(0) : null) ?? null

  return (
    <Card m='3'>
      <BinderCardDeleteButton binderId={binder.id} />
      <Text as='div' size='2' weight='bold' mb='2'>
        {binder.name}
      </Text>
      <Text as='div' color='gray' size='2'>
        {latestNote !== null ? latestNote.content : ''}
      </Text>
    </Card>
  )
}
