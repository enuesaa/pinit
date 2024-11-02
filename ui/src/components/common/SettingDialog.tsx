import { useGetSetting, useUpdateSetting } from '@/lib/api/setting'
import { Button, Dialog, TextField } from '@radix-ui/themes'
import { FormEventHandler } from 'react'
import { MdSettings } from 'react-icons/md'
import { FaCheck } from 'react-icons/fa6'

export const SettingDialog = () => {
  const setting = useGetSetting()
  const updateSetting = useUpdateSetting()

  const handleSubmit: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault()
    const data = {
      openaiToken: e.currentTarget.openaiToken.value as string,
    }
    updateSetting.mutate(data)
    setTimeout(() => updateSetting.reset(), 1000)
  }

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='soft' className='absolute right-3 top-1 m-2'><MdSettings /></Button>
      </Dialog.Trigger>

      <Dialog.Content className='max-w-[500px] relative'>
        <Dialog.Title>Setting</Dialog.Title>

        {updateSetting.isSuccess && (
          <div className='absolute top-5 right-5'><FaCheck /></div>
        )}

        <form onSubmit={handleSubmit}>
          <label className='font-semibold'>
            OpenAI Token
            <TextField.Root
              type='text'
              name='openaiToken'
              className='mt-1'
              defaultValue={setting.data?.openaiToken}
              onBlur={e => e.currentTarget.form?.requestSubmit()}
            />
          </label>
        </form>

        <Dialog.Close>
          <Button variant='surface' className='mt-5'>OK</Button>
        </Dialog.Close>
      </Dialog.Content>
    </Dialog.Root>
  )
}
