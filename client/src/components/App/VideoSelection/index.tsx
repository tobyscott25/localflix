import { FunctionComponent, ReactElement, useState, useEffect } from 'react'
import { Flex } from '@chakra-ui/react'
import { Video } from './Video'
import { File, LibraryEndpointReturnShape } from '../../../utils/api/library'
import { getLibrary } from '../../../utils/api/library'

export const VideoSelection: FunctionComponent = (): ReactElement => {
	const [videos, setVideos] = useState<File[]>()

	useEffect(() => {
		async function fetchFiles() {
			try {
				const response = await getLibrary()
				const data =
					(await response.json()) as LibraryEndpointReturnShape
				setVideos(data.videos)
			} catch (error) {
				console.error('Error fetching videos:', error)
			}
		}

		fetchFiles()
	}, [])

	return (
		<Flex wrap={'wrap'} gap={5}>
			{videos &&
				videos.map((video) => <Video key={video.id} video={video} />)}
		</Flex>
	)
}
