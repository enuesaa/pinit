import { Box, Container, Flex, Link, Separator } from '@radix-ui/themes'
import { FaBolt } from 'react-icons/fa'
import styles from './Header.css'
import { PlaygorundMenu } from './PlaygorundMenu'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4' pt='1' pr='5' pb='2'>
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
