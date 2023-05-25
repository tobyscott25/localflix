import { FunctionComponent, ReactElement } from 'react'
import { Outlet } from 'react-router-dom'

import { Flex } from '@chakra-ui/react'
import { NavBar } from './NavBar'
import { Footer } from './Footer'

export const AppRoot: FunctionComponent = (): ReactElement => {
	return (
		<Flex flexDir={'column'} minHeight={'100vh'}>
			<NavBar />

			<Flex
				flexDir={'column'}
				p={8}
				// border={'1px solid green'}
				flexGrow={1}
			>
				{/* An <Outlet> renders the component for the child route that is currently active. */}
				<Outlet />
			</Flex>

			<Footer />
		</Flex>
	)
}
