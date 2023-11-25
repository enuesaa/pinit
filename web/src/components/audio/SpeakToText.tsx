import useWhisper from '@chengsokdara/use-whisper'
import { CodeBlock } from '@enuesaa/fileslook'

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
  
  return (
    <div>
      <CodeBlock className='md'>{transcript.text}</CodeBlock>
      <button onClick={() => startRecording()}>Start</button>
      <button onClick={() => stopRecording()}>Stop</button>
    </div>
  )
}
