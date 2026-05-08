const calculator = require('./calculator');

const [,, operator, ...args] = process.argv;

if (!['+', '-', '*', '/'].includes(operator)) {
  console.log('无效的操作符，请使用 +、-、* 或 /');
  process.exit(1);
}

try {
  let result;
  switch (operator) {
    case '+':
      result = calculator.add(...args);
      break;
    case '-':
      result = calculator.subtract(...args);
      break;
    case '*':
      result = calculator.multiply(...args);
      break;
    case '/':
      result = calculator.divide(...args);
      break;
  }
  console.log('计算结果：', result);
} catch (error) {
  console.error('发生错误：', error.message);
}
