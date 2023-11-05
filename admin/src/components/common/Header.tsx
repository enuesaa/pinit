import Link from 'next/link'
import { AiTwotoneSetting } from 'react-icons/ai'
import { FaGuitar } from 'react-icons/fa'
import { Flex, Box, Container } from '@radix-ui/themes'
import { css } from '@emotion/react'

export const Header = () => {
  const styles = {
    main: css({
      height: '50px',
      lineHeight: '50px',
      fontSize: '25px',
      fontWeight: 'bold',
    }),
    heading: css({
      color: 'white',
      margin: '0 10px',
      textDecoration: 'none',
      'svg': {
        verticalAlign: 'middle',
        margin: '0 10px',
      },
    }),
    setting: css({
      display: 'block',
      width: '100px',
      color: 'white',
      textDecoration: 'none',
      'svg': {
        verticalAlign: 'middle',
        margin: '0 10px',
      },
    }),
  }

  return (
    <header css={styles.main}>
      <Container size='4'>
        <Flex>
          <Box grow='1'>
            <Link href='/' css={styles.heading}>
              <FaGuitar />
              my-nextjs-template
            </Link>
          </Box>

          <Link href='/setting' css={styles.setting}>
            <AiTwotoneSetting />
          </Link>
        </Flex>
      </Container>
    </header>
  )
}
