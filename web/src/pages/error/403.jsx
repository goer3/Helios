import { Button, Result } from 'antd';
import { useNavigate } from 'react-router-dom';
import { ArrowLeftOutlined } from '@ant-design/icons';

const Forbidden = () => {
  const navigate = useNavigate();
  return (
    <Result
      className="s-error-result"
      status="403"
      title="403"
      subTitle="抱歉，您无权限访问该资源"
      extra={
        <Button type="primary" icon={<ArrowLeftOutlined />} onClick={() => navigate('/')}>
          返回首页
        </Button>
      }
    />
  );
};

export default Forbidden;
