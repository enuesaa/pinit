import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { SpeakToTextDialog } from '@/components/audio/SpeakToTextDialog'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <SpeakToTextDialog />
      </Main>
    </>
  )
}
