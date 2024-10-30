import { Box } from '@radix-ui/themes'

type Props = {
  name: string
}
export const BinderStoryName = ({ name }: Props) => {
  return (
    <Box mt='6' ml='3'>
      {name}
    </Box>
  )
}
