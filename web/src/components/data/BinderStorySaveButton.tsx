import { useCreateBinder } from '@/lib/api'
import { useGetStory } from '@/lib/state'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaCheck } from 'react-icons/fa'
import styles from './BinderStoryChatButton.css'

export const BinderStorySaveButton = () => {
  const story = useGetStory()
  const createBinder = useCreateBinder()

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const name = story.input.slice(0, 100)
    const content = story.output
    await createBinder.mutateAsync({ name, content })
  }

  return (
    <Button variant='surface' m='2' className={styles.main} onClick={handleClick} style={{ fontSize: '20px' }}>
      <FaCheck />
    </Button>
  )
}
