import { useRecog } from '@/lib/api'
import { useSetStoryInput, useSetStoryOuptut } from '@/lib/state'
import { PaperPlaneIcon, PauseIcon } from '@radix-ui/react-icons'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { FaMicrophone } from 'react-icons/fa'
import { useReactMediaRecorder } from 'react-media-recorder'

export const BinderStoryRecorder = () => {
  const { status, startRecording, stopRecording, mediaBlobUrl } = useReactMediaRecorder({})
  const invokeRecogApi = useRecog()
  const setInput = useSetStoryInput()

  const handleSend: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    if (mediaBlobUrl === undefined || mediaBlobUrl === null) {
      return
    }
    const text = await invokeRecogApi.mutateAsync(mediaBlobUrl)
    setInput(text)
  }

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

  if (status === 'stopped') {
    return (
      <Button m='2' variant='soft' style={{cursor: 'pointer'}} onClick={handleSend}>
        <PaperPlaneIcon />
      </Button>
    )
  }

  return (
    <Button m='2' variant='soft' style={{cursor: 'pointer'}} onClick={handleStart}>
      <FaMicrophone />
    </Button>
  )
}
