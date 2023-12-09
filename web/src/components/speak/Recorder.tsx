import { MouseEventHandler } from 'react'
import { useReactMediaRecorder } from 'react-media-recorder'

export const Recorder = () => {
  const { status, startRecording, stopRecording, mediaBlobUrl } = useReactMediaRecorder({})

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
    <div>
      <p>{status}</p>
      <button onClick={startRecording}>Start Recording</button>
      <button onClick={stopRecording}>Stop Recording</button>
      <button onClick={handleSend}>send</button>
    </div>
  )
}
