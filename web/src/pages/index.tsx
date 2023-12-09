import { Header } from '@/components/common/Header'
import { ListBinders } from '@/components/data/ListBinders'
import { Container } from '@radix-ui/themes'
import dynamic from 'next/dynamic'

const Recorder = dynamic(() => import("../components/speak/Recorder").then(m => m.Recorder), {ssr: false})

export default function TopPage() {
  return (
    <>
      <Header />
      <Container size='4' p='4'>
        <ListBinders />
        <Recorder />
      </Container>
    </>
  )
}
