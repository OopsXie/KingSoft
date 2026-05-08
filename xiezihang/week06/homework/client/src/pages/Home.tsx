import React, { useEffect, useState } from 'react'
import ReactMarkdown from 'react-markdown'
import axios from 'axios'
import { Card, Spin, message } from 'antd'

const Home: React.FC = () => {
  const [content, setContent] = useState('')
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    const fetchReadme = async () => {
      setLoading(true)
      try {
        const res = await axios.get('/api/readme')
        if (res.data.code === 0) {
          setContent(res.data.data)
        } else {
          message.error('获取学习心得失败')
        }
      } catch (error) {
        message.error('请求失败')
      } finally {
        setLoading(false)
      }
    }

    fetchReadme()
  }, [])

  return (
    <Card title="学习心得" style={{ margin: 24 }}>
      {loading ? <Spin /> : <ReactMarkdown>{content}</ReactMarkdown>}
    </Card>
  )
}

export default Home
