import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { commonRoutes } from './routes';


export default function Router() {

  const router = createBrowserRouter(commonRoutes);

  return <RouterProvider router={router} />;
}

