import { useRecog } from '@/lib/api/recog'
import { useGetStory, useSetStoryInput } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler, useEffect } from 'react'
import { FaMicrophone } from 'react-icons/fa'
import { useReactMediaRecorder } from 'react-media-recorder'

export const StoryRecorder = () => {
  const { status, startRecording, stopRecording, mediaBlobUrl } = useReactMediaRecorder({})
  const invokeRecogApi = useRecog()
  const story = useGetStory()
  const setInput = useSetStoryInput()

  useEffect(() => {
    if (mediaBlobUrl === undefined || mediaBlobUrl === null) {
      return
    }
    ;(async () => {
      const text = await invokeRecogApi.mutateAsync(mediaBlobUrl)
      setInput(`${story.content}\n\n${text}`)
    })()
  }, [mediaBlobUrl])

  const handleStart: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    startRecording()
  }

  if (status === 'recording') {
    return (
      <Button m='2' variant='soft' onClick={stopRecording}>
        stop
      </Button>
    )
  }

  return (
    <Button m='2' variant='soft' onClick={handleStart}>
      <FaMicrophone />
    </Button>
  )
}
