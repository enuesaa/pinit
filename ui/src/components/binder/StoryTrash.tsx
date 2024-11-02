import { useSetStoryInput } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaTrash } from 'react-icons/fa'

export const StoryTrash = () => {
  const setInput = useSetStoryInput()

  const handleClick: MouseEventHandler<HTMLSpanElement> = (e) => {
    e.preventDefault()
    setInput('')
  }

  return (
    <Button m='2' variant='soft' style={{ cursor: 'pointer' }} onClick={handleClick}>
      <FaTrash />
    </Button>
  )
}
