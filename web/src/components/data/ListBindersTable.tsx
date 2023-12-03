import { useListBinders } from '@/lib/api'
import { Table } from '@radix-ui/themes'

export const ListBindersTable = () => {
  const {data: binders} = useListBinders()
  if (binders === undefined) {
    return (<></>)
  }

  return (
    <Table.Root>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell>Binder Id</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Name</Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {binders.map((binder, i) => (
          <Table.Row key={i}>
            <Table.Cell>{binder.id}</Table.Cell>
            <Table.Cell>{binder.name}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  )
}