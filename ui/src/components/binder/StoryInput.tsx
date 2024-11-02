import { useUpdateBinder } from '@/lib/api/binders'
import { useGetStory, useSetStoryInput } from '@/lib/state/story'
import { TextArea } from '@radix-ui/themes'
import { FocusEventHandler, KeyboardEventHandler, useEffect, useRef } from 'react'
import { FaCheck } from 'react-icons/fa'

type Props = {
  name: string
}
export const StoryInput = ({ name }: Props) => {
  const textareaRef = useRef<HTMLTextAreaElement>(null)
  const story = useGetStory()
  const setInput = useSetStoryInput()
  const updateBinder = useUpdateBinder(name)

  useEffect(() => {
    if (textareaRef.current === null) {
      return
    }
    textareaRef.current.value = story.content
  }, [story.content])

  const handleKeyUp: KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    setInput(e.currentTarget.value)
  }

  const handleSave: FocusEventHandler<HTMLTextAreaElement> = (e) => {
    updateBinder.mutate({ content: story.content })
    setTimeout(() => updateBinder.reset(), 1000)
  }

  return (
    <>
      <TextArea
        className='min-h-[80vh] p-[10px] outline-none'
        ref={textareaRef}
        size='3'
        defaultValue={story.content}
        onKeyUp={handleKeyUp}
        onBlur={handleSave}
      />
      {updateBinder.isSuccess && (
        <div className='absolute top-5 right-5'><FaCheck /></div>
      )}
    </>
  )
}
