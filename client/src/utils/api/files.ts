import { baseUrl } from './helper'

export interface File {
	name: string
	size: string
	path: string
	lastModified: string
	checksum: string
}

export type FilesEndpointReturnShape = {
	files: File[]
}

export function getFilesList(): Promise<Response> {
	return fetch(`${baseUrl}/files`)
}

export type VideoDetailsEndpointReturnShape = File

export function getVideoDetails(checksum: string): Promise<Response> {
	return fetch(`${baseUrl}/files/checksum/${checksum}`)
}
