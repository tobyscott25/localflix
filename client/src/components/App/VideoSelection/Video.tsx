import { FunctionComponent, ReactElement } from 'react'
import { Box, Text } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'
import { File } from '../../../utils/api/files'

interface VideoProps {
	file: File
}

export const Video: FunctionComponent<VideoProps> = ({
	file,
}): ReactElement => {
	const navigate = useNavigate()

	function formatDate(dateString: string) {
		const date = new Date(dateString)

		console.log('dateString', dateString)
		console.log('date', date)

		// Options for formatting the date and time
		const options: Intl.DateTimeFormatOptions = {
			weekday: 'long',
			year: 'numeric',
			month: 'long',
			day: 'numeric',
			hour: 'numeric',
			minute: 'numeric',
			second: 'numeric',
			timeZoneName: 'short',
		}

		// Convert the date to a human-readable format
		const formattedDate = date.toLocaleString('en-US', options)

		console.log('formattedDate', formattedDate)

		return formattedDate
	}

	return (
		<Box
			width={'300px'}
			onClick={() => {
				console.log('clicked')
				navigate(`/video/${encodeURIComponent(file.name)}`)
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
				Last modified: {formatDate(file.lastModified)}
			</Text>
		</Box>
	)
}
