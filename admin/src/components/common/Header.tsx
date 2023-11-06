import Link from 'next/link'
import { FaBolt } from 'react-icons/fa'
import { Container, Separator } from '@radix-ui/themes'
import { css } from '@emotion/react'

export const Header = () => {
  const styles = {
    main: css({
      height: '50px',
      lineHeight: '50px',
      fontSize: '25px',
      fontWeight: 'bold',
      margin: '0 0 10px 0',
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
  }

  return (
    <header css={styles.main}>
      <Container size='4'>
        <Link href='/' css={styles.heading}>
          <FaBolt />
          pinit
        </Link>
      </Container>
      <Separator size='4' />
    </header>
  )
}
