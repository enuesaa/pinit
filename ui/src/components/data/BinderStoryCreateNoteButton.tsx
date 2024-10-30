import { useCreateBinderNote } from '@/lib/api/binders'
import { useSetNoteName } from '@/lib/state/story'
import { Button } from '@radix-ui/themes'
import { MouseEventHandler, useEffect } from 'react'

type Props = {
  name: string
}
export const BinderStoryCreateNoteButton = ({ name }: Props) => {
  const createNote = useCreateBinderNote(name)
  const setNoteName = useSetNoteName()

  const handleClick: MouseEventHandler<HTMLButtonElement> = (e) => {
    e.preventDefault()
    createNote.mutate({})
  }

  useEffect(() => {
    if (createNote.data?.name !== undefined) {
      setNoteName(createNote.data.name)
    }
  }, [createNote.data?.name])

  return (
    <Button onClick={handleClick}>
      Create
    </Button>
  )
}
