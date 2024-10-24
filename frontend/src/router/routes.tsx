import { RouteObject } from "react-router-dom"

import Home from "../pages/Home"
import SignUp from "../views/SignUp"
import Shell from "../components/Shell"

export const commonRoutes: RouteObject[] = [
  {
    path: '/',
    element: (
      <Shell />
    ),
    errorElement: <div>404</div>,
    children: [
      {
        path: '',
        element: <Home />,
      },
      {
        path: 'sign-up',
        element: <SignUp />,
      }
    ]
    
  },
]