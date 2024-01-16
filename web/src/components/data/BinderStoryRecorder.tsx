import { useRecog } from '@/lib/api'
import { useSetStoryInput } from '@/lib/state'
import { PauseIcon } from '@radix-ui/react-icons'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler, useEffect } from 'react'
import { FaMicrophone } from 'react-icons/fa'
import { useReactMediaRecorder } from 'react-media-recorder'

export const BinderStoryRecorder = () => {
  const { status, startRecording, stopRecording, mediaBlobUrl } = useReactMediaRecorder({})
  const invokeRecogApi = useRecog()
  const setInput = useSetStoryInput()

  useEffect(() => {
    if (mediaBlobUrl === undefined || mediaBlobUrl === null) {
      return
    }
    (async() => {
      const text = await invokeRecogApi.mutateAsync(mediaBlobUrl)
      setInput(text)
    })()
  }, [mediaBlobUrl])

  const handleStart: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    startRecording()
  }

  if (status === 'recording') {
    return (
      <Button m='2' variant='soft' style={{cursor: 'pointer'}} onClick={stopRecording}>
        <PauseIcon />
      </Button>
    )
  }

  return (
    <Button m='2' variant='soft' style={{cursor: 'pointer'}} onClick={handleStart}>
      <FaMicrophone />
    </Button>
  )
}
