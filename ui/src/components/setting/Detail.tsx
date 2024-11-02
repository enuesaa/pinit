import { useGetSetting, useUpdateSetting } from '@/lib/api/setting'
import { Button } from '@radix-ui/themes'
import { FormEventHandler } from 'react'

export const Detail = () => {
  const setting = useGetSetting()
  const updateSetting = useUpdateSetting()

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()

    const data = {
      openaiToken: e.currentTarget.openaiToken.value as string,
    }
    updateSetting.mutate(data)
  }

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input type='text' name='openaiToken' defaultValue={setting.data?.openaiToken} />
        <Button>Save</Button>
      </form>
    </div>
  )
}
