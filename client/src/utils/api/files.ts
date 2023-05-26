import { baseUrl } from './helper'

export interface File {
	name: string
	size: string
	path: string
	lastModified: string
}

export type FilesEndpointReturnShape = {
	files: File[]
}

export function getFilesList(): Promise<Response> {
	return fetch(`${baseUrl}/files`)
}
