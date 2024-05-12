import { WithSidebar } from '@/components/common/WithSidebar'
import { BinderStory } from '@/components/data/BinderStory'
import { Greet2 } from '../../wailsjs/go/main/App.js'
import { MouseEventHandler } from 'react'

export default function Page() {

  const handleClick: MouseEventHandler<HTMLButtonElement> = async (e) => {
    e.preventDefault()
    const res = await Greet2('a')
    console.log(res)
  }

  return (
    <WithSidebar>
      <BinderStory />
      <button onClick={handleClick}>call</button>
    </WithSidebar>
  )
}
