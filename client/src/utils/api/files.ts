import { baseUrl } from './helper'

export function getFilesList(): Promise<Response> {
	return fetch(`${baseUrl}/files`)
}
