import { Layout } from 'antd';
import { Outlet } from 'react-router-dom';
import { LogoBlack } from '@/components/image';
import { FooterText } from '@/components/text';

const { Content, Footer } = Layout;

const ErrorLayout = () => {
  return (
    <Layout>
      <div className="s-base-logo">
        <img src={LogoBlack} alt="logo" />
      </div>
      <Content>
        <Outlet />
      </Content>
      <Footer style={{ textAlign: 'center', lineHeight: '50px' }}>{FooterText}</Footer>
    </Layout>
  );
};

export default ErrorLayout;
