import { Button, Dialog, Flex, TextFieldInput } from '@radix-ui/themes'
import { Dispatch, KeyboardEventHandler, SetStateAction, useEffect, useState } from 'react'
import { SpeakToText } from './SpeakToText'
import { useGetConfig } from '@/lib/api'

type Props = {
  open: boolean
  setOpen: Dispatch<SetStateAction<boolean>>
}
export const SpeakToTextDialog = ({ open, setOpen }: Props) => {
  const { data: config } = useGetConfig()
  const [apiKey, setApiKey] = useState<string>('')
  const handleKeyDown: KeyboardEventHandler<HTMLInputElement> = (e) => {
    e.stopPropagation()
    setApiKey(e.currentTarget.value)
  }
  useEffect(() => {
    setApiKey(config?.token ?? '')
  }, [config])

  return (
    <Dialog.Root open={open}>
      <Dialog.Content style={{ maxWidth: 450 }}>
        <Dialog.Title>Speak</Dialog.Title>
        <Dialog.Description mt='2' mb='4'>OpenAI API Token</Dialog.Description>

        <TextFieldInput defaultValue={apiKey} onKeyUp={handleKeyDown} />

        {apiKey.length > 0 && <SpeakToText apiKey={apiKey} />}

        <Flex mt='6' justify='end'>
          <Dialog.Close onClick={() => setOpen(false)}>
            <Button variant='soft' color='gray' style={{ cursor: 'pointer' }}>
              Close
            </Button>
          </Dialog.Close>
        </Flex>
      </Dialog.Content>
    </Dialog.Root>
  )
}
