import { Button, Dialog, Flex } from '@radix-ui/themes'
import { Dispatch, SetStateAction } from 'react'
import { SpeakToText } from './SpeakToText'

type Props = {
  open: boolean
  setOpen: Dispatch<SetStateAction<boolean>>
}
export const SpeakToTextDialog = ({ open, setOpen }: Props) => {
  return (
    <Dialog.Root open={open}>
      <Dialog.Content style={{ maxWidth: 450 }}>
        <Dialog.Title>Speak c</Dialog.Title>
        <Dialog.Description mt='2' mb='4'>
          OpenAI API Token
        </Dialog.Description>

        <SpeakToText />
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
