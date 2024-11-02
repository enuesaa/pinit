import { globalStyle, style } from '@vanilla-extract/css'

const textarea = style({
  minHeight: '80vh',
  padding: '10px',
})

globalStyle(`${textarea} textarea`, {
  outline: 'none',
})

export default {
  textarea,
}
