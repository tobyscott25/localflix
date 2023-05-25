import { StrictMode } from 'react'
import { Root, createRoot } from 'react-dom/client'
import { RouterProvider } from 'react-router-dom'
import { ChakraProvider, ColorModeScript } from '@chakra-ui/react'
import reportWebVitals from './reportWebVitals'
import * as serviceWorker from './serviceWorker'

import router from './router'
import theme from './theme'

const container = document.getElementById('root') as HTMLElement
if (!container) throw new Error('Failed to find the root element')
const root: Root = createRoot(container)

root.render(
	<StrictMode>
		<ChakraProvider theme={theme}>
			<ColorModeScript />
			<RouterProvider router={router} />
		</ChakraProvider>
	</StrictMode>
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://cra.link/PWA
serviceWorker.unregister()

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
