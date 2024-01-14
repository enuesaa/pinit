import { Header } from '@/components/common/Header'
import { Container } from '@radix-ui/themes'
import { BinderStory } from '@/components/data/BinderStory'

export default function Page() {
  return (
    <>
      <Header />
      <Container size='4' p='4'>
        <BinderStory />
      </Container>
    </>
  )
}
