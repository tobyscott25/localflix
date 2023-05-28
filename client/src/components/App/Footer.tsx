import { FunctionComponent, ReactElement } from 'react'
import { Flex, Link, useColorMode } from '@chakra-ui/react'
import { ColorModeSwitcher } from '../ColorModeSwitcher'

export const Footer: FunctionComponent = (): ReactElement => {
	const { colorMode } = useColorMode()
	return (
		<Flex
			alignItems={'center'}
			justifyContent={'space-between'}
			borderTop={1}
			borderStyle={'solid'}
			borderColor={colorMode === 'light' ? 'gray.100' : 'gray.900'}
			shadow={'sm'}
			px={10}
			py={1}
			fontSize={'sm'}
		>
			<Flex alignItems={'center'}>
				<Link
					mr={10}
					href={'https://github.com/tobyscott25'}
					target={'_blank'}
				>
					&copy; Toby Scott 2023
				</Link>
				<Link
					mr={10}
					href={'https://github.com/tobyscott25/localflix'}
					target={'_blank'}
				>
					View this project on GitHub
				</Link>
			</Flex>

			<ColorModeSwitcher />
		</Flex>
	)
}
