import { globalStyle, style } from '@vanilla-extract/css'

const textarea = style({
  height: '300px',
  padding: '10px',
})

globalStyle(`${textarea} textarea`, {
  outline: 'none',
})

export default {
  textarea,
}
