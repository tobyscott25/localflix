import { baseUrl } from './helper'

export interface Library {
	version: string
	videos: File[]
}

export interface File {
	id: string
	title: string
	description: string
	file_name: string
	file_size: string
	last_modified: string
	checksum_sha256: string
}

export type LibraryEndpointReturnShape = Library
export type VideoDetailsEndpointReturnShape = File

export function getLibrary(): Promise<Response> {
	return fetch(`${baseUrl}/library`)
}

export function getVideoDetails(id: string): Promise<Response> {
	return fetch(`${baseUrl}/library/videos/${id}`)
}
