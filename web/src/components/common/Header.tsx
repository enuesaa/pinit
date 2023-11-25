import Link from 'next/link'
import { FaBolt } from 'react-icons/fa'
import { Container, Flex, Separator, Box } from '@radix-ui/themes'
import { PlaygorundMenu } from './PlaygorundMenu'
import styles from './Header.css'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4'>
        <Flex width='100%' height='7'>
          <Link href='/' className={styles.heading}>
            <FaBolt />
            pinit
          </Link>
          <Box grow='1' style={{ textAlign: 'right' }}>
            <PlaygorundMenu />
          </Box>
        </Flex>
      </Container>
      <Separator size='4' />
    </header>
  )
}
