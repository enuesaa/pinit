import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { MagicWandIcon } from '@radix-ui/react-icons'
import { IconButton, TextArea } from '@radix-ui/themes'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <TextArea />
        <IconButton>
          <MagicWandIcon />
        </IconButton>
      </Main>
    </>
  )
}
