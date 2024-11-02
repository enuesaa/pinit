import { useDeleteBinder } from '@/lib/api/binders'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaTrash } from 'react-icons/fa'

type Props = {
  binderName: string
}
export const SidebarBinderDeleteButton = ({ binderName }: Props) => {
  const deleteBinder = useDeleteBinder(binderName)

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await deleteBinder.mutateAsync()
  }

  return (
    <Button onClick={handleClick} variant='outline' className='px-2 py-1 absolute right-5'>
      <FaTrash />
    </Button>
  )
}
