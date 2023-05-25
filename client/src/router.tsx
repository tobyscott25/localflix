import { createBrowserRouter } from 'react-router-dom'
import { Router } from '@remix-run/router'

import { AppRoot } from './components/App'
import { RouterError } from './components/App/RouterError'
import { Home } from './components/App/Home'
import { NotFound } from './components/App/NotFound'
import { VideoViewer } from './components/App/VideoViewer'

const router: Router = createBrowserRouter([
	{
		path: '/',
		element: <AppRoot />,
		errorElement: <RouterError />,
		children: [
			{
				index: true,
				element: <Home />,
			},
			{
				path: 'video/:fileName',
				element: <VideoViewer />,
			},
			{
				path: '*',
				element: <NotFound />,
			},
		],
	},
])

export default router
