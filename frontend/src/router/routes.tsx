import { RouteObject } from "react-router-dom"

import SignUp from "../views/SignUp"
import Shell from "../components/layout/Shell"
import Landing from "../pages/Landing"
import TierList from "../components/tierlist/TierList"
import ErrorBoundary from "../components/utils/ErrorBoundary"


export const commonRoutes: RouteObject[] = [
  {
    path: '/',
    element: (
      <Shell />
    ),
    errorElement: <ErrorBoundary />,
    children: [
      {
        path: '',
        element: <Landing />,
      },
      {
        path: '/sign-up',
        element: <SignUp />,
      },
      {
        path: '/new',
        element: <TierList />,
      }
    ]
  },
]