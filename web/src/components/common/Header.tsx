import { Container, Flex, Link, Separator } from '@radix-ui/themes'
import { FaBolt } from 'react-icons/fa'
import styles from './Header.css'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4' pt='1' pr='5' pb='2'>
        <Link href='/' className={styles.heading}>
          <FaBolt />
          pinit
        </Link>
      </Container>
      <Separator size='4' />
    </header>
  )
}
