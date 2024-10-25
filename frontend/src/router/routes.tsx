import { RouteObject } from "react-router-dom"

import Home from "../pages/Home"
import SignUp from "../views/SignUp"
import Shell from "../components/layout/Shell"
import LandingShell from "../components/layout/LandingShell"
import Landing from "../pages/Landing"


export const commonRoutes: RouteObject[] = [
  {
    path: '/',
    element: (
      <LandingShell />
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