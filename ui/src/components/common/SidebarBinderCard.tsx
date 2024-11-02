import { useGetBinder } from '@/lib/api/binders'
import { Card, Text } from '@radix-ui/themes'
import { SidebarBinderDeleteButton } from './SidebarBinderDeleteButton'
import { useSetStory } from '@/lib/state/story'
import { MouseEventHandler } from 'react'

type Props = {
  binderName: string
}
export const SidebarBinderCard = ({ binderName }: Props) => {
  const setStory = useSetStory()
  const binder = useGetBinder(binderName)

  const handleClick: MouseEventHandler<HTMLDivElement> = (e) => {
    e.preventDefault()
    setStory(binderName, binder.data?.content ?? '')
  }

  return (
    <Card m='3' onClick={handleClick}>
      <SidebarBinderDeleteButton binderName={binderName} />
      <Text as='div' size='2' weight='bold' mb='2'>
        {binderName}
      </Text>
      <Text as='div' color='gray' size='2'>
        {binder.data?.content ?? ''}
      </Text>
    </Card>
  )
}
