import { Helmet } from 'react-helmet';
import { Button, Checkbox, Divider, Flex, Form, Input, App } from 'antd';
import { LogoBlack } from '@/components/Image';
import { DingdingOutlined } from '@ant-design/icons';
import HTTP from '@/utils/axios';
import { SYSTEM_BACKEND_API } from '@/config';
import { useNavigate } from 'react-router-dom';
import { SaveToken } from '@/utils/token';

const Login = () => {
  const { message } = App.useApp();
  const navigate = useNavigate();

  // 登录方法
  const [loginForm] = Form.useForm();
  const loginHandler = async values => {
    try {
      const res = await HTTP.POST(SYSTEM_BACKEND_API.OPEN.LOGIN.URL, values);
      if (res.code === 200) {
        SaveToken(res.data.token, res.data.expire_time);
        message.success('登录成功');
        navigate('/');
      } else {
        message.error(res.msg);
      }
    } catch (e) {
      message.error(e.message || '网络异常，请稍后重试');
    }
  };

  return (
    <>
      <Helmet>
        <title>用户登录</title>
      </Helmet>
      <div className="s-login-container">
        <div style={{ textAlign: 'center', marginBottom: '20px' }}>
          <img src={LogoBlack} alt="logo" style={{ height: '25px', userSelect: 'none' }} />
        </div>
        <Divider plain>欢迎回来，立即登录</Divider>
        <div>
          <Form
            name="login"
            form={loginForm}
            initialValues={{ remember: true }}
            onFinish={loginHandler}
            variant="filled"
            layout="vertical"
          >
            <Form.Item
              label="账号"
              name="account"
              rules={[{ required: true, message: '请输入用户名 / 手机号 / 邮箱' }]}
            >
              <Input placeholder="支持用户名 / 手机号 / 邮箱登录" autoComplete="off" />
            </Form.Item>
            <Form.Item
              label="密码"
              name="password"
              rules={[{ required: true, message: '请输入密码' }]}
            >
              <Input.Password placeholder="密码" />
            </Form.Item>
            <Form.Item
              label="验证码"
              name="code"
              rules={[{ required: true, message: '请输入验证码' }]}
            >
              <Input placeholder="双因子认证验证码（未绑定则随便输入）" />
            </Form.Item>
            <Form.Item>
              <Flex justify="space-between" align="center">
                <Form.Item name="remember" valuePropName="checked" noStyle>
                  <Checkbox>记住我</Checkbox>
                </Form.Item>
                <a href="">忘记密码？</a>
              </Flex>
            </Form.Item>
            <Form.Item>
              <Button size="large" block color="default" variant="solid" htmlType="submit">
                登录
              </Button>
            </Form.Item>
            <Divider plain>切换登录方式</Divider>
            <Form.Item>
              <Button
                size="large"
                block
                htmlType="submit"
                type="default"
                icon={<DingdingOutlined />}
                variant="solid"
                color="primary"
              >
                使用钉钉扫码登录
              </Button>
            </Form.Item>
          </Form>
          <div className="login-version">版本号：1.0.0</div>
        </div>
      </div>
    </>
  );
};

export default Login;
