import { style } from '@vanilla-extract/css'

export default {
  button: style({
    selectors: {
      'button.&': {
        cursor: 'pointer',
      }
    },
  }),
}