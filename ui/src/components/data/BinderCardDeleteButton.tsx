import { useDeleteBinder } from '@/lib/api/binders'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaTrash } from 'react-icons/fa'
import styles from './BinderCardDeleteButton.css'

type Props = {
  binderName: string
}
export const BinderCardDeleteButton = ({ binderName }: Props) => {
  const deleteBinder = useDeleteBinder(binderName)

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await deleteBinder.mutateAsync()
  }

  return (
    <Button onClick={handleClick} variant='outline' size='1' className={styles.main}>
      <FaTrash />
    </Button>
  )
}
