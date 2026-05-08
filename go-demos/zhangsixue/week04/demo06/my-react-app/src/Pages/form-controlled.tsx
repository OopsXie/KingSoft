import React, { useRef, useEffect } from "react";

// 非受控组件
function UncontrolledInput() {
  const inputRef = useRef<HTMLInputElement>(null);

  const handleClick = () => {
    alert(`Input value: ${inputRef?.current?.value}`);
  };
  console.log('render UncontrolledInput')
  return (
    <div>
      <input type="text" ref={inputRef} defaultValue="Default value" />
      <button onClick={handleClick}>Show Input Value</button>
    </div>
  );
}

// 受控组件
function ControlledInput() {
  const [value, setValue] = React.useState("");

  const handleClick = () => {
    alert(value);
  };
  console.log('render ControlledInput')
  return (
    <div>
      <input
        type="text"
        onChange={(e) => setValue(e.target.value)}
        defaultValue="Default value"
      />
      <button onClick={handleClick}>Show Input Value</button>
    </div>
  );
}
export default () => {
  return (
    <div>
      <UncontrolledInput />
      <hr />
      <ControlledInput />
    </div>
  );
};
