import { useListActions } from '@/lib/api'
import { Section } from '@radix-ui/themes'
import { BinderStoryAction } from './BinderStoryAction'
import { BinderStoryChatButton } from './BinderStoryChatButton'
import { BinderStoryInput } from './BinderStoryInput'
import { BinderStoryOutput } from './BinderStoryOutput'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStorySaveButton } from './BinderStorySaveButton'
import { BinderStoryTrash } from './BinderStoryTrash'

export const BinderStory = () => {
  const { data: actions } = useListActions()

  return (
    <Section p='1'>
      {actions && actions.map((a, i) => <BinderStoryAction key={i} action={a} />)}
      <BinderStoryRecorder />
      <BinderStoryTrash />
      <BinderStoryInput />
      <BinderStoryChatButton />
      <BinderStorySaveButton />
      <BinderStoryOutput />
    </Section>
  )
}
