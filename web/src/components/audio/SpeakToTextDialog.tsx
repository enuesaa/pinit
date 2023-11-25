import { Button, Dialog, Flex, TextFieldInput } from '@radix-ui/themes'
import { Dispatch, KeyboardEventHandler, SetStateAction, useState } from 'react'
import { SpeakToText } from './SpeakToText'

type Props = {
  open: boolean;
  setOpen: Dispatch<SetStateAction<boolean>>
}
export const SpeakToTextDialog = ({ open, setOpen }: Props) => {
  const [apiKey, setApiKey] = useState<string>('')
  const handleKeyDown: KeyboardEventHandler<HTMLInputElement> = (e) => {
    e.stopPropagation()
    setApiKey(e.currentTarget.value)
  }

  return (
    <Dialog.Root open={open}>
      <Dialog.Content style={{ maxWidth: 450 }}>
        <Dialog.Title>Speak</Dialog.Title>
        <Dialog.Description mb='4'></Dialog.Description>

        <TextFieldInput placeholder='OpenAI API Key' onKeyUp={handleKeyDown} />

        {apiKey.length > 0 && (<SpeakToText apiKey={apiKey} />)}

        <Flex mt='6' justify='end'>
          <Dialog.Close onClick={() => setOpen(false)}>
            <Button variant='soft' color='gray' style={{ cursor: 'pointer' }}>Close</Button>
          </Dialog.Close>
        </Flex>

      </Dialog.Content>
    </Dialog.Root>
  )
}