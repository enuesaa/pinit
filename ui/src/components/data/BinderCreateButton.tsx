import { useCreateBinder } from '@/lib/api/binders'
import { useSetBinderName } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler, useEffect } from 'react'

export const BinderCreateButton = () => {
  const createBinder = useCreateBinder()
  const setBinderName = useSetBinderName()

  const handleClick: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    createBinder.mutate({})
  }

  useEffect(() => {
    if (createBinder.data?.name !== undefined) {
      setBinderName(createBinder.data.name)
    }
  }, [createBinder.data?.name])

  return (
    <Button onClick={handleClick} style={{ width: '90%', display: 'block', margin: '0 auto', cursor: 'pointer' }}>
      Create
    </Button>
  )
}
