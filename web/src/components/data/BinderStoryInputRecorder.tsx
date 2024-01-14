import { useRecog } from '@/lib/api'
import { useChooseAction } from '@/lib/state'
import { PaperPlaneIcon, PauseIcon, TriangleRightIcon } from '@radix-ui/react-icons'
import { IconButton } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { useReactMediaRecorder } from 'react-media-recorder'

export const BinderStoryInputRecorder = () => {
  const { status, startRecording, stopRecording, mediaBlobUrl } = useReactMediaRecorder({})
  const invokeRecogApi = useRecog()
  const choose = useChooseAction()

  const handleSend: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    if (mediaBlobUrl === undefined || mediaBlobUrl === null) {
      return
    }
    const text = await invokeRecogApi.mutateAsync(mediaBlobUrl)
    console.log(text)
    // TODO change hook because recorder is not action.
    choose({ id: 0, name: 'record', template: text })
  }

  const handleStart: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    startRecording()
  }

  if (status === 'recording') {
    return (
      <IconButton variant='soft' style={{ cursor: 'pointer' }} onClick={stopRecording}>
        <PauseIcon />
      </IconButton>
    )
  }

  if (status === 'stopped') {
    return (
      <IconButton variant='soft' style={{ cursor: 'pointer', width: '33%' }} onClick={handleSend}>
        <PaperPlaneIcon />
      </IconButton>
    )
  }

  return (
    <IconButton variant='soft' style={{ cursor: 'pointer' }} onClick={handleStart}>
      <TriangleRightIcon />
    </IconButton>
  )
}
