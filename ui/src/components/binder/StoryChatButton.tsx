import { useChat } from '@/lib/api/chat'
import { useGetStory, useSetStoryInput } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaHourglassStart } from 'react-icons/fa'
import { IoWarningOutline, IoChatbubbleEllipsesOutline } from 'react-icons/io5'

export const StoryChatButton = () => {
  const chat = useChat()
  const story = useGetStory()
  const setStoryInput = useSetStoryInput()

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const output = await chat.mutateAsync({ message: story.content })
    setStoryInput(story.content + '\n\n' + output.message)
  }

  if (chat.isLoading) {
    return (
      <Button variant='soft' className='px-3 py-1 m-2 text-xl' onClick={handleClick}>
        <FaHourglassStart />
      </Button>
    )
  }

  if (chat.isError) {
    return (
      <Button variant='soft' className='px-3 py-1 m-2 text-xl' onClick={handleClick}>
        <IoWarningOutline />
      </Button>
    )
  }

  return (
    <Button variant='soft' className='px-3 py-1 text-xl m-2' onClick={handleClick}>
      <IoChatbubbleEllipsesOutline />
    </Button>
  )
}
