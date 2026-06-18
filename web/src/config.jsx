// 当前版本
const CURRENT_VERSION = 'v1.10.11';

// 后端域名地址，从 window.CONFIG 获取，如果未定义则使用空字符串避免报错
const SYSTEM_BACKEND_URL = window.CONFIG?.SYSTEM_BACKEND_URL || '';

// 系统 API 前缀
const SYSTEM_API_PREFIX = '/api/v1';

// 系统开放 API 前缀
const SYSTEM_OPEN_API_PREFIX = '/openapi/v1';

// 接口地址
const SYSTEM_BACKEND_API = {
  OPEN: {
    HEALTH: { METHOD: 'GET', URL: SYSTEM_BACKEND_URL + SYSTEM_OPEN_API_PREFIX + '/health' },
    LOGIN: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_OPEN_API_PREFIX + '/login' },
    LOGIN_DINGTALK: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_OPEN_API_PREFIX + '/login/dingtalk' },
    LOGIN_FEISHU: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_OPEN_API_PREFIX + '/login/feishu' },
    LOGIN_WECHAT: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_OPEN_API_PREFIX + '/login/wechat' }
  },
  NO_PERMISSION: {
    LOGOUT: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/logout' }
  },
  SYSTEM: {
    USER: {
      LIST: { METHOD: 'GET', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/user/list' },
      CREATE: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/user/create' }
    },
    MENU: {
      LIST: { METHOD: 'GET', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/menu/list' },
      CREATE: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/menu/create' }
    },
    ROLE: {
      LIST: { METHOD: 'GET', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/role/list' },
      CREATE: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/role/create' }
    },
    API_CATEGORY: {
      LIST: { METHOD: 'GET', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/api-category/list' },
      CREATE: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/api-category/create' }
    },
    API: {
      LIST: { METHOD: 'GET', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/api/list' },
      CREATE: { METHOD: 'POST', URL: SYSTEM_BACKEND_URL + SYSTEM_API_PREFIX + '/system/api/create' }
    }
  }
};

export { CURRENT_VERSION, SYSTEM_BACKEND_URL, SYSTEM_API_PREFIX, SYSTEM_OPEN_API_PREFIX, SYSTEM_BACKEND_API };
