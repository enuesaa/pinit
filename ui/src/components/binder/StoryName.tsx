import { useRenameBinder } from '@/lib/api/binder-rename'
import { Button, TextField } from '@radix-ui/themes'
import { FormEventHandler, MouseEventHandler, useState } from 'react'

type Props = {
  name: string
}
export const StoryName = ({ name }: Props) => {
  const [modify, setModify] = useState(false)
  const renameBinder = useRenameBinder(name)

  if (!modify) {
    const handleClick: MouseEventHandler<HTMLDivElement> = (e) => {
      e.preventDefault()
      setModify(true)
    }
  
    return <div onClick={handleClick} className='my-5 font-semibold text-xl'>{name}</div>
  }

  const handleSave: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const data = {
      newName: e.currentTarget.newName.value as string,
    }
    if (data.newName === name) {
      setModify(false)
      return;
    }
    renameBinder.mutate(data)
    // データ更新にタイムラグがある
    setTimeout(() => location.reload(), 1000)
    setModify(false)
  }

  return (
    <form onSubmit={handleSave} className='my-5'>
      <TextField.Root 
        name='newName'
        defaultValue={name}
        pattern='[a-z0-9\-]+'
        className='w-[90%] max-w-[500px] inline-block font-semibold text-xl'
      />
      <Button type='submit' variant='surface' className='inline-block mx-2'>OK</Button>
    </form>
  )
}