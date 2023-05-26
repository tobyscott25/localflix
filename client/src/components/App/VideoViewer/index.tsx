import { FunctionComponent, ReactElement } from 'react'
import { Box, Text } from '@chakra-ui/react'
import { baseUrl } from '../../../utils/api/helper'
import { useParams } from 'react-router-dom'
import { VideoSelection } from '../VideoSelection'

export const VideoViewer: FunctionComponent = (): ReactElement => {
	const params = useParams()
	return (
		<Box>
			<Box
				as="video"
				controls
				src={`${baseUrl}/assets/${params.fileName}`}
				// poster="thumbnail_image_url_goes_here"
				// objectFit="contain"
				// sx={{
				// 	aspectRatio: '16/9',
				// }}
				maxH={'50vh'}
			/>
			<Box my={6}>
				<Text fontSize={'xl'} fontWeight={'bold'}>
					{params.fileName}
				</Text>
				<Box mt={2}>
					<Text>Last Modified: </Text>
					<Text>Size: </Text>
				</Box>
			</Box>

			<hr />
			<Text fontSize={'xl'}>More videos</Text>
			<VideoSelection />
		</Box>
	)
}
