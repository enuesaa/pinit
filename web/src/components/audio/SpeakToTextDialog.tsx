import { Button, Dialog, Flex, TextFieldInput } from '@radix-ui/themes'
import { KeyboardEventHandler, useState } from 'react'
import { SpeakToText } from './SpeakToText'

export const SpeakToTextDialog = () => {
  const [apiKey, setApiKey] = useState<string>('')
  const handleKeyDown: KeyboardEventHandler<HTMLInputElement> = (e) => {
    e.stopPropagation()
    setApiKey(e.currentTarget.value)
  }

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='ghost' style={{cursor: 'pointer'}}>speak</Button>
      </Dialog.Trigger>

      <Dialog.Content style={{ maxWidth: 450 }}>
        <Dialog.Title>Speak</Dialog.Title>
        <Dialog.Description size='2' mb='4'></Dialog.Description>

        <TextFieldInput placeholder='OpenAI API Key' onKeyUp={handleKeyDown} />

        {apiKey.length > 0 && (<SpeakToText apiKey={apiKey} />)}

        <Flex mt='4' justify='end'>
          <Dialog.Close>
            <Button variant='soft' color='gray' style={{ cursor: 'pointer' }}>Close</Button>
          </Dialog.Close>
        </Flex>

      </Dialog.Content>
    </Dialog.Root>
  )
}