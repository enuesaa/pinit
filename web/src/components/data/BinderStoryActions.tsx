import { Box, Button } from '@radix-ui/themes'
import styles from './BinderStoryActions.css'

export const BinderStoryActions = () => {
  return (
    <>
      <ActionItem />
      <ActionItem />
    </>
  )
}

const ActionItem = () => {
  return (
    <Button m='2' className={styles.button}>
      action
    </Button>
  )
}