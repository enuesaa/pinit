import { CheckIcon, CopyIcon } from '@radix-ui/react-icons'
import { MouseEventHandler, useState } from 'react'
import styles from './SpeakToTextCopyButton.css'

export const SpeakToTextCopyButton = ({ text }: { text: string }) => {
  const [clicked, setClicked] = useState<boolean>(false)

  const handleCopy: MouseEventHandler<HTMLSpanElement> = async (e) => {
    e.preventDefault()
    await globalThis.navigator.clipboard.writeText(text)
    setClicked(true)
    setTimeout(() => {
      setClicked(false)
    }, 3000)
  }

  return (
    <span onClick={handleCopy} className={styles.main}>
      {clicked ? <CheckIcon /> : <CopyIcon />}
    </span>
  )
}
