import { useListBinders } from '@/lib/api/binders'
import { Section } from '@radix-ui/themes'
import { BinderCard } from './BinderCard'

export const ListBinders = () => {
  const { data: binders } = useListBinders()
  if (binders === undefined) {
    return <></>
  }

  return (
    <Section p='2'>
      {binders.items.map((binder, i) => (
        <BinderCard key={i} binder={binder} />
      ))}
    </Section>
  )
}
