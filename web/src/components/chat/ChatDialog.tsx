import { Button, Dialog, Flex, TextFieldInput } from '@radix-ui/themes'
import { Dispatch, KeyboardEventHandler, SetStateAction, useEffect, useState } from 'react'
import { Chat } from './Chat'
import { useGetConfig } from '@/lib/api'

type Props = {
  open: boolean
  setOpen: Dispatch<SetStateAction<boolean>>
}
export const ChatDialog = ({ open, setOpen }: Props) => {
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
        <Dialog.Title>Chat</Dialog.Title>
        <Dialog.Description mt='2' mb='4'>OpenAI API Token</Dialog.Description>

        <TextFieldInput defaultValue={apiKey} onKeyUp={handleKeyDown} />

        {apiKey.length > 0 && <Chat apiKey={apiKey} />}

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
