import { globalStyle, style } from '@vanilla-extract/css'

const textarea = style({
  height: '300px',
  padding: '10px',
})

globalStyle(`${textarea} textarea`, {
  outline: 'none',
})

export default {
  main: style({
    position: 'relative',
  }),
  textarea,
  speakButton: style({
    position: 'absolute',
    right: '10px',
    top: '10px',
    cursor: 'pointer',
    color: 'var(--gray-10)',
    ':hover': {
      color: 'var(--accent-a11)',
    },
  }),
  trashButton: style({
    position: 'absolute',
    right: '10px',
    top: '-50px',
    cursor: 'pointer',
    color: 'var(--gray-10)',
    ':hover': {
      color: 'var(--accent-a11)',
    },
  }),
}
