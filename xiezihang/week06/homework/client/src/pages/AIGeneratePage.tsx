import React, { useState } from 'react'
import { Card, Select, InputNumber, Button, message, Row, Col, } from 'antd'
import { InfoCircleOutlined } from '@ant-design/icons'
import axios from 'axios'

const { Option } = Select

const AIGeneratePage: React.FC = () => {
  const [type, setType] = useState<string>('单选题')
  const [count, setCount] = useState<number>(1)
  const [language, setLanguage] = useState<string>('go语言')
  const [difficulty, setDifficulty] = useState<string>('简单')
  const [result, setResult] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(false)
  const [questions, setQuestions] = useState<any[]>([])

  const handleGenerate = async () => {
    if (count < 1 || count > 10) {
      message.warning('题目数量必须在1到10之间')
      return
    }
    try {
      setLoading(true)
      const res = await axios.post('/api/questions/ai_generate', {
        type,
        count,
        language,
        difficulty
      })
      if (res.data.code === 0) {
        setQuestions(res.data.data || [])
        setResult('')
      } else {
        setQuestions([])
        setResult(res.data.msg || '生成失败')
      }
    } catch (err) {
      setQuestions([])
      setResult('生成失败')
      message.error('生成失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={{ background: '#fff', padding: 24, minHeight: '100%' }}>
      <Row gutter={24}>
        <Col span={6}>
          <Card title="AI 生成试题" bordered={false}>
            <p>题型：</p>
            <Select value={type} onChange={setType} style={{ width: '100%', marginBottom: 12 }}>
              <Option value="单选题">单选题</Option>
              <Option value="多选题">多选题</Option>
              <Option value="编程题">编程题</Option>
            </Select>

            <p>题目数量：</p>
            <div style={{ marginBottom: 12 }}>
              <InputNumber
                min={1}
                max={10}
                value={count}
                onChange={(v) => setCount(v || 1)}
                style={{ width: '100%' }}
                placeholder="请输入1-10之间的数字"
              />
              <div style={{ marginTop: 4, color: '#999', fontSize: '12px' }}>
                <InfoCircleOutlined style={{ marginRight: 4 }} />
                题目数量必须在1到10之间
              </div>
            </div>

            <p>语言：</p>
            <Select value={language} onChange={setLanguage} style={{ width: '100%', marginBottom: 12 }}>
              <Option value="go语言">go语言</Option>
              <Option value="java">Java</Option>
              <Option value="python">Python</Option>
              <Option value="JavaSprit">JavaSprit</Option>
            </Select>

            <p>难度：</p>
            <Select value={difficulty} onChange={setDifficulty} style={{ width: '100%', marginBottom: 12 }}>
              <Option value="简单">简单</Option>
              <Option value="中等">中等</Option>
              <Option value="困难">困难</Option>
            </Select>

            <Button type="primary" block onClick={handleGenerate} loading={loading}>
              生成并预览题库
            </Button>
          </Card>
        </Col>

        <Col span={18}>
          <Card title="AI 生成区域" bordered style={{ height: '100%' }}>
            {result ? (
              <div>{result}</div>
            ) : (
              questions.length === 0 ? (
                <div>暂无题目</div>
              ) : (
                questions.map((q, idx) => (
                  <div key={q.id || idx} style={{ marginBottom: 24 }}>
                    <div>
                      <b>题目{idx + 1}：</b>{q.title}
                    </div>
                    {q.options && q.options.split('|').length > 1 && (
                      <ol type="A" style={{ marginLeft: 24 }}>
                        {q.options.split('|').map((opt: string, i: number) => (
                          <li key={i}>{opt.replace(/^[A-D][.、:：\\s]?/, '')}</li>
                        ))}
                      </ol>
                    )}
                    <div>
                      <b>答案：</b>
                      {q.answer}
                    </div>
                    <div>
                      <b>难度：</b>
                      {q.difficulty}
                    </div>
                    <div>
                      <b>题型：</b>
                      {q.type}
                    </div>
                  </div>
                ))
              )
            )}
          </Card>
        </Col>
      </Row>
    </div>
  )
}

export default AIGeneratePage