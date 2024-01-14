import { style } from '@vanilla-extract/css'

export default {
  textarea: style({
    height: '300px',
  }),
  speakButton: style({
    position: 'absolute',
    right: '10px',
    bottom: '5px',
    cursor: 'pointer',
  })
}
