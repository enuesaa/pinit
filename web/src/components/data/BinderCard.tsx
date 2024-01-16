import { Binder, useListBinderNotes } from '@/lib/api'
import { Card, Text } from '@radix-ui/themes'

type Props = {
  binder: Binder
}
export const BinderCard = ({ binder }: Props) => {
  const { data: notes } = useListBinderNotes(binder.id)
  const latestNote = (notes !== undefined ? notes.at(0) : null) ?? null

  return (
    <Card>
      <Text as='div' size='2' weight='bold'>
        {binder.name}
      </Text>
      <Text as='div' color='gray' size='2'>
        {latestNote !== null ? latestNote.content : ''}
      </Text>
    </Card>
  )
}
