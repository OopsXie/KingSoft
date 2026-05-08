function add(...nums) {
    return nums.reduce((acc, val) => acc + Number(val), 0);
  }
  
  function subtract(...nums) {
    if (nums.length === 0) return 0;
    return nums.slice(1).reduce((acc, val) => acc - Number(val), Number(nums[0]));
  }
  
  function multiply(...nums) {
    return nums.reduce((acc, val) => acc * Number(val), 1);
  }
  
  function divide(...nums) {
    if (nums.length === 0) return 0;
    return nums.slice(1).reduce((acc, val) => {
      if (Number(val) === 0) {
        throw new Error('除数不能为0');
      }
      return acc / Number(val);
    }, Number(nums[0]));
  }
  
  module.exports = {
    add,
    subtract,
    multiply,
    divide
  };
  