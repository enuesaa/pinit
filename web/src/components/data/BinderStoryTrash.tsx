import { Button } from '@radix-ui/themes'
import { FaTrash } from 'react-icons/fa'

export const BinderStoryTrash = () => {
  // const handleTrash: MouseEventHandler<HTMLSpanElement> = (e) => {
  //   e.preventDefault()
  //   if (textareaRef.current === null) {
  //     return
  //   }
  //   textareaRef.current.value = ''
  // }
  return (
    <Button m='2' variant='soft' style={{cursor: 'pointer'}}>
      <FaTrash />
    </Button>
  )
}