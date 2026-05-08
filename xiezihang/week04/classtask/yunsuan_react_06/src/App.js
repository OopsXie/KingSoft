import React, { useState, useEffect } from 'react';
import './App.css';

const App = () => {
  const [tasks, setTasks] = useState(() => {
    const saved = localStorage.getItem('tasks');
    return saved ? JSON.parse(saved) : [];
  });
  const [input, setInput] = useState('');
  const [theme, setTheme] = useState('light');

  useEffect(() => {
    localStorage.setItem('tasks', JSON.stringify(tasks));
  }, [tasks]);

  const addTask = () => {
    if (!input.trim()) return;
    const now = new Date();
    const timestamp = `${now.getFullYear()}年${now.getMonth() + 1}月${now.getDate()}日 ${now.toLocaleTimeString('zh-CN', { hour12: false })}`;
    const newTask = { text: input.trim(), done: false, time: timestamp };
    setTasks([...tasks, newTask]);
    setInput('');
  };

  const toggleDone = (index) => {
    const newTasks = [...tasks];
    newTasks[index].done = !newTasks[index].done;
    setTasks(newTasks);
  };

  const deleteTask = (index) => {
    const newTasks = tasks.filter((_, i) => i !== index);
    setTasks(newTasks);
  };

  const toggleTheme = () => {
    setTheme(theme === 'light' ? 'dark' : 'light');
  };

  const total = tasks.length;
  const completed = tasks.filter(t => t.done).length;

  return (
    <div className={`app-container ${theme}`}>
      <h1 className="title">Todo List</h1>
      <div className="input-container">
        <input
          value={input}
          onChange={e => setInput(e.target.value)}
          onKeyDown={e => e.key === 'Enter' && addTask()}
          placeholder="What needs to be done?"
        />
        <button onClick={addTask}>Add</button>
        <button onClick={toggleTheme}>切换主题</button>
      </div>
      <ul className="task-list">
        {tasks.map((task, index) => (
          <li key={index} className="task-item">
            <div className="left">
              <input type="checkbox" checked={task.done} onChange={() => toggleDone(index)} />
              <span className={task.done ? 'done' : ''}>{task.text}</span>
            </div>
            <div className="right">
              <span className="timestamp">{task.time}</span>
              <button className="delete-btn" onClick={() => deleteTask(index)}>🗑️</button>
            </div>
          </li>
        ))}
      </ul>
      <div className="footer">
        总任务数: {total}，已完成: {completed}
      </div>
    </div>
  );
};

export default App;
