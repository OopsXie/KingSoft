import React, { useEffect, useState } from 'react'
import ReactMarkdown from 'react-markdown'
import { Card, Spin, message } from 'antd'
import axios from 'axios'

const LearningNote: React.FC = () => {
  const [content, setContent] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(true)

  useEffect(() => {
    const fetchReadme = async () => {
      try {
        const res = await axios.get('/api/readme')
        if (res.data.code === 0) {
          setContent(res.data.data)
        } else {
          message.error(res.data.msg || '加载学习心得失败')
        }
      } catch {
        message.error('加载学习心得失败')
      } finally {
        setLoading(false)
      }
    }
    fetchReadme()
  }, [])

  return (
    <Card title="学习心得" style={{ minHeight: 400 }}>
      {loading ? <Spin /> : <ReactMarkdown>{content}</ReactMarkdown>}
    </Card>
  )
}

export default LearningNote
