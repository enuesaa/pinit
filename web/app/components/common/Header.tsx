import { Box, Container, Flex, Separator } from '@radix-ui/themes'
import Link from 'next/link'
import { FaBolt } from 'react-icons/fa'
import styles from './Header.css'
import { PlaygorundMenu } from './PlaygorundMenu'

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
