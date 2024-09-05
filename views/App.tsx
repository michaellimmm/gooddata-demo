import JourneyDashboard from './features/dashboard/Dashboard';
import Login from './features/login/Login';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Register from './features/register/Register';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Login />,
  },
  {
    path: '/dashboard',
    element: <JourneyDashboard />,
  },
  {
    path: '/register',
    element: <Register />,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
