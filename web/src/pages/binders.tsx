import { Header } from '@/components/common/Header'
import { ListBinders } from '@/components/data/ListBinders'
import { Container } from '@radix-ui/themes'

export default function TopPage() {
  return (
    <>
      <Header />
      <Container size='4' p='4'>
        <ListBinders />
      </Container>
    </>
  )
}
