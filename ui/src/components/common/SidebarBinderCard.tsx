import { useGetBinder } from '@/lib/api/binders'
import { useSetStory } from '@/lib/state/story'
import { Card, Text } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { SidebarBinderDeleteButton } from './SidebarBinderDeleteButton'

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
    <Card onClick={handleClick} className='m-3'>
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
