import { useRecog } from '@/lib/api'
import { PaperPlaneIcon, PauseIcon, TriangleRightIcon } from '@radix-ui/react-icons'
import { Flex, IconButton, TextArea } from '@radix-ui/themes'
import { MouseEventHandler, useState } from 'react'
import { useReactMediaRecorder } from 'react-media-recorder'
import { SpeakToTextCopyButton } from './SpeakToTextCopyButton'

export const SpeakToText = () => {
  const { startRecording, stopRecording, mediaBlobUrl } = useReactMediaRecorder({})
  const invokeRecogApi = useRecog()
  const [text, setText] = useState('')

  const handleSend: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    if (mediaBlobUrl === undefined || mediaBlobUrl === null) {
      return
    }
    const text = await invokeRecogApi.mutateAsync(mediaBlobUrl)
    setText(text)
  }

  return (
    <div style={{ position: 'relative' }}>
      <Flex gap='3' mt='6' mb='4' justify='center'>
        <IconButton variant='soft' style={{ cursor: 'pointer', width: '33%' }} onClick={startRecording}>
          <TriangleRightIcon />
        </IconButton>
        <IconButton variant='soft' style={{ cursor: 'pointer', width: '33%' }} onClick={stopRecording}>
          <PauseIcon />
        </IconButton>
        <IconButton variant='soft' style={{ cursor: 'pointer', width: '33%' }} onClick={handleSend}>
          <PaperPlaneIcon />
        </IconButton>
      </Flex>
      <TextArea value={text} readOnly size='3' style={{ height: '300px' }} />
      <SpeakToTextCopyButton text={text} />
    </div>
  )
}
