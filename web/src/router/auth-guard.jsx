import { useState, useEffect } from 'react';
import { Navigate, useLocation } from 'react-router-dom';
import { Spin } from 'antd';
import { GetToken } from '@/utils/token';
import { SYSTEM_BACKEND_API } from '@/config';
import HTTP from '@/utils/axios';

// 路由守卫组件，用于验证用户是否登录
const AuthGuard = ({ children }) => {
  const [verified, setVerified] = useState(null); // null=验证中, true=通过, false=失败
  const location = useLocation();

  useEffect(() => {
    const checkAuth = async () => {
      const token = GetToken();
      if (!token) {
        setVerified(false);
        return;
      }

      try {
        const res = await HTTP.GET(SYSTEM_BACKEND_API.NO_PERMISSION.TOKEN_VERIFY.URL);
        if (res.code === 200) {
          setVerified(true);
        } else {
          setVerified(false);
        }
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

  if (!verified) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children;
};

export default AuthGuard;
