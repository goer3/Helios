import { createRoot } from 'react-dom/client';
import { ConfigProvider, App } from 'antd';

// 国际化
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
import zhCN from 'antd/locale/zh_CN';
dayjs.locale('zh-cn');

// 样式
import 'antd/dist/antd.css';
import '@/assets/css/light.less';

// 默认主题
let customTheme = localStorage.getItem('theme') || 'light';
// 根据不同主题引入不同的样式
if (customTheme === 'dark') {
  import('@/assets/css/dark.less');
}

// 字体
import 'misans/lib/Normal/MiSans-Regular.min.css';

// 路由组件
import { BrowserRouter } from 'react-router-dom';
import { GenerateRoutes } from '@/router/rules';

// 颜色
const customColorBlue = '#0052d9';
const customColorLight = '#ffffff';
const customColorDark = '#000000';
const customColorLightGray = '#f4f4f4';
const customColorFillGray = '#e5e5e5';
const customColorGray = '#f7f8f9';

// 通用样式
const customControlHeight = 28;
const customTableLineHeight = 32;
const customPaddingInline = 10;
const customMainBg = '#ffffff';

// light 样式
let customSiderBg = customColorGray;
let customSiderColor = customColorDark;
let customMenuItemSelectedColor = customColorDark;
let customMenuItemSelectedBg = customColorFillGray;

// dark 样式
if (customTheme === 'dark') {
  customSiderBg = customColorDark;
  customSiderColor = customColorLight;
  customMenuItemSelectedColor = '#ff0000';
  customMenuItemSelectedBg = customSiderBg;
}

// 主题定制
const themeConfig = {
  zeroRuntime: true,
  token: {
    colorPrimary: customColorDark,
    colorPrimaryHover: customColorDark,
    colorLinkHover: customColorDark,
    colorLink: customColorBlue,
    controlItemBgActive: customColorLightGray,
    controlItemBgHover: customColorLightGray,
    lineHeight: '20px',
    fontFamily: 'MiSans, serif',
    fontSize: 13,
    controlInteractiveSize: 13,
    borderRadius: 0
  },
  components: {
    Layout: {
      headerBg: customMainBg,
      headerPadding: 0,
      bodyBg: customMainBg,
      footerBg: customMainBg,
      footerPadding: 0,
      headerHeight: 50,
      siderBg: customSiderBg
    },
    Menu: {
      iconSize: 13,
      collapsedIconSize: 16,
      itemColor: customSiderColor,
      itemHeight: 30,
      itemBg: customSiderBg,
      subMenuItemBg: customSiderBg,
      subMenuItemSelectedColor: customSiderColor,
      itemMarginInline: 20,
      itemPaddingInline: 0,
      itemMarginBlock: 4,
      itemActiveBg: customSiderBg,
      itemSelectedBg: customMenuItemSelectedBg,
      itemSelectedColor: customMenuItemSelectedColor,
      itemHoverColor: customSiderColor,
      dropdownWidth: 'auto'
    },
    Button: {
      contentFontSize: 12,
      contentFontSizeSM: 12,
      defaultShadow: 'none',
      primaryShadow: 'none',
      dangerShadow: 'none',
      controlHeight: customControlHeight,
      controlHeightSM: 24,
      paddingInline: customPaddingInline,
      paddingInlineSM: customPaddingInline,
      blue6: customColorBlue
    },
    Form: {
      labelHeight: customControlHeight
    },
    Input: {
      activeShadow: 'none',
      errorActiveShadow: 'none',
      warningActiveShadow: 'none',
      controlHeight: customControlHeight,
      activeBorderColor: customColorFillGray,
      paddingBlockSM: 4
    },
    InputNumber: {
      controlHeight: customControlHeight,
      activeBorderColor: customColorFillGray
    },
    Slider: {
      railSize: 8,
      handleSize: 8,
      handleLineWidth: 1
    },
    Select: {
      optionHeight: customControlHeight,
      optionLineHeight: customControlHeight + 'px',
      optionPadding: '0 10px',
      optionSelectedBg: customColorFillGray
    },
    Table: {
      lineHeight: customTableLineHeight + 'px',
      cellPaddingInline: customPaddingInline,
      cellPaddingBlock: 0
    },
    Pagination: {
      itemSize: 20,
      itemActiveBg: 'transparent',
      itemActiveColor: 'red',
      itemActiveColorHover: 'red'
    },
    Dropdown: {
      paddingBlock: 5
    }
  }
};

createRoot(document.getElementById('root')).render(
  <ConfigProvider locale={zhCN} theme={themeConfig}>
    <App>
      <BrowserRouter>
        <GenerateRoutes />
      </BrowserRouter>
    </App>
  </ConfigProvider>
);
