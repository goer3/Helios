import { useState } from 'react';
import {
  AutoComplete,
  Button,
  Cascader,
  Checkbox,
  Col,
  Form,
  Input,
  Slider,
  InputNumber,
  DatePicker,
  Row,
  Select,
  Space
} from 'antd';
import { UserOutlined } from '@ant-design/icons';
import STextArea from '@/components/ui/s-textarea';
import SCard from '@/components/ui/s-card';
const { RangePicker } = DatePicker;

const residences = [
  {
    value: 'zhejiang',
    label: 'Zhejiang',
    children: [
      {
        value: 'hangzhou',
        label: 'Hangzhou',
        children: [
          {
            value: 'xihu',
            label: 'West Lake'
          }
        ]
      }
    ]
  },
  {
    value: 'jiangsu',
    label: 'Jiangsu',
    children: [
      {
        value: 'nanjing',
        label: 'Nanjing',
        children: [
          {
            value: 'zhonghuamen',
            label: 'Zhong Hua Men'
          }
        ]
      }
    ]
  }
];
const formItemLayout = {
  labelCol: {
    xs: { span: 24 },
    sm: { span: 8 }
  },
  wrapperCol: {
    xs: { span: 24 },
    sm: { span: 16 }
  }
};
const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 24,
      offset: 0
    },
    sm: {
      span: 16,
      offset: 8
    }
  }
};
const PhoneInput = ({ id, value = {}, onChange }) => {
  const [prefix, setPrefix] = useState('86');
  const [phone, setPhone] = useState('');
  const triggerChange = changedValue => {
    onChange?.({ ...value, ...changedValue });
  };
  const onPrefixChange = newPrefix => {
    if (!('prefix' in value)) {
      setPrefix(newPrefix);
    }
    triggerChange({ prefix: newPrefix });
  };
  const onPhoneChange = e => {
    const newPhone = e.target.value;
    if (!('phone' in value)) {
      setPhone(newPhone);
    }
    triggerChange({ phone: newPhone });
  };
  return (
    <span id={id}>
      <Space.Compact block>
        <Select
          value={value.prefix || prefix}
          onChange={onPrefixChange}
          style={{ width: 70 }}
          options={[
            { label: '+86', value: '86' },
            { label: '+87', value: '87' }
          ]}
        />
        <Input value={value.phone || phone} onChange={onPhoneChange} style={{ width: '100%' }} />
      </Space.Compact>
    </span>
  );
};
const DonationInput = ({ id, value = {}, onChange }) => {
  const [amount, setAmount] = useState();
  const [currency, setCurrency] = useState('USD');
  const triggerChange = changedValue => {
    onChange?.({ ...value, ...changedValue });
  };
  const onAmountChange = newAmount => {
    if (!('amount' in value)) {
      setAmount(newAmount ?? undefined);
    }
    triggerChange({ amount: newAmount ?? undefined });
  };
  const onCurrencyChange = newCurrency => {
    if (!('currency' in value)) {
      setCurrency(newCurrency);
    }
    triggerChange({ currency: newCurrency });
  };
  return (
    <span id={id}>
      <Space.Compact block>
        <InputNumber
          value={value.amount ?? amount}
          onChange={onAmountChange}
          style={{ width: '100%' }}
        />
        <Select
          value={value.currency || currency}
          onChange={onCurrencyChange}
          style={{ width: 70 }}
          options={[
            { label: '$', value: 'USD' },
            { label: '¥', value: 'CNY' }
          ]}
        />
      </Space.Compact>
    </span>
  );
};

const options = [];
for (let i = 10; i < 36; i++) {
  options.push({
    label: i.toString(36) + i,
    value: i.toString(36) + i
  });
}

const handleChange = value => {
  console.log(`selected ${value}`);
};

const TemplateForm = () => {
  const [form] = Form.useForm();
  const onFinish = values => {
    console.log('Received values of form: ', values);
  };

  return (
    <Row>
      <Col span={24}>
        <SCard title="Form Group" description="New design form based on Ant Design Form">
          <Form
            {...formItemLayout}
            variant="filled"
            form={form}
            name="register"
            onFinish={onFinish}
            initialValues={{
              residence: ['zhejiang', 'hangzhou', 'xihu'],
              phone: { prefix: '86' },
              donation: { currency: 'USD' }
            }}
            style={{ maxWidth: 600 }}
            scrollToFirstError
          >
            <Form.Item
              name="email"
              label="E-mail"
              colon={false}
              rules={[
                {
                  type: 'email',
                  message: 'The input is not valid E-mail!'
                },
                {
                  required: true,
                  message: 'Please input your E-mail!'
                }
              ]}
            >
              <Input />
            </Form.Item>

            <Form.Item
              name="password"
              label="Password"
              colon={false}
              rules={[
                {
                  required: true,
                  message: 'Please input your password!'
                }
              ]}
              hasFeedback
            >
              <Input.Password />
            </Form.Item>

            <Form.Item
              name="confirm"
              label="Confirm Password"
              dependencies={['password']}
              colon={false}
              hasFeedback
              rules={[
                {
                  required: true,
                  message: 'Please confirm your password!'
                },
                ({ getFieldValue }) => ({
                  validator(_, value) {
                    if (!value || getFieldValue('password') === value) {
                      return Promise.resolve();
                    }
                    return Promise.reject(
                      new Error('The new password that you entered do not match!')
                    );
                  }
                })
              ]}
            >
              <Input.Password />
            </Form.Item>

            <Form.Item
              name="username"
              label="Username"
              colon={false}
              rules={[{ required: true, message: 'Please input your Username!' }]}
            >
              <Input prefix={<UserOutlined />} placeholder="Username" />
            </Form.Item>

            <Form.Item
              name="nickname"
              label="Nickname"
              colon={false}
              tooltip="What do you want others to call you?"
              rules={[{ required: true, message: 'Please input your nickname!', whitespace: true }]}
            >
              <Input />
            </Form.Item>

            <Form.Item
              name="balance"
              label="Balance"
              colon={false}
              tooltip="What do you want others to call you?"
              rules={[{ required: true, message: 'Please input your nickname!', whitespace: true }]}
            >
              <Input prefix="￥" suffix="RMB" />
            </Form.Item>

            <Form.Item
              name="residence"
              label="Habitual Residence"
              colon={false}
              rules={[
                { type: 'array', required: true, message: 'Please select your habitual residence!' }
              ]}
            >
              <Cascader options={residences} />
            </Form.Item>

            <Form.Item name="age" label="Age" colon={false}>
              <Slider defaultValue={30} />
            </Form.Item>

            <Form.Item name="ageRange" label="Age Range" colon={false}>
              <Slider range defaultValue={[20, 50]} />
            </Form.Item>

            <Form.Item
              name="phone"
              label="Phone Number"
              colon={false}
              rules={[{ required: true, message: 'Please input your phone number!' }]}
            >
              <PhoneInput />
            </Form.Item>

            <Form.Item
              name="donation"
              label="Donation"
              colon={false}
              rules={[{ required: true, message: 'Please input donation amount!' }]}
            >
              <DonationInput />
            </Form.Item>

            <Form.Item
              name="intro"
              label="Intro"
              colon={false}
              rules={[{ required: true, message: 'Please input Intro' }]}
            >
              <Input.TextArea autoSize={{ minRows: 5 }} showCount maxLength={100} />
            </Form.Item>

            <Form.Item
              name="sIntro"
              label="Textarea"
              colon={false}
              rules={[{ required: true, message: 'Please input Intro' }]}
            >
              <STextArea autoSize={{ minRows: 5 }} showCount maxLength={100} />
            </Form.Item>

            <Form.Item
              name="gender"
              label="Gender"
              colon={false}
              rules={[{ required: true, message: 'Please select gender!' }]}
            >
              <Select
                placeholder="select your gender"
                defaultValue={'male'}
                options={[
                  { label: 'Male', value: 'male' },
                  { label: 'Female', value: 'female' },
                  { label: 'Other', value: 'other' }
                ]}
              />
            </Form.Item>

            <Form.Item name="multipleSelect" label="Multiple Select" colon={false}>
              <Select
                mode="multiple"
                allowClear
                style={{ width: '100%' }}
                placeholder="Please select"
                defaultValue={['a10', 'c12']}
                onChange={handleChange}
                options={options}
              />
            </Form.Item>

            <Form.Item
              label="DatePicker"
              name="DatePicker"
              colon={false}
              rules={[{ required: true, message: 'Please input!' }]}
            >
              <DatePicker />
            </Form.Item>

            <Form.Item
              label="RangePicker"
              name="RangePicker"
              colon={false}
              rules={[{ required: true, message: 'Please input!' }]}
            >
              <RangePicker />
            </Form.Item>
            <Form.Item
              label="Captcha"
              colon={false}
              extra="We must make sure that your are a human."
            >
              <Row gutter={8}>
                <Col span={12}>
                  <Form.Item
                    name="captcha"
                    noStyle
                    rules={[{ required: true, message: 'Please input the captcha you got!' }]}
                  >
                    <Input />
                  </Form.Item>
                </Col>
                <Col span={12}>
                  <Button size="large">Get captcha</Button>
                </Col>
              </Row>
            </Form.Item>

            <Form.Item
              name="agreement"
              valuePropName="checked"
              colon={false}
              rules={[
                {
                  validator: (_, value) =>
                    value ? Promise.resolve() : Promise.reject(new Error('Should accept agreement'))
                }
              ]}
              {...tailFormItemLayout}
            >
              <Checkbox>
                I have read the <a href="">agreement</a>
              </Checkbox>
            </Form.Item>
            <Form.Item {...tailFormItemLayout}>
              <Button type="primary" htmlType="submit">
                Register
              </Button>
            </Form.Item>
          </Form>
        </SCard>
      </Col>
    </Row>
  );
};
export default TemplateForm;
