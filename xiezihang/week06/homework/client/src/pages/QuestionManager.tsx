import React, { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import {
  Button,
  Input,
  Space,
  Table,
  Dropdown,
  Menu,
  message,
  Row,
  Col,
  Modal,
  Form,
  Select,
  Radio,
  Checkbox,
  Popconfirm
} from 'antd'
import type { ColumnsType } from 'antd/es/table'
import { DownOutlined } from '@ant-design/icons'
interface Question {
  id: number
  title: string
  type: string
  difficulty?: string
  options?: string
  answer?: string
  key: string
}

import axios from 'axios'

const { Option } = Select

const QuestionManager: React.FC = () => {
  const navigate = useNavigate()
  const [selectedType, setSelectedType] = useState('全部')
  const [searchKeyword, setSearchKeyword] = useState('')
  const [data, setData] = useState<Question[]>([])
  const [selectedRowKeys, setSelectedRowKeys] = useState<number[]>([])
  const [loading, setLoading] = useState(false)
  const [modalVisible, setModalVisible] = useState(false)
  const [form] = Form.useForm()
  const [questionType, setQuestionType] = useState('单选题')
  const [pagination, setPagination] = useState({ current: 1, pageSize: 10, total: 0 })
  const [isEdit, setIsEdit] = useState(false)
  const [currentQuestion, setCurrentQuestion] = useState<Question | null>(null)

  useEffect(() => {
    fetchData(pagination.current, pagination.pageSize)
  }, [selectedType, searchKeyword])

  const fetchData = async (page: number, pageSize: number) => {
    try {
      setLoading(true)
      const res = await axios.get('/api/questions', {
        params: {
          type: selectedType === '全部' ? '' : selectedType,
          keyword: searchKeyword,
          page,
          pageSize
        }
      })
      const list = (res.data.data || []).map((item: any) => ({
        ...item,
        key: item.id
      }))
      setData(list)
      setPagination(prev => ({ ...prev, total: res.data.total || 0, current: page, pageSize }))
    } catch (err) {
      message.error('加载题目失败')
    } finally {
      setLoading(false)
    }
  }

  const handleDelete = async (key: string) => {
    try {
      await axios.delete(`/api/questions/${key}`)
      message.success('已删除')
      fetchData(pagination.current, pagination.pageSize)
    } catch {
      message.error('删除失败')
    }
  }

  const handleSearch = (value: string) => {
    setSearchKeyword(value)
  }

  const handleTypeChange = (type: string) => {
    setSelectedType(type)
  }

  const handleBatchDelete = async () => {
    try {
      await axios.post('/api/questions/delete', {
        ids: selectedRowKeys
      })
      setSelectedRowKeys([])
      message.success('批量删除成功')
      fetchData(pagination.current, pagination.pageSize)
    } catch {
      message.error('批量删除失败')
    }
  }

  const handleTableChange = (paginationConfig: any) => {
    setPagination(paginationConfig)
    fetchData(paginationConfig.current, paginationConfig.pageSize)
  }

  const handleManualCreate = () => {
    setIsEdit(false)
    setCurrentQuestion(null)
    form.resetFields()
    setQuestionType('单选题')
    setModalVisible(true)
  }

  const handleEdit = (record: Question) => {
    setIsEdit(true)
    setCurrentQuestion(record)
    setModalVisible(true)
    
    const options = record.options?.split('|') || []
    
    form.setFieldsValue({
      type: record.type,
      title: record.title,
      difficulty: record.difficulty,
      optionA: options[0] || '',
      optionB: options[1] || '',
      optionC: options[2] || '',
      optionD: options[3] || '',
      answer: record.type === '多选题' ? record.answer?.split('') : record.answer
    })
    
    setQuestionType(record.type)
  }

  const handleModalOk = async () => {
    try {
      const values = await form.validateFields()
      const payload = {
        title: values.title,
        type: values.type,
        difficulty: values.difficulty,
        options: ['A', 'B', 'C', 'D'].map(letter => values[`option${letter}`]).join('|'),
        answer: values.type === '多选题' ? (values.answer?.join('') || '') : values.answer,
      }
      if (values.type === '编程题') {
        payload.options = ''
      }

      if (isEdit && currentQuestion) {
        await axios.put(`/api/questions/${currentQuestion.id}`, payload)
        message.success('修改成功')
      } else {
        await axios.post('/api/questions', payload)
        message.success('创建成功')
      }
      
      setModalVisible(false)
      setIsEdit(false)
      setCurrentQuestion(null)
      fetchData(pagination.current, pagination.pageSize)
    } catch {
      message.error(isEdit ? '修改失败' : '创建失败')
    }
  }

  const handleModalCancel = () => {
    setModalVisible(false)
    setIsEdit(false)
    setCurrentQuestion(null)
    form.resetFields()
  }

  const columns: ColumnsType<Question> = [
    {
      title: '题目',
      dataIndex: 'title',
      render: text => <a>{text}</a >,
    },
    {
      title: '题型',
      dataIndex: 'type',
      align: 'center',
    },
    {
    title: '难度',
    dataIndex: 'difficulty',
    align: 'center',
    },
    {
      title: '操作',
      align: 'center',
      render: (_, record) => (
        <Space size="middle">
          <a onClick={() => handleEdit(record)}>编辑</a>
          <Popconfirm
            title="确定要删除这道题吗？"
            onConfirm={() => handleDelete(record.id.toString())}
            okText="确定"
            cancelText="取消"
          >
            <a style={{ color: 'red' }}>删除</a>
          </Popconfirm>
        </Space>
      ),
    },
  ]

  const rowSelection = {
    selectedRowKeys,
    onChange: (newSelectedRowKeys: React.Key[]) => {
      setSelectedRowKeys(newSelectedRowKeys as number[])
    },
  }

  const menu = (
    <Menu
      onClick={({ key }) => {
        if (key === 'ai') {
          navigate('/ai-generate')
        } else if (key === 'manual') {
          handleManualCreate()
        }
      }}
    >
      <Menu.Item key="ai">AI 出题</Menu.Item>
      <Menu.Item key="manual">手动出题</Menu.Item>
    </Menu>
  )

  return (
    <div style={{ background: '#fff', padding: 24, minHeight: '100%' }}>
      <Row justify="space-between" align="middle" gutter={16} style={{ marginBottom: 16 }}>
        <Col flex="auto">
          <Space>
            <span>题型：</span>
            {['全部', '单选题', '多选题', '编程题'].map(type => (
              <Button
                key={type}
                type={selectedType === type ? 'primary' : 'default'}
                onClick={() => handleTypeChange(type)}
              >
                {type}
              </Button>
            ))}
          </Space>
        </Col>
        <Col flex="500px">
          <Input.Search
            allowClear
            placeholder="请输入试题名称"
            onSearch={handleSearch}
            style={{ width: '100%' }}
          />
        </Col>
      </Row>

      <Row justify="end" align="middle" gutter={16} style={{ marginBottom: 16 }}>
        <Col>
          <Dropdown overlay={menu}>
            <Button type="primary">
              出题 <DownOutlined />
            </Button>
          </Dropdown>
        </Col>
        <Col>
          <Popconfirm
            title="确定要删除选中的题目吗？"
            onConfirm={handleBatchDelete}
            okText="确定"
            cancelText="取消"
            disabled={selectedRowKeys.length === 0}
          >
            <Button danger disabled={selectedRowKeys.length === 0}>
              批量删除
            </Button>
          </Popconfirm>
        </Col>
      </Row>

      <Table
        rowSelection={rowSelection}
        columns={columns}
        dataSource={data}
        loading={loading}
        pagination={{
          current: pagination.current,
          pageSize: pagination.pageSize,
          total: pagination.total,
          showSizeChanger: true,
          showQuickJumper: true,
          showTotal: total => `共 ${total} 条`,
          position: ['bottomCenter'],
          locale: {
            items_per_page: '条/页',
            jump_to: '跳至',
            jump_to_confirm: '确定',
            page: '页',
            prev_page: '上一页',
            next_page: '下一页',
            prev_5: '向前 5 页',
            next_5: '向后 5 页',
            prev_3: '向前 3 页',
            next_3: '向后 3 页',
          }
        }}
        onChange={handleTableChange}
        locale={{ emptyText: '暂无数据' }}
      />

      <Modal
        title={isEdit ? "编辑题目" : "创建题目"}
        open={modalVisible}
        onCancel={handleModalCancel}
        onOk={handleModalOk}
        width={800}
      >
        <Form form={form} layout="vertical">
          <Row gutter={16}>
            <Col span={12}>
              <Form.Item name="type" label="题型" rules={[{ required: true }]}> 
                <Select onChange={(value) => setQuestionType(value)}>
                  <Option value="单选题">单选题</Option>
                  <Option value="多选题">多选题</Option>
                  <Option value="编程题">编程题</Option>
                </Select>
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="language" label="语言" initialValue="go语言" rules={[{ required: true }]}>
                <Select>
                  <Option value="go语言">go语言</Option>
                  <Option value="java">Java</Option>
                  <Option value="python">Python</Option>
                  <Option value="JavaSprit">JavaSprit</Option>
                </Select>
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item name="difficulty" label="难度" rules={[{ required: true }]}>
                <Select>
                  <Option value="简单">简单</Option>
                  <Option value="中等">中等</Option>
                  <Option value="困难">困难</Option>
                </Select>
              </Form.Item>
            </Col>
          </Row>

          <Form.Item name="title" label="标题" rules={[{ required: true }]}> 
            <Input />
          </Form.Item>

          <Form.Item name="content" label="内容" rules={[{ required: false }]}> 
            <Input.TextArea rows={4} showCount maxLength={500} />
          </Form.Item>

          {questionType !== '编程题' && (
            <>
              {['A', 'B', 'C', 'D'].map((label,) => (
                <Form.Item
                  key={label}
                  name={`option${label}`}
                  label={`选项 ${label}`}
                  rules={[{ required: true }]}
                >
                  <Input />
                </Form.Item>
              ))}

              <Form.Item 
                name="answer" 
                label="答案" 
                rules={[
                  { required: true, message: '请选择答案' },
                  {
                    validator: (_, value) => {
                      if (questionType === '多选题' && (!value || value.length < 2)) {
                        return Promise.reject('多选题至少需要选择两个选项');
                      }
                      return Promise.resolve();
                    }
                  }
                ]}
              > 
                {questionType === '单选题' ? (
                  <Radio.Group>
                    <Radio value="A">A</Radio>
                    <Radio value="B">B</Radio>
                    <Radio value="C">C</Radio>
                    <Radio value="D">D</Radio>
                  </Radio.Group>
                ) : (
                  <Checkbox.Group>
                    <Space>
                      <Checkbox value="A" style={{ marginRight: 8 }}>A</Checkbox>
                      <Checkbox value="B" style={{ marginRight: 8 }}>B</Checkbox>
                      <Checkbox value="C" style={{ marginRight: 8 }}>C</Checkbox>
                      <Checkbox value="D">D</Checkbox>
                    </Space>
                  </Checkbox.Group>
                )}
              </Form.Item>
            </>
          )}
        </Form>
      </Modal>
    </div>
  )
}

export default QuestionManager