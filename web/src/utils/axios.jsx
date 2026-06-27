import axios from 'axios';
import { GetToken } from '@/utils/token';

// 默认 Axios 配置
const defaultConfig = {
  timeout: 10000,
  headers: { 'X-HELIOS-NAME': 'Helios' }
};

// HTTP 状态码错误消息映射
const HTTP_ERROR_MESSAGES = {
  400: '请求错误',
  401: '未授权，请重新登录',
  403: '无权限访问该资源',
  404: '未找到相关资源',
  408: '请求超时',
  500: '服务器内部错误',
  502: '网关错误',
  503: '服务暂不可用',
  504: '网关超时'
};

/**
 * HTTP 客户端，基于 Axios 封装
 * 支持拦截器配置、自动处理响应数据、统一错误处理
 */
class HttpClient {
  // 构造函数，初始化 Axios 实例和取消控制器
  constructor(customConfig) {
    this.instance = axios.create({ ...defaultConfig, ...customConfig });
    this.initInterceptors();
  }

  // 初始化拦截器
  initInterceptors() {
    // 请求拦截器
    this.instance.interceptors.request.use(
      config => {
        const token = GetToken();
        if (token) {
          // 往 Header 中添加 Authorization Token
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      error => Promise.reject(error)
    );

    // 响应拦截器
    this.instance.interceptors.response.use(
      response => response.data,
      error => {
        // 请求被取消
        if (axios.isCancel(error)) {
          return Promise.reject(new Error('请求已被取消'));
        }

        // 网络错误或超时
        if (!error.response) {
          if (error.code === 'ECONNABORTED' || error.message?.includes('timeout')) {
            return Promise.reject(new Error('请求超时，请稍后重试'));
          }
          return Promise.reject(new Error('网络错误，请检查网络连接'));
        }

        // HTTP 状态码错误
        const status = error.response.status;
        const message = HTTP_ERROR_MESSAGES[status] || `请求失败，状态码：${status}`;
        return Promise.reject(new Error(message));
      }
    );
  }

  // 生成请求唯一标识，用于日志追踪
  generateRequestId() {
    // 时间戳（13位）+ 7位随机数 = 20位
    const rand = String(Math.floor(Math.random() * 10000000)).padStart(7, '0');
    return `REQ-${Date.now()}${rand}`;
  }

  // 通用请求方法
  async request(method, url, config = {}) {
    const requestId = this.generateRequestId();
    const requestConfig = {
      ...config,
      method,
      url,
      headers: {
        ...config.headers,
        'X-HELIOS-REQUEST-ID': requestId
      }
    };
    return this.instance.request(requestConfig);
  }

  // GET 请求
  GET(url, config) {
    return this.request('get', url, config);
  }

  // POST 请求
  POST(url, data, config) {
    return this.request('post', url, { ...config, data });
  }

  // PUT 请求
  PUT(url, data, config) {
    return this.request('put', url, { ...config, data });
  }

  // DELETE 请求
  DELETE(url, config) {
    return this.request('delete', url, config);
  }

  // PATCH 请求
  PATCH(url, data, config) {
    return this.request('patch', url, { ...config, data });
  }
}

// 导出默认实例
const HTTP = new HttpClient();

export { HttpClient };
export default HTTP;
