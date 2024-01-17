import { useDeleteBinder } from '@/lib/api'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaTrash } from 'react-icons/fa'
import styles from './BinderCardDeleteButton.css'

type Props = {
  binderId: number
}
export const BinderCardDeleteButton = ({ binderId }: Props) => {
  const deleteBinder = useDeleteBinder()

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await deleteBinder.mutateAsync(binderId)
  }

  return (
    <Button onClick={handleClick} variant='outline' size='1' className={styles.main}>
      <FaTrash />
    </Button>
  )
}