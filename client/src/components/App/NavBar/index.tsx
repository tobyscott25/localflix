import { FunctionComponent, ReactElement, useState } from 'react'

import { Link as ReactRouterLink } from 'react-router-dom'
import { Button, Flex, Link, Text, useColorMode } from '@chakra-ui/react'
import {
	LibraryEndpointReturnShape,
	getLibrary,
	syncLibrary,
} from '../../../utils/api/library'
// We need to use Chakra UI's <Link> component for consistency with the rest of the UI.
// But we need to use React Router's <Link> component for the routing to work properly.
// So we import Chakra UI's <Link> component, and then import React Router's <Link> component as ReactRouterLink.
// We can then pass the "as" prop to Chakra UI's <Link> component. See: https://chakra-ui.com/docs/components/link/usage#usage-with-routing-library

export const NavBar: FunctionComponent = (): ReactElement => {
	const [librarySyncing, setLibrarySyncing] = useState<boolean>(false)

	async function handleSyncLibrary() {
		if (librarySyncing) return

		setLibrarySyncing(true)
		const sync = await syncLibrary()
		if (!sync.ok) {
			console.log('Failed to sync library')
			setLibrarySyncing(false)
			return
		}
		console.log('Successfully synced library')

		const response = await getLibrary()
		const data = (await response.json()) as LibraryEndpointReturnShape
		setLibrarySyncing(false)
		console.log(data)
	}

	const { colorMode } = useColorMode()
	return (
		<Flex
			alignItems={'center'}
			justifyContent={'space-between'}
			borderBottom={1}
			borderStyle={'solid'}
			borderColor={colorMode === 'light' ? 'gray.100' : 'gray.900'}
			shadow={'sm'}
			px={10}
			py={5}
			fontSize={'sm'}
		>
			<Flex alignItems={'center'}>
				<Text
					mr={10}
					fontSize={'xl'}
					fontWeight={'extrabold'}
					textShadow={'2px 2px #FF0000'}
				>
					LOCALFLIX
				</Text>
				<Link as={ReactRouterLink} to="/browse" mr={10}>
					Browse
				</Link>
				<Link as={ReactRouterLink} to="/search" mr={10}>
					Search
				</Link>
			</Flex>

			<Flex>
				<Button onClick={handleSyncLibrary} isDisabled={librarySyncing}>
					{librarySyncing ? 'Syncing...' : 'Sync Library'}
				</Button>
				<Link as={ReactRouterLink} to="/settings" ml={10}>
					Settings
				</Link>
			</Flex>
		</Flex>
	)
}
