import { Navigate, useRoutes } from 'react-router-dom';
import RouteLazyLoad from '@/router/lazyload';
import AdminLayout from '@/components/layout';
import ErrorLayout from '@/components/error';

// 路由列表
export const RouteRules = [
  {
    path: '/',
    element: <Navigate to="/dashboard" />
  },
  {
    path: '/',
    element: <AdminLayout />,
    children: [
      {
        path: '/dashboard',
        element: RouteLazyLoad(() => import('@/pages/dashboard/dashboard'))
      },
      {
        path: '/template',
        children: [
          {
            path: '/template/button',
            element: RouteLazyLoad(() => import('@/pages/template/button'))
          },
          {
            path: '/template/form',
            element: RouteLazyLoad(() => import('@/pages/template/form'))
          },
          {
            path: '/template/table',
            element: RouteLazyLoad(() => import('@/pages/template/table'))
          }
        ]
      }
    ]
  },
  {
    path: '/error',
    element: <ErrorLayout />,
    children: [
      {
        path: '/error/404',
        element: RouteLazyLoad(() => import('@/pages/error/404'))
      }
    ]
  },
  {
    path: '*',
    element: <Navigate to="/error/404" />
  }
];

// 生成 React-Router 路由
export const GenerateRoutes = () => useRoutes(RouteRules);
