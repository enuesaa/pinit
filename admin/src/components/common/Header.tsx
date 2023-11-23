import Link from 'next/link'
import { FaBolt } from 'react-icons/fa'
import { Container, Separator } from '@radix-ui/themes'
import styles from './Header.css'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4'>
        <Link href='/' className={styles.heading}>
          <FaBolt />
          pinit
        </Link>
      </Container>
      <Separator size='4' />
    </header>
  )
}
