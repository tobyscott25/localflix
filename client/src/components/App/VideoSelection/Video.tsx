import { FunctionComponent, ReactElement } from 'react'
import { Box, Text } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'
// import { FilesEndpointReturnShape } from '../../../utils/api/helper'

interface VideoProps {
	fileName: string
}

export const Video: FunctionComponent<VideoProps> = ({
	fileName,
}): ReactElement => {
	const navigate = useNavigate()
	return (
		<Box
			width={'300px'}
			onClick={() => {
				console.log('clicked')
				navigate(`/video/${encodeURIComponent(fileName)}`)
			}}
			borderRadius={'md'}
			shadow={'md'}
			_hover={{
				shadow: 'lg',
				cursor: 'pointer',
			}}
			m={2}
			p={4}
		>
			<Box
				height={'30px'}
				bgColor={'grey'}
				borderRadius={'md'}
				mb={2}
			></Box>
			<Text fontWeight={'bold'}>{fileName}</Text>
		</Box>
	)
}
