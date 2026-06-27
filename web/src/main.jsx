import { createRoot } from 'react-dom/client';
import { ConfigProvider, App } from 'antd';

// 国际化
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
import zhCN from 'antd/locale/zh_CN';
dayjs.locale('zh-cn');

// 样式
import 'antd/dist/antd.css';
import '@/assets/css/theme.less';

// 字体
import 'misans/lib/Normal/MiSans-Regular.min.css';

// 路由组件
import { BrowserRouter } from 'react-router-dom';
import { GenerateRoutes } from '@/router/rules';

// 颜色定义
const customColor = {
  lightBgColor: '#f7f8f9',
  blueColor: '#0052d9',
  blueBgColor: '#e6f4ff',
  borderColor: '#e5e5e5',
  filledBgColor: '#e5e5e555',
  errorColor: '#CC0033',
  errorBgColor: '#ffe6e84a'
};

// 主题定制
const themeConfig = {
  zeroRuntime: true,
  token: {
    colorPrimary: '#000000',
    fontFamily: 'MiSans, serif',
    fontSize: 13,
    borderRadius: 0,
    margin: 10,
    marginXS: 5,
    marginSM: 10,
    marginMD: 15,
    marginLG: 20,
    paddingXS: 5,
    paddingSM: 10,
    padding: 10,
    paddingMD: 15,
    paddingLG: 20,
    marginXL: 30,
    paddingXL: 30,
    marginXXL: 40,
    colorError: customColor.errorColor,
    colorErrorBg: customColor.errorBgColor,
    colorErrorBgActive: customColor.errorBgColor,
    colorErrorBgHover: customColor.errorBgColor,
    lineHeight: '20px',
    controlHeight: 28,
    colorLink: customColor.blueColor
  },
  components: {
    Layout: {
      headerBg: '#ffffff',
      bodyBg: '#ffffff',
      footerBg: '#ffffff',
      footerPadding: 0,
      headerHeight: 50,
      headerPadding: '0 10px'
    },
    Dropdown: {
      paddingBlock: 5, // 菜单内边距垂直
      controlPaddingHorizontal: 15, // 菜单内边距水平
      lineHeight: '20px', // 菜单高度
      fontSizeSM: 13, // 图标字体大小
      marginXS: 10, // 图标和文字距离
      controlItemBgActive: customColor.filledBgColor,
      controlItemBgActiveHover: customColor.filledBgColor,
      controlItemBgHover: customColor.filledBgColor
    },
    Button: {
      contentFontSize: 12,
      contentFontSizeSM: 12,
      defaultShadow: 'none',
      primaryShadow: 'none',
      dangerShadow: 'none',
      controlHeightSM: 20,
      controlHeight: 26,
      controlHeightLG: 28,
      paddingInline: 10,
      paddingInlineSM: 5,
      colorLink: customColor.blueColor,
      colorLinkHover: customColor.blueColor,
      colorLinkActive: customColor.blueColor,
      colorPrimary: customColor.blueColor,
      colorPrimaryActive: customColor.blueColor,
      colorPrimaryBg: customColor.blueBgColor,
      colorPrimaryBgHover: customColor.blueBgColor,
      colorPrimaryBorder: customColor.blueBgColor,
      colorPrimaryHover: customColor.blueColor,
      // 错误按钮样式
      colorErrorBgFilledHover: customColor.errorBgColor,
      colorErrorBgFilledHover: customColor.errorBgColor,
      colorErrorBgActive: customColor.errorBgColor,
      colorErrorActive: customColor.errorColor,
      colorErrorHover: customColor.errorColor,
      // 普通按钮样式
      colorBgSolidActive: '#000000',
      colorBgSolidHover: '#000000',
      // 普通按钮填充颜色
      colorFill: 'rgba(0,0,0,0.04)',
      colorFillSecondary: 'rgba(0,0,0,0.04)',
      colorFillTertiary: 'rgba(0,0,0,0.04)',
      blue6: customColor.blueColor
    },
    Form: {
      labelHeight: 28,
      controlHeight: 28,
      lineHeight: '28px'
    },
    Input: {
      activeShadow: 'none',
      errorActiveShadow: 'none',
      warningActiveShadow: 'none',
      activeBorderColor: customColor.borderColor,
      hoverBorderColor: customColor.borderColor,
      paddingBlock: 3.43,
      paddingBlockSM: 2.43,
      paddingInline: 10,
      paddingInlineSM: 10
    },
    InputNumber: {
      activeBorderColor: customColor.borderColor
    },
    Slider: {
      railSize: 8,
      handleSize: 8,
      handleLineWidth: 1
    },
    Select: {
      optionPadding: '3px 10px',
      optionSelectedBg: customColor.filledBgColor,
      controlItemBgActiveHover: customColor.filledBgColor,
      multipleItemHeight: 20
    },
    Cascader: {
      optionPadding: '3px 10px',
      optionSelectedBg: customColor.filledBgColor
    },
    DatePicker: {
      paddingBlock: 3.43,
      paddingInline: 10,
      activeBorderColor: customColor.borderColor,
      cellActiveWithRangeBg: customColor.borderColor
    },
    Table: {
      lineHeight: '32px',
      cellPaddingInline: 10,
      cellPaddingBlock: 0,
      rowSelectedBg: customColor.lightBgColor,
      colorLinkHover: customColor.blueColor
    },
    Pagination: {
      itemSize: 20,
      itemActiveBg: 'transparent',
      itemActiveColor: customColor.errorColor,
      itemActiveColorHover: customColor.errorColor
    },
    Message: {
      contentPadding: '5px 10px',
    }
  }
};

createRoot(document.getElementById('root')).render(
  <ConfigProvider locale={zhCN} theme={themeConfig} message={{ duration: 3 }}>
    <App>
      <BrowserRouter>
        <GenerateRoutes />
      </BrowserRouter>
    </App>
  </ConfigProvider>
);
