import { forwardRef } from 'react';
import { Card } from 'antd';

// 带介绍的卡片
const SCard = forwardRef(({ title, description, ...rest }, ref) => {
  const cardTitle =
    title !== undefined ? (
      <div>
        <div>{title}</div>
        {description && <div className="s-card__description">{description}</div>}
      </div>
    ) : undefined;

  return <Card className="s-card" ref={ref} title={cardTitle} {...rest} />;
});

export default SCard;
