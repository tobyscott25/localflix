import { FunctionComponent, ReactElement } from 'react'
import { Box, Text } from '@chakra-ui/react'
import { VideoSelection } from '../VideoSelection'

export const Browse: FunctionComponent = (): ReactElement => {
	return (
		<Box>
			<Text>Welcome to your localflix library!</Text>
			<VideoSelection />
		</Box>
	)
}
