import { useListActions } from '@/lib/api'
import { Section } from '@radix-ui/themes'
import { BinderStoryAction } from './BinderStoryAction'
import { BinderStoryInputArea } from './BinderStoryInputArea'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStoryTrash } from './BinderStoryTrash'

export const BinderStory = () => {
  const { data: actions } = useListActions()

  return (
    <Section p='1'>
      {actions && actions.map((a, i) => <BinderStoryAction key={i} action={a} />)}
      <BinderStoryRecorder />
      <BinderStoryTrash />
      <BinderStoryInputArea />
    </Section>
  )
}
