import { Input } from 'antd';
import { useState, useEffect, useContext } from 'react';
import { FormItemInputContext } from 'antd/es/form/context';

const { TextArea } = Input;

const STextArea = (props) => {
  const [lineCount, setLineCount] = useState(1);

  const { validateStatus: propsValidateStatus, ...textAreaProps } = props;

  // 从 Form context 获取校验状态
  const { status: contextStatus } = useContext(FormItemInputContext);
  const validateStatus = propsValidateStatus || contextStatus;

  useEffect(() => {
    const value = props.value || '';
    setLineCount(value.split('\n').length);
  }, [props.value]);

  const handleChange = (e) => {
    const value = e.target.value;
    setLineCount(value.split('\n').length);

    if (textAreaProps.onChange) {
      textAreaProps.onChange(e);
    }
  };

  const containerCls = `s-textarea ${validateStatus ? `s-textarea--${validateStatus}` : ''}`;

  return (
    <div className={containerCls}>
      <div className="s-textarea__linenum">
        {Array.from({ length: lineCount }, (_, i) => (
          <div key={i + 1} className="s-textarea__linenum-item">{i + 1}</div>
        ))}
      </div>
      <div className="s-textarea__wrapper">
        <TextArea
          {...textAreaProps}
          onChange={handleChange}
          className="s-textarea__input"
        />
      </div>
    </div>
  );
};

export default STextArea;
