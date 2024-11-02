import { Link } from '@radix-ui/themes'
import { FaBolt } from 'react-icons/fa'
import { SettingDialog } from './SettingDialog'

export const Header = () => {
  return (
    <header className='p-3 relative'>
      <Link href='/' className='text-2xl font-bold text-white block'>
        <FaBolt className='text-xl align-middle mx-2 inline' />
        pinit
      </Link>
      <SettingDialog />
    </header>
  )
}
