import { RouteObject } from "react-router-dom"

import SignUp from "../views/SignUp"
import Shell from "../components/layout/Shell"
import Landing from "../pages/Landing"


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
        element: <Landing />,
      },
      {
        path: 'sign-up',
        element: <SignUp />,
      }
    ]
    
  },
]