import { FunctionComponent, ReactElement } from 'react'
import { Box, Text } from '@chakra-ui/react'
import { baseUrl } from '../../../utils/api/helper'
import { useParams } from 'react-router-dom'
import { VideoSelection } from '../VideoSelection'

export const VideoViewer: FunctionComponent = (): ReactElement => {
	const params = useParams()
	return (
		<Box>
			<Text fontSize={'xl'}>{params.fileName}</Text>
			<Box
				as="video"
				controls
				src={`${baseUrl}${params.fileName}`}
				// poster="thumbnail_image_url_goes_here"
				// objectFit="contain"
				// sx={{
				// 	aspectRatio: '16/9',
				// }}
				maxH={'50vh'}
			/>
			<Text fontSize={'xl'}>More videos</Text>
			<VideoSelection />
		</Box>
	)
}
