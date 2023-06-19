import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from 'react-router-dom'
import './styles/index.css'

import Root from './routes/root'
import ErrorPage from './routes/error-page'
import PlayComputer from './routes/play-computer'
import PlayOnline from './routes/play-online'
import Home from './routes/home-page'
import Settings from './routes/settings-page'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />, 
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/home",
        element: <Home />,
      },
      {
        path: "/play/computer",
        element: <PlayComputer />,
      },
      {
        path: "/play/online",
        element: <PlayOnline />,
      },
      {
        path: "/settings",
        element: <Settings />,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <RouterProvider router={router} />
)
