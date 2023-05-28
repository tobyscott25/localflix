import { FunctionComponent, ReactElement, useEffect, useState } from 'react'
import { Box, Code, Text } from '@chakra-ui/react'
import { baseUrl } from '../../../utils/api/helper'
import { useLocation, useParams } from 'react-router-dom'
import { VideoSelection } from '../VideoSelection'
import {
	File,
	VideoDetailsEndpointReturnShape,
	getVideoDetails,
} from '../../../utils/api/files'
import { formatDate } from '../../../utils/formatDate'

export const VideoViewer: FunctionComponent = (): ReactElement => {
	const location = useLocation()
	const params = useParams()

	const [videoDetails, setVideoDetails] = useState<File>()

	useEffect(() => {
		async function fetchVideoDetails() {
			try {
				const response = await getVideoDetails(
					params.checksum?.toString() as string
				)
				const data =
					(await response.json()) as VideoDetailsEndpointReturnShape
				setVideoDetails(data)
			} catch (error) {
				console.error('Error fetching files:', error)
			}
		}

		fetchVideoDetails()
	}, [location])

	return (
		<Box>
			<Box
				as="video"
				controls
				autoPlay
				src={`${baseUrl}/assets${videoDetails?.path}`}
				// poster="thumbnail_image_url_goes_here"
				// objectFit="contain"
				// sx={{
				// 	aspectRatio: '16/9',
				// }}
				maxH={'50vh'}
			/>
			<Box my={6}>
				<Text fontSize={'xl'} fontWeight={'bold'}>
					{videoDetails?.name}
				</Text>
				<Box mt={2}>
					<Text>
						Last Modified:{' '}
						{formatDate(
							videoDetails?.lastModified as string,
							'long'
						)}
					</Text>
					<Text>Size: {videoDetails?.size}</Text>
					<Text>
						Checksum: <Code>{videoDetails?.checksum}</Code>
					</Text>
				</Box>
			</Box>

			<hr />
			<Text fontSize={'xl'}>More videos</Text>
			<VideoSelection />
		</Box>
	)
}
