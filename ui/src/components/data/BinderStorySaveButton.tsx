import { useUpdateBinder } from '@/lib/api/binders'
import { useGetStory } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaCheck } from 'react-icons/fa'
import styles from './BinderStoryChatButton.css'

type Props = {
  binderName: string
}
export const BinderStorySaveButton = ({ binderName }: Props) => {
  const story = useGetStory()
  const updateBinder = useUpdateBinder(binderName)

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await updateBinder.mutateAsync({ content: story.content })
  }

  return (
    <Button m='2' className={styles.main} onClick={handleClick} style={{ fontSize: '20px' }}>
      <FaCheck />
    </Button>
  )
}
