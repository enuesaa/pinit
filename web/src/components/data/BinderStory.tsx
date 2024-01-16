import { useListActions } from '@/lib/api'
import { Section } from '@radix-ui/themes'
import { BinderStoryAction } from './BinderStoryAction'
import { BinderStoryInputArea } from './BinderStoryInputArea'

export const BinderStory = () => {
  const { data: actions } = useListActions()

  return (
    <Section p='1'>
      {actions && actions.map((a, i) => <BinderStoryAction key={i} action={a} />)}
      <BinderStoryInputArea />
    </Section>
  )
}
