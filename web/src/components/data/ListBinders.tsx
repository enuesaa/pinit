import { Heading } from '@radix-ui/themes'
import { ListBindersTable } from './ListBindersTable'

export const ListBinders = () => {
  return (
    <>
      <Heading as='h2' size='5'>Binders</Heading>
      <ListBindersTable />
    </>
  )
}