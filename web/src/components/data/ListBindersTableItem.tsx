import { useListBinderNotes } from '@/lib/api'
import { Binder } from '@/lib/schema'
import { Table } from '@radix-ui/themes'

type Props = {
  binder: Binder
}
export const ListBindersTableItem = ({ binder }: Props) => {
  const {data: notes} = useListBinderNotes(binder.id)
  const latestNote = (notes !== undefined ? notes.at(0) : null) ?? null

  return (
    <Table.Row>
      <Table.Cell>{binder.id}</Table.Cell>
      <Table.Cell>{binder.name}</Table.Cell>
      <Table.Cell>{latestNote !== null ? latestNote.content : ''}</Table.Cell>
    </Table.Row>
  )
}
