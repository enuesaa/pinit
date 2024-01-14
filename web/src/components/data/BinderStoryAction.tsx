import { Action } from '@/lib/api'
import { useChooseAction } from '@/lib/state'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'

type Props = {
  action: Action
}
export const BinderStoryAction = ({ action }: Props) => {
  const choose = useChooseAction()

  const handleClick: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    choose(action)
  }

  return (
    <Button m='2' style={{ cursor: 'pointer' }} onClick={handleClick}>
      {action.name}
    </Button>
  )
}
