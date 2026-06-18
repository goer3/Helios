import { QuestionOutlined, WomanOutlined, ManOutlined } from '@ant-design/icons';

// 生成性别徽章
const GenerateGenderBadge = (gender) => {
  const icons = {
    1: <ManOutlined style={{ backgroundColor: '#165dff' }} />,
    2: <WomanOutlined style={{ backgroundColor: '#ff4d4f' }} />,
    default: <QuestionOutlined style={{ backgroundColor: '#999999' }} />
  };
  return icons[gender] || icons.default;
};

export {GenerateGenderBadge};
