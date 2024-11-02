import { TextArea } from '@radix-ui/themes'

type Props = {
  content: string
}
export const StoryInputReadonly = ({ content }: Props) => {
  return (
    <TextArea
      className='min-h-[80vh] p-[10px] outline-none'
      size='3'
      defaultValue={content}
      readOnly
      style={{ minHeight: '10vh', height: 'fit-content', margin: '10px' }}
    />
  )
}
