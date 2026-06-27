import { Layout, Space } from 'antd';
import { LoginBgImage, LogoBlack } from '@/components/Image';
import { GithubOutlined } from '@ant-design/icons';
import { Outlet } from 'react-router';

const { Header, Content, Footer } = Layout;

const LoginAndErrorLayout = () => {
  return (
    <>
      <Layout className="s-login" style={{ backgroundImage: `url(${LoginBgImage})` }}>
        <Header className="s-login-header">
          <img className="s-login-logo" src={LogoBlack} alt="logo" />
          <Space size={15}>
            <a href="https://www.baidu.com" target="_blank" className="s-login-github">
              <GithubOutlined />
            </a>
          </Space>
        </Header>
        <Content className="s-login-content">
          <Outlet />
        </Content>
        <Footer className="s-login-footer">
          <Space style={{ marginBottom: '15px' }}>
            <a href="">关于我们</a>
            <a href="">联系我们</a>
            <a href="">隐私政策</a>
            <a href="">服务条款</a>
            <a href="">版权声明</a>
          </Space>
          <div>© 2026 Helios，Created by DK</div>
        </Footer>
      </Layout>
    </>
  );
};

export default LoginAndErrorLayout;
