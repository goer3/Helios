// 获取本地存储中的 token
const GetToken = () => {
  const token = localStorage.getItem('token');
  const expireTimeStr = localStorage.getItem('expire-time');

  if (!expireTimeStr || !token) {
    localStorage.clear();
    return null;
  }

  // expire-time 格式为 YYYY-MM-DD HH:MM:SS
  if (new Date(expireTimeStr) < new Date()) {
    localStorage.clear();
    return null;
  }

  return token;
};

// 保存 token 到本地存储
const SaveToken = (token, expireTime) => {
  localStorage.setItem('token', token);
  localStorage.setItem('expire-time', expireTime);
};

export { GetToken, SaveToken };
