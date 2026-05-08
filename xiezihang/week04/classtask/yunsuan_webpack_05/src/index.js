import './style.css';

const taskInput = document.getElementById('taskInput');
const taskList = document.getElementById('taskList');
const totalCountEl = document.getElementById('totalCount');
const doneCountEl = document.getElementById('doneCount');

taskInput.addEventListener('keypress', function (e) {
  if (e.key === 'Enter') addTask();
});

function saveTasks() {
  const tasks = [];
  document.querySelectorAll('.task').forEach(task => {
    const name = task.querySelector('span').textContent;
    const timestamp = task.querySelector('.timestamp').textContent;
    const completed = task.classList.contains('completed');
    tasks.push({ name, timestamp, completed });
  });
  localStorage.setItem('tasks', JSON.stringify(tasks));
  updateCounts();
}

function updateCounts() {
  const total = document.querySelectorAll('.task').length;
  const done = document.querySelectorAll('.task.completed').length;
  totalCountEl.textContent = total;
  doneCountEl.textContent = done;
}

function addTask(nameFromStorage = '', timestampFromStorage = '', completedFromStorage = false) {
  const taskName = nameFromStorage || taskInput.value.trim();
  if (!taskName) return;

  const now = new Date();
  const timestamp = timestampFromStorage || `${now.getFullYear()}年${now.getMonth() + 1}月${now.getDate()}日 ${now.toLocaleTimeString('zh-CN', { hour12: false })}`;

  const li = document.createElement('li');
  li.className = 'task';

  const left = document.createElement('div');
  left.className = 'left';

  const checkbox = document.createElement('input');
  checkbox.type = 'checkbox';
  checkbox.checked = completedFromStorage;
  checkbox.onchange = () => {
    li.classList.toggle('completed', checkbox.checked);
    saveTasks();
  };

  const span = document.createElement('span');
  span.textContent = taskName;

  left.appendChild(checkbox);
  left.appendChild(span);

  const right = document.createElement('div');
  right.style.display = 'flex';
  right.style.alignItems = 'center';

  const time = document.createElement('span');
  time.className = 'timestamp';
  time.textContent = timestamp;

  const del = document.createElement('button');
  del.className = 'delete-btn';
  del.innerHTML = '🗑️';
  del.onclick = () => {
    li.remove();
    saveTasks();
  };

  right.appendChild(time);
  right.appendChild(del);

  li.appendChild(left);
  li.appendChild(right);

  if (completedFromStorage) li.classList.add('completed');

  taskList.appendChild(li);
  if (!nameFromStorage) taskInput.value = '';
  saveTasks();
}

function toggleTheme() {
  document.body.classList.toggle('dark');
  document.body.classList.toggle('light');
}

document.querySelector('.add-btn').addEventListener('click', addTask);
document.querySelector('.theme-btn').addEventListener('click', toggleTheme);

function loadTasks() {
  const saved = JSON.parse(localStorage.getItem('tasks')) || [];
  saved.forEach(task => {
    addTask(task.name, task.timestamp, task.completed);
  });
}

loadTasks();
