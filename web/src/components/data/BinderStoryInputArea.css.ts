import { style } from '@vanilla-extract/css'

export default {
  main: style({
    position: 'relative',
  }),
  textarea: style({
    height: '300px',
  }),
  speakButton: style({
    position: 'absolute',
    right: '10px',
    bottom: '50px',
    cursor: 'pointer',
  }),
  trashButton: style({
    position: 'absolute',
    right: '10px',
    top: '5px',
    cursor: 'pointer',
  }),
}
