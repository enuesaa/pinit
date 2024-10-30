import { useChat } from '@/lib/api/chat'
import { useGetStory } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaCaretRight, FaHourglassStart } from 'react-icons/fa'
import { IoWarningOutline } from 'react-icons/io5'
import styles from './BinderStoryChatButton.css'

export const BinderStoryChatButton = () => {
  const chat = useChat()
  const story = useGetStory()
  // const setStoryOutput = useSetStoryOuptut()

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const output = await chat.mutateAsync({message: story.noteInput})
    // setStoryOutput(output.message)
  }

  if (chat.isLoading) {
    return (
      <Button variant='surface' m='2' className={styles.main} onClick={handleClick}>
        <FaHourglassStart />
      </Button>
    )
  }

  if (chat.isError) {
    return (
      <Button variant='surface' m='2' className={styles.main} onClick={handleClick}>
        <IoWarningOutline />
      </Button>
    )
  }

  return (
    <Button variant='surface' m='2' className={styles.main} onClick={handleClick}>
      <FaCaretRight />
    </Button>
  )
}
