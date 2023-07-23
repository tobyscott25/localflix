import { FunctionComponent, ReactElement } from 'react'
import { Box, Text } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'
import { File } from '../../../utils/api/library'
import { formatDate } from '../../../utils/formatDate'

interface VideoProps {
	video: File
}

export const Video: FunctionComponent<VideoProps> = ({
	video,
}): ReactElement => {
	const navigate = useNavigate()

	return (
		<Box
			width={'300px'}
			onClick={() => {
				console.log('clicked')
				navigate(`/video/${encodeURIComponent(video.id)}`)
			}}
			borderRadius={'md'}
			// shadow={'md'}
			_hover={{
				shadow: 'lg',
				cursor: 'pointer',
			}}
			// m={2}
			p={4}
		>
			<Box
				height={'150px'}
				bgColor={'grey'}
				borderRadius={'md'}
				mb={2}
			></Box>
			<Text fontWeight={'bold'}>{video.title}</Text>
			<Text fontSize={'sm'}>Size: {video.file_size}</Text>
			<Text fontSize={'sm'}>
				Last modified: {formatDate(video.last_modified, 'short')}
			</Text>
		</Box>
	)
}
