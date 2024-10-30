import { Binder, useListBinderNotes } from '@/lib/api/binders'
import { Card, Text } from '@radix-ui/themes'
import { BinderCardDeleteButton } from './BinderCardDeleteButton'
import { useSetBinderName } from '@/lib/state/story'
import { MouseEventHandler } from 'react'

type Props = {
  binder: Binder
}
export const BinderCard = ({ binder }: Props) => {
  const setBinderName = useSetBinderName()
  const { data: notes } = useListBinderNotes(binder.name)
  const latestNote = (notes !== undefined ? notes?.items.at(0) : null) ?? null

  const handleClick: MouseEventHandler<HTMLDivElement> = (e) => {
    e.preventDefault()
    setBinderName(binder.name)
  }

  return (
    <Card m='3' onClick={handleClick}>
      <BinderCardDeleteButton binderName={binder.name} />
      <Text as='div' size='2' weight='bold' mb='2'>
        {binder.name}
      </Text>
      <Text as='div' color='gray' size='2'>
        {latestNote !== null ? latestNote.content : ''}
      </Text>
    </Card>
  )
}
