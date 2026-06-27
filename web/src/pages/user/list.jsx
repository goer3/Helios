import { Table, Row, Col, Space, Tag, Button, Input } from 'antd';
import {
  EditOutlined,
  DeleteOutlined,
  PlusOutlined,
  CloudDownloadOutlined,
  SearchOutlined,
  FileExcelOutlined
} from '@ant-design/icons';
import Helmet from 'react-helmet';
import { TitleSuffix } from '@/components/text';
import SCard from '@/components/ui/s-card';

// 页面基础信息
const pageTitle = '用户列表';
const pageTitleEn = 'USER LIST';
const pageDescription = 'Administrators can use this list to modify the personal information and permissions of specific users.';

const columns = [
  {
    title: 'Full Name',
    dataIndex: 'name',
    key: 'name',
    showSorterTooltip: { target: 'full-header' },
    filters: [
      {
        text: 'Joe',
        value: 'Joe'
      },
      {
        text: 'Jim',
        value: 'Jim'
      },
      {
        text: 'Submenu',
        value: 'Submenu',
        children: [
          {
            text: 'Green',
            value: 'Green'
          },
          {
            text: 'Black',
            value: 'Black'
          }
        ]
      }
    ],
    onFilter: (value, record) => record.name.indexOf(value) === 0,
    fixed: 'start'
  },
  {
    title: 'Age',
    dataIndex: 'age',
    key: 'age',
    sorter: (a, b) => a.age - b.age,
    fixed: 'start'
  },
  {
    title: 'Column 1',
    dataIndex: 'address',
    key: '1'
  },
  {
    title: 'Column 2',
    dataIndex: 'hobby',
    key: '2',
    render: hobby => {
      const presets = [
        'magenta',
        'red',
        'volcano',
        'orange',
        'green',
        'cyan',
        'blue',
        'geekblue',
        'purple'
      ];
      return (
        <Space>
          {hobby.map(item => {
            const randomIndex = Math.floor(Math.random() * presets.length);
            return (
              <Tag variant="filled" color={presets[randomIndex]} key={item}>
                {item}
              </Tag>
            );
          })}
        </Space>
      );
    }
  },
  {
    title: 'Column 3',
    dataIndex: 'address',
    key: '3'
  },
  {
    title: 'Column 4',
    dataIndex: 'address',
    key: '4'
  },
  {
    title: 'Column 5',
    dataIndex: 'address',
    key: '5'
  },
  {
    title: 'Column 6',
    dataIndex: 'address',
    key: '6'
  },
  {
    title: 'Action',
    key: 'operation',
    fixed: 'end',
    render: () => (
      <Space>
        <Button icon={<EditOutlined />} color="blue" variant="link">
          Edit
        </Button>
        <Button icon={<DeleteOutlined />} color="pink" variant="link">
          Delete
        </Button>
      </Space>
    )
  }
];

const dataSource = Array.from({ length: 100 }).map((_, i) => ({
  key: i,
  name: `Edward King ${i}`,
  age: Math.floor(Math.random() * 100) + 1,
  hobby: ['Reading', 'Playing'],
  address: `London, Park Lane no. ${i}`
}));

const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
  },
  getCheckboxProps: record => ({
    disabled: record.name === 'Disabled User',
    name: record.name
  })
};

const UserList = () => {
  return (
    <>
      <Helmet>
        <title>{pageTitle} - {TitleSuffix}</title>
      </Helmet>
      <Row>
        <Col span={24}>
          <SCard title={`${pageTitle} / ${pageTitleEn}`} description={pageDescription}>
            <div className="s-card__btn-group">
              <Space>
                <Button icon={<PlusOutlined />} type="primary">
                  添加数据
                </Button>
                <Input size="small" placeholder="支持模糊搜索" suffix={<SearchOutlined />} />
              </Space>
              <Space>
                <Button icon={<CloudDownloadOutlined />}>下载模板</Button>
                <Button icon={<FileExcelOutlined />}>导出数据</Button>
              </Space>
            </div>
            <Table
              bordered
              columns={columns}
              dataSource={dataSource}
              scroll={{ x: 'max-content' }}
              rowSelection={{ type: 'checkbox', ...rowSelection }}
              expandable={{
                expandedRowRender: record => <h1>Hello World</h1>,
                rowExpandable: record => record.name !== 'Not Expandable'
              }}
              pagination={{
                showTotal: (total, range) => `总共 ${total} 条记录`,
                showSizeChanger: true,
                showQuickJumper: true
              }}
            />
          </SCard>
        </Col>
      </Row>
    </>
  );
};

export default UserList;
