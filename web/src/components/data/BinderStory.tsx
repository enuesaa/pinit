import { useListActions } from '@/lib/api'
import { Section } from '@radix-ui/themes'
import { BinderStoryAction } from './BinderStoryAction'
import { BinderStoryInput } from './BinderStoryInput'
import { BinderStoryOutput } from './BinderStoryOutput'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStoryTrash } from './BinderStoryTrash'
import { BinderStoryChatButton } from './BinderStoryChatButton'

export const BinderStory = () => {
  const { data: actions } = useListActions()

  return (
    <Section p='1'>
      {actions && actions.map((a, i) => <BinderStoryAction key={i} action={a} />)}
      <BinderStoryRecorder />
      <BinderStoryTrash />
      <BinderStoryInput />
      <BinderStoryChatButton />
      <BinderStoryOutput />
    </Section>
  )
}
