import { useListBinders } from '@/lib/api'
import { Table } from '@radix-ui/themes'
import { ListBindersTableItem } from './ListBindersTableItem'

export const ListBindersTable = () => {
  const { data: binders } = useListBinders()
  if (binders === undefined) {
    return <></>
  }

  return (
    <Table.Root>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell>Binder Id</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Content</Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {binders.map((binder, i) => (
          <ListBindersTableItem key={i} binder={binder} />
        ))}
      </Table.Body>
    </Table.Root>
  )
}
