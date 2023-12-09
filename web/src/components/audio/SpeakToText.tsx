import { PauseIcon, TriangleRightIcon, PaperPlaneIcon } from '@radix-ui/react-icons'
import { Flex, IconButton, TextArea } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { SpeakToTextCopyButton } from './SpeakToTextCopyButton'
import { useReactMediaRecorder } from 'react-media-recorder'

export const SpeakToText = () => {
  const { startRecording, stopRecording, mediaBlobUrl } = useReactMediaRecorder({})

  const handleSend: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()  
    console.log(mediaBlobUrl)
    if (mediaBlobUrl === undefined) {
      return
    }
    const response = await fetch(mediaBlobUrl);
    const blob = await response.blob()
    const file = new File([blob], 'hello', { type: blob.type })
    const res = await fetch('/api/recog', {
      method: 'POST',
      headers: {
        'Content-Type': 'audio/wav',
      },
      body: file,
    })
    const resbody = await res.json()
    console.log(resbody?.text)
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
      <TextArea value={''} readOnly size='3' style={{ height: '300px' }} />
      <SpeakToTextCopyButton text={''} />
    </div>
  )
}
