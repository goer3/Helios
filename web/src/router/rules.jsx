import { Navigate, useRoutes } from 'react-router-dom';
import RouteLazyLoad from '@/router/lazyload';
import AdminLayout from '@/components/layout';
import LoginAndErrorLayout from '@/components/login-and-error-layout';

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
        path: '/user',
        children: [
          {
            path: '/user/list',
            element: RouteLazyLoad(() => import('@/pages/user/list'))
          },
          {
            path: '/user/group',
            element: RouteLazyLoad(() => import('@/pages/user/group'))
          },
          {
            path: '/user/project',
            element: RouteLazyLoad(() => import('@/pages/user/project'))
          },
          {
            path: '/user/duty-roster',
            element: RouteLazyLoad(() => import('@/pages/user/duty-roster'))
          }
        ]
      },
      {
        path: '/system',
        children: [
          {
            path: '/system/menu',
            element: RouteLazyLoad(() => import('@/pages/system/menu'))
          }
        ]
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
    path: '/',
    element: <LoginAndErrorLayout />,
    children: [
      {
        path: '/login',
        element: RouteLazyLoad(() => import('@/pages/login/login'))
      },
      {
        path: '/error',
        children: [
          {
            path: '/error/403',
            element: RouteLazyLoad(() => import('@/pages/error/403'))
          },
          {
            path: '/error/404',
            element: RouteLazyLoad(() => import('@/pages/error/404'))
          },
          {
            path: '/error/500',
            element: RouteLazyLoad(() => import('@/pages/error/500'))
          }
        ]
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
