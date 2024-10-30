import { Section } from '@radix-ui/themes'
import { BinderStoryChatButton } from './BinderStoryChatButton'
import { BinderStoryInput } from './BinderStoryInput'
import { BinderStoryRecorder } from './BinderStoryRecorder'
import { BinderStorySaveButton } from './BinderStorySaveButton'
import { BinderStoryTrash } from './BinderStoryTrash'

export const BinderStory = () => {
  return (
    <Section p='1'>
      binder name
      <BinderStoryRecorder />
      <BinderStoryTrash />
      <BinderStoryInput />
      <BinderStoryChatButton />
      <BinderStorySaveButton />
    </Section>
  )
}
