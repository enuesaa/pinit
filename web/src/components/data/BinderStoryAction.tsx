import { Action } from '@/lib/api'
import { useSetStoryInput } from '@/lib/state'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'

type Props = {
  action: Action
}
export const BinderStoryAction = ({ action }: Props) => {
  const setInput = useSetStoryInput()

  const handleClick: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    setInput(action.template)
  }

  return (
    <Button m='2' variant='soft' style={{ cursor: 'pointer' }} onClick={handleClick}>
      {action.name}
    </Button>
  )
}
