import { Navigate, createBrowserRouter } from 'react-router-dom'
import { Router } from '@remix-run/router'

import { AppRoot } from './components/App'
import { RouterError } from './components/App/RouterError'
import { Browse } from './components/App/Browse'
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
				element: <Navigate to={'/browse'} />,
			},
			{
				path: 'browse',
				element: <Browse />,
			},
			{
				path: 'video/:id',
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
