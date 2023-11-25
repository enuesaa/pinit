import useWhisper from '@chengsokdara/use-whisper'
import { PauseIcon, TriangleRightIcon } from '@radix-ui/react-icons'
import { Flex, IconButton, TextArea } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import { SpeakToTextCopyButton } from './SpeakToTextCopyButton'

type Props = {
  apiKey: string
}
export const SpeakToText = ({ apiKey }: Props) => {
  const { transcript, startRecording, stopRecording } = useWhisper({
    apiKey,
    streaming: true,
    timeSlice: 1000,
    whisperConfig: {
      language: 'ja',
    },
  })

  const handleStart: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    startRecording()
  }
  const handleStop: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    stopRecording()
  }

  return (
    <div style={{ position: 'relative' }}>
      <Flex gap='3' mt='6' mb='4' justify='center'>
        <IconButton variant='soft' style={{ width: '45%' }} onClick={handleStart}>
          <TriangleRightIcon />
        </IconButton>
        <IconButton variant='soft' style={{ width: '45%' }} onClick={handleStop}>
          <PauseIcon />
        </IconButton>
      </Flex>
      <TextArea value={transcript.text ?? ''} readOnly size='3' style={{ height: '300px' }} />
      <SpeakToTextCopyButton text={transcript.text ?? ''} />
    </div>
  )
}
