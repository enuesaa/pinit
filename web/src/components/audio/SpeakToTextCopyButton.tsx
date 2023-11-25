import { CheckIcon, CopyIcon } from '@radix-ui/react-icons'
import { MouseEventHandler, useState } from 'react'

export const SpeakToTextCopyButton = ({ text }: { text: string }) => {
  const [clicked, setClicked] = useState<boolean>(false)

  const handleCopy: MouseEventHandler<HTMLSpanElement> = async (e) => {
    e.preventDefault()
    await globalThis.navigator.clipboard.writeText(text)
    setClicked(true)
    setTimeout(() => { setClicked(false) }, 3000)
  }

  return (
    <span onClick={handleCopy} style={{ marginTop: '-25px', right: '10px', position: 'absolute', cursor: 'pointer' }}>
      {clicked ? (<CheckIcon />): (<CopyIcon />)}
    </span>
  )
}