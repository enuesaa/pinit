import { Header } from '@/components/common/Header'
import { BinderStory } from '@/components/data/BinderStory'
import { Container } from '@radix-ui/themes'

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
