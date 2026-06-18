import { Button, Result } from 'antd';
import { useNavigate } from 'react-router-dom';
import { ArrowLeftOutlined } from '@ant-design/icons';

const NotFound = () => {
  const navigate = useNavigate();
  return (
    <Result
      className="s-error-result"
      status="404"
      title="404"
      subTitle="抱歉，您访问的页面不存在"
      extra={
        <Button type="primary" icon={<ArrowLeftOutlined />} onClick={() => navigate('/')}>
          返回首页
        </Button>
      }
    />
  );
};

export default NotFound;
