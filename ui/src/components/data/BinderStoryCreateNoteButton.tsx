import { useCreateBinderNote } from '@/lib/api/binders'
import { useSetNote } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler, useEffect } from 'react'

type Props = {
  name: string
}
export const BinderStoryCreateNoteButton = ({ name }: Props) => {
  const createNote = useCreateBinderNote(name)
  const setNote = useSetNote()

  const handleClick: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    createNote.mutate({})
  }

  useEffect(() => {
    if (createNote.data?.name !== undefined) {
      setNote(createNote.data.name, '')
    }
  }, [createNote.data?.name])

  return (
    <Button onClick={handleClick} m='2' mt='6' variant='surface' style={{ display: 'block', width: '100px', cursor: 'pointer' }}>
      Add
    </Button>
  )
}
