import { FunctionComponent, ReactElement } from 'react'
import { Box, Text } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'
import { File } from '../../../utils/api/files'
import { formatDate } from '../../../utils/formatDate'

interface VideoProps {
	file: File
}

export const Video: FunctionComponent<VideoProps> = ({
	file,
}): ReactElement => {
	const navigate = useNavigate()

	return (
		<Box
			width={'300px'}
			onClick={() => {
				console.log('clicked')
				navigate(`/video/${encodeURIComponent(file.checksum)}`)
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
			<Text fontWeight={'bold'}>{file.name}</Text>
			<Text fontSize={'sm'}>Size: {file.size}</Text>
			<Text fontSize={'sm'}>
				Last modified: {formatDate(file.lastModified, 'short')}
			</Text>
		</Box>
	)
}
