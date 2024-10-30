import { useCreateBinder } from '@/lib/api/binders'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'

export const BinderCreateBtn = () => {
  const createBinder = useCreateBinder()

  const handleClick: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    createBinder.mutate()
  }

  return (
    <Button onClick={handleClick}>
      Create
    </Button>
  )
}
