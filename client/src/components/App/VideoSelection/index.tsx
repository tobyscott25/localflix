import { FunctionComponent, ReactElement, useState, useEffect } from 'react'
import { Flex } from '@chakra-ui/react'
import { Video } from './Video'
import { FilesEndpointReturnShape } from '../../../utils/api/helper'
import { getFilesList } from '../../../utils/api/files'

export const VideoSelection: FunctionComponent = (): ReactElement => {
	const [files, setFiles] = useState<string[]>()

	useEffect(() => {
		async function fetchFiles() {
			try {
				const response = await getFilesList()
				const data = (await response.json()) as FilesEndpointReturnShape
				setFiles(data.files)
			} catch (error) {
				console.error('Error fetching files:', error)
			}
		}

		fetchFiles()
	}, [])

	return (
		<Flex>
			{files &&
				files.map((fileName) => (
					<Video key={fileName} fileName={fileName} />
				))}
		</Flex>
	)
}
