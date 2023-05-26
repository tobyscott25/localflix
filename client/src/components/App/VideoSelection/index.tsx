import { FunctionComponent, ReactElement, useState, useEffect } from 'react'
import { Flex } from '@chakra-ui/react'
import { Video } from './Video'
import { File, FilesEndpointReturnShape } from '../../../utils/api/files'
import { getFilesList } from '../../../utils/api/files'

export const VideoSelection: FunctionComponent = (): ReactElement => {
	const [files, setFiles] = useState<File[]>()

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
		<Flex wrap={'wrap'} gap={5}>
			{files &&
				files.map((file) => <Video key={file.name} file={file} />)}
		</Flex>
	)
}
