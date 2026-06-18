import { Suspense, lazy } from 'react';
import { Spin } from 'antd';

// 路由加载动画
const RouteLoading = () => (
  <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
    <Spin size="large" />
  </div>
);

// 路由懒加载包装
const RouteLazyLoad = (importFn) => {
  const Component = lazy(importFn);
  return (
    <Suspense fallback={<RouteLoading />}>
      <Component />
    </Suspense>
  );
};

export default RouteLazyLoad;
