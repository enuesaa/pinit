import { useListBinders } from '@/lib/api/binders'
import { Section } from '@radix-ui/themes'
import { SidebarBinderCard } from './SidebarBinderCard'

export const SidebarBinders = () => {
  const binders = useListBinders()
  if (binders === undefined) {
    return <></>
  }

  return (
    <Section p='2'>
      {binders.data?.items.map((binder, i) => (
        <SidebarBinderCard key={i} binderName={binder.name} />
      ))}
    </Section>
  )
}
