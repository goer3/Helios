import { Button, Result } from 'antd';
import { useNavigate } from 'react-router-dom';
import { ArrowLeftOutlined } from '@ant-design/icons';

const InternalServerError = () => {
  const navigate = useNavigate();
  return (
    <Result
      className="s-error-result"
      status="500"
      title="500"
      subTitle="抱歉，服务器内部错误，请联系管理员"
      extra={
        <Button type="primary" icon={<ArrowLeftOutlined />} onClick={() => navigate('/')}>
          返回首页
        </Button>
      }
    />
  );
};

export default InternalServerError;
