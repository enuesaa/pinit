import { Button, Dialog, Flex } from '@radix-ui/themes'
import { Dispatch, SetStateAction } from 'react'
import { Chat } from './Chat'

type Props = {
  open: boolean
  setOpen: Dispatch<SetStateAction<boolean>>
}
export const ChatDialog = ({ open, setOpen }: Props) => {  
  return (
    <Dialog.Root open={open}>
      <Dialog.Content style={{ maxWidth: 450 }}>
        <Dialog.Title>Chat</Dialog.Title>
        <Chat />

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
