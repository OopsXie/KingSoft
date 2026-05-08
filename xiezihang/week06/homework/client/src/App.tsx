import kingsoftLogo from './assets/kingsoft.svg'

import { useState } from 'react'
import { Layout, Typography, theme, Menu } from 'antd'
import {
  FileTextOutlined,
  ProfileOutlined,
} from '@ant-design/icons'
import { useNavigate, BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import QuestionManager from './pages/QuestionManager'
import AIGeneratePage from './pages/AIGeneratePage'
import LearningNote from './pages/LearningNote'


const { Header, Sider, Content } = Layout
const { Title } = Typography

function MainLayout() {
  const [collapsed, setCollapsed] = useState(false)
  const navigate = useNavigate()
  const {
    token: { colorBgContainer },
  } = theme.useToken()

  const items = [
    {
      key: '1',
      icon: <FileTextOutlined />,
      label: '学习心得',
      onClick: () => navigate('/'),
    },
    {
      key: '2',
      icon: <ProfileOutlined />,
      label: '题库管理',
      onClick: () => navigate('/questions'),
    },
  ]

  return (
    <Layout style={{ height: '100vh',width: '100vw' }}>
      {/* 顶部导航栏 */}
      <Header style={{ background: '#001529', paddingLeft: 16, display: 'flex', alignItems: 'center'}}>
        <img
          src={kingsoftLogo}
          alt="logo"
          style={{ height: 32, marginRight: 16 }}
        />
        <Title level={4} style={{ color: '#fff', margin: 0 }}>
          武汉科技大学 谢子航 大作业
        </Title>
      </Header>

      {/* 左侧导航栏 + 主体内容 */}
      <Layout style={{ height: 'calc(100vh - 64px)' }}>
        <Sider
          collapsible
          collapsed={collapsed}
          onCollapse={setCollapsed}
          theme="dark"
          style={{ height: '100%' }}
        >
          <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']} items={items} />
        </Sider>

        <Content
          style={{
            padding: 24,
            background: colorBgContainer,
            height: '100%',
            width: '100%',
            overflow: 'auto',
            display: 'flex',
            flexDirection: 'column',
          }}
        >
          <Routes>
            <Route path="/" element={<LearningNote />} />
            <Route path="/questions" element={<QuestionManager />} />
            <Route path="/ai-generate" element={<AIGeneratePage />} />
          </Routes>
        </Content>
      </Layout>
    </Layout>
  )
}

export default function App() {
  return (
    <Router>
      <MainLayout />
    </Router>
  )
}
