import useWhisper from '@chengsokdara/use-whisper'
import { CodeBlock } from '@enuesaa/fileslook'
import { Button, Flex, IconButton, Separator } from '@radix-ui/themes';
import { PauseIcon, TriangleRightIcon } from '@radix-ui/react-icons'
import { MouseEventHandler } from 'react';

type Props = {
  apiKey: string;
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
    <div>
      <Flex gap='3' mt='6' mb='4' justify='center'>
        <IconButton variant='soft' style={{width: '45%'}} onClick={handleStart}><TriangleRightIcon /></IconButton>
        <IconButton variant='soft'  style={{width: '45%'}} onClick={handleStop}><PauseIcon /></IconButton>
      </Flex>
      <CodeBlock className='md'>{transcript.text ?? ''}</CodeBlock>
    </div>
  )
}
