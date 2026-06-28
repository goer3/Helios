import { useState, useEffect } from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { Spin } from 'antd';
import { GetToken } from '@/utils/token';
import { SYSTEM_BACKEND_API } from '@/config';
import HTTP from '@/utils/axios';

// 路由守卫组件
// requireAuth=true (默认): 未登录 → /login，只有登录后才能访问的页面
// requireAuth=false: 已登录 → /dashboard，登录和没登录都能访问的页面，但是检测到已经登录后会跳转到 /dashboard
const AuthGuard = ({ children, requireAuth = true }) => {
  const [verified, setVerified] = useState(null); // null=验证中, true=通过, false=未登录
  const location = useLocation();

  useEffect(() => {
    const checkAuth = async () => {
      const token = GetToken();
      // 无本地 token，直接判定未登录
      if (!token) {
        setVerified(false);
        return;
      }

      // 有本地 token，请求后端验证 token 是否有效（服务端可能已主动失效）
      try {
        const res = await HTTP.GET(SYSTEM_BACKEND_API.NO_PERMISSION.TOKEN_VERIFY.URL);
        setVerified(res.code === 200);
      } catch (e) {
        console.error('Token 验证失败:', e);
        setVerified(false);
      }
    };

    checkAuth();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (verified === null) {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <Spin size="large" />
      </div>
    );
  }

  // requireAuth=true: 未登录跳转 /login
  // requireAuth=false: 已登录跳转 /dashboard
  if ((requireAuth && !verified) || (!requireAuth && verified)) {
    return <Navigate to={requireAuth ? '/login' : '/dashboard'} state={{ from: location }} replace />;
  }

  return children;
};

export default AuthGuard;
