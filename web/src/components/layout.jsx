import React, { useState, useEffect } from 'react';
import { Outlet, useNavigate, useLocation } from 'react-router-dom';
import {
  FundOutlined,
  DesktopOutlined,
  AlertOutlined,
  BorderlessTableOutlined,
  AudioMutedOutlined,
  ExceptionOutlined,
  FieldTimeOutlined,
  MailOutlined,
  AppstoreAddOutlined,
  SnippetsOutlined,
  HddOutlined,
  UsergroupAddOutlined,
  CalendarOutlined,
  ClusterOutlined,
  CommentOutlined,
  SolutionOutlined,
  SlidersOutlined,
  IdcardOutlined,
  SisternodeOutlined,
  ApiOutlined,
  AuditOutlined,
  SettingOutlined,
  CreditCardOutlined,
  QuestionCircleOutlined,
  GithubOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  MoreOutlined,
  WarningTwoTone,
  CheckOutlined,
  LogoutOutlined,
  HeartTwoTone,
  BgColorsOutlined
} from '@ant-design/icons';
import { Layout, Menu, Button, Avatar, Dropdown, Badge } from 'antd';
const { Header, Content, Footer, Sider } = Layout;
import { LogoWhite, LogoBlack, LogoCollapsed, DefaultAvatar } from '@/components/image';
import { SiderDivider } from '@/components/common';
import { GenerateGenderBadge } from '@/components/badge';
import { CURRENT_VERSION } from '@/config';

// 修改主题
let customTheme = localStorage.getItem('theme') || 'light';
const switchThemeHandler = () => {
  customTheme = customTheme === 'light' ? 'dark' : 'light';
  localStorage.setItem('theme', customTheme);
  location.reload();
};

const item1 = [
  {
    key: '/dashboard',
    label: '工作空间',
    icon: <DesktopOutlined />
  },
  {
    key: '/query',
    label: '数据查询',
    icon: <FundOutlined />
  },
  {
    key: '/datasource',
    label: '数据来源',
    icon: <HddOutlined />
  },
  {
    key: '/alarm',
    label: '监控告警',
    icon: <AlertOutlined />,
    popupOffset: [20, 0],
    children: [
      { key: '/alarm/events', label: '告警事件', icon: <ExceptionOutlined /> },
      { key: '/alarm/rules', label: '告警策略', icon: <BorderlessTableOutlined /> },
      { key: '/alarm/shielding', label: '屏蔽策略', icon: <AudioMutedOutlined /> },
      {
        key: '/alarm/history',
        label: '告警历史',
        icon: <FieldTimeOutlined />,
        children: [
          { key: '/alarm/history/todo', label: '等待处理', icon: <WarningTwoTone twoToneColor="#eb2f96" /> },
          { key: '/alarm/history/finish', label: '完成处理', icon: <CheckOutlined /> }
        ]
      }
    ]
  },
  {
    key: '/message',
    label: '消息通知',
    icon: <MailOutlined />,
    popupOffset: [20, 0],
    children: [
      { key: '/message/type', label: '告警媒介', icon: <AppstoreAddOutlined /> },
      { key: '/message/template', label: '通知模板', icon: <SnippetsOutlined /> }
    ]
  },
  {
    key: '/user',
    label: '人员组织',
    icon: <UsergroupAddOutlined />,
    popupOffset: [20, 0],
    children: [
      { key: '/user/list', label: '用户列表', icon: <SolutionOutlined /> },
      { key: '/user/groups', label: '用户分组', icon: <CommentOutlined /> },
      { key: '/user/projects', label: '项目团队', icon: <ClusterOutlined /> },
      { key: '/user/scheduling', label: '人员排班', icon: <CalendarOutlined /> }
    ]
  },
  {
    key: '/system',
    label: '系统设置',
    icon: <SlidersOutlined />,
    popupOffset: [20, 0],
    children: [
      { key: '/system/role', label: '角色授权', icon: <IdcardOutlined /> },
      { key: '/system/menu', label: '菜单配置', icon: <SisternodeOutlined /> },
      { key: '/system/api', label: '接口配置', icon: <ApiOutlined /> },
      { key: '/system/setting', label: '通用设置', icon: <SettingOutlined /> }
    ]
  },
  {
    key: '/log',
    label: '日志审计',
    icon: <AuditOutlined />
  }
];

const item2 = [
  {
    key: '/profile',
    label: '个人中心',
    icon: <CreditCardOutlined />
  },
  {
    key: '/template',
    label: '开发模板',
    icon: <HeartTwoTone />,
    popupOffset: [20, 0],
    children: [
      { key: '/template/button', label: '按钮' },
      { key: '/template/form', label: '表单' },
      { key: '/template/table', label: '表格' }
    ]
  },
  {
    key: '/help',
    label: '获取帮助',
    icon: <QuestionCircleOutlined />
  },
  {
    key: '/update',
    label: '版本更新',
    icon: <GithubOutlined />
  }
];

const dropdownItems = [
  {
    key: '1',
    label: (
      <div className="s-dropdown-user">
        <Badge size="small" className="s-badge-gender" count={GenerateGenderBadge(1)} offset={[-5, 23]}>
          <Avatar className="s-dropdown-user__avatar" size={30} src={DefaultAvatar} />
        </Badge>
        <div className="s-dropdown-user__info">
          <div className="s-dropdown-user__name">吴彦祖</div>
          <div className="s-dropdown-user__desc">互联网不知名换皮工程师</div>
        </div>
      </div>
    )
  },
  {
    type: 'divider'
  },
  {
    key: 'profile',
    label: '个人中心',
    icon: <CreditCardOutlined />
  },
  {
    key: 'switchTheme',
    label: '切换主题',
    icon: <BgColorsOutlined />,
    onClick: switchThemeHandler
  },
  {
    type: 'divider'
  },
  {
    key: 'logout',
    label: '注销登录',
    danger: true,
    icon: <LogoutOutlined />
  }
];

const AdminLayout = () => {
  // 菜单展开折叠
  const [collapsed, setCollapsed] = useState(false);
  // 路由跳转
  const navigate = useNavigate();
  const { pathname } = useLocation();

  // 版本提示
  const [showVersionTip, setShowVersionTip] = useState(false);

  useEffect(() => {
    const savedVersion = localStorage.getItem('version');
    if (savedVersion !== CURRENT_VERSION) {
      setShowVersionTip(true);
    }
  }, []);

  const dismissVersionTip = () => {
    localStorage.setItem('version', CURRENT_VERSION);
    setShowVersionTip(false);
  };

  // 点击菜单跳转
  const clickMenuItemHandler = React.useCallback(({ key }) => navigate(key), [navigate]);

  return (
    <>
      <Layout hasSider>
        <Sider
          className="s-sider"
          width={220}
          collapsedWidth={60}
          breakpoint="lg"
          onBreakpoint={broken => {
            console.log(broken);
          }}
          onCollapse={(collapsed, type) => {
            console.log(collapsed, type);
          }}
          trigger={null}
          collapsible
          collapsed={collapsed}
        >
          {/* logo */}
          <div className="s-logo">
            <img src={collapsed ? LogoCollapsed : customTheme === 'light' ? LogoBlack : LogoWhite} alt="logo" />
          </div>
          {/* 菜单 */}
          <div className="s-body">
            {collapsed ? null : <div className="s-group">平台功能</div>}
            <Menu
              theme="light"
              mode="inline"
              inlineIndent={15}
              defaultSelectedKeys={['alarm']}
              items={item1}
              selectedKeys={[pathname]}
              onClick={clickMenuItemHandler}
            />
            {collapsed ? <SiderDivider /> : <div className="s-group">更多功能</div>}
            <Menu
              theme="light"
              mode="inline"
              inlineIndent={15}
              defaultSelectedKeys={['4']}
              items={item2}
              selectedKeys={[pathname]}
              onClick={clickMenuItemHandler}
            />
          </div>
          {/* 用户 */}
          <div className="s-trigger">
            {collapsed ? (
              <Dropdown className="s-trigger__menu" menu={{ items: dropdownItems }} placement="rightBottom">
                <Badge size="small" className="s-badge-gender" count={GenerateGenderBadge(1)} offset={[-5, 23]}>
                  <Avatar className="s-trigger__avatar" size={30} src={DefaultAvatar} style={{ cursor: 'pointer' }} />
                </Badge>
              </Dropdown>
            ) : (
              <>
                <Badge size="small" className="s-badge-gender" count={GenerateGenderBadge(1)} offset={[-5, 23]}>
                  <Avatar className="s-trigger__avatar" size={30} src={DefaultAvatar} />
                </Badge>
                <div className="s-trigger__info">
                  <div className="s-trigger__name">吴彦祖</div>
                  <div className="s-trigger__desc">互联网换皮工程师</div>
                </div>
                <Dropdown className="s-trigger__menu" menu={{ items: dropdownItems }} placement="rightBottom">
                  <MoreOutlined style={{ cursor: 'pointer' }} />
                </Dropdown>
              </>
            )}
          </div>
        </Sider>
        <Button
          className="s-collapsed"
          icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
          onClick={() => setCollapsed(!collapsed)}
        ></Button>
        <Layout>
          <Header></Header>
          <Content>
            <Outlet />
          </Content>
          <Footer style={{ textAlign: 'center', lineHeight: '50px' }}>Ant Design ©2026 Created by Ant UED</Footer>
        </Layout>
      </Layout>
    </>
  );
};

export default AdminLayout;
