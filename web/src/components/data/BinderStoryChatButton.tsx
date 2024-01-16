import { useChat } from '@/lib/api'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaCaretRight } from 'react-icons/fa'
import styles from './BinderStoryChatButton.css'
import { useGetStory, useSetStoryOuptut } from '@/lib/state'

export const BinderStoryChatButton = () => {
  const chat = useChat()
  const story = useGetStory()
  const setStoryOutput = useSetStoryOuptut()

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const output = await chat.mutateAsync(story.input)
    setStoryOutput(output)
  }

  return (
  <Button variant='surface' m='2' className={styles.main} onClick={handleClick}>
    <FaCaretRight />
  </Button>
  )
}