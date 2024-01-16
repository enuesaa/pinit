import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import styles from './BinderStoryChatButton.css'
import { useGetStory } from '@/lib/state'
import { useCreateBinder } from '@/lib/api'
import { GiSaveArrow } from 'react-icons/gi'

export const BinderStorySaveButton = () => {
  const story = useGetStory()
  const createBinder = useCreateBinder()

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    await createBinder.mutateAsync(story.input.slice(0, 100))
  }

  return (
    <Button variant='surface' m='2' className={styles.main} onClick={handleClick} style={{ fontSize: '20px' }}>
      <GiSaveArrow />
    </Button>
  )
}