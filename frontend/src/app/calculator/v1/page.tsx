'use client'

import React from 'react';
import { createClient } from "@connectrpc/connect";
import { CalculatorService } from "@/gen/calculator_pb";
import { createConnectTransport } from "@connectrpc/connect-web";
import Button from '@/components/Button';


const Calculator = () => {
  // This transport is going to be used throughout the app
  const transport = createConnectTransport({
    baseUrl: "/",
  });
  const client = createClient(CalculatorService, transport);
  const [expression, setExpression] = React.useState('');
  const [badResult, setBadResult] = React.useState(false);


  const handleButtonClick = (value) => {
    if (badResult) {
      setBadResult(!badResult)
      setExpression('')
    }
    setExpression(prev => prev + value);
  };

  const handleButtonDel = () => {
    setExpression(prev => prev.slice(0, prev.length - 1));
  };

  const handleButtonClr = () => {
    setExpression('');
  };

  const handleSubmit = async () => {
    const res = await client.calculate({
      expression: expression,
    });
    if (res.code != 0) {
      setExpression(`${res.message}`)
      setBadResult(true)
    } else {
      setExpression(res.result.toString())
    }
  };

  return (
    <div className="bg-[#1d1a1a] text-primary-foreground min-h-screen flex flex-col justify-end items-center">
      <div className="flex flex-col justify-end items-end p-4 w-full">
        <div className="text-4xl font-bold">{expression || ''}</div>
      </div>
      <div className="bg-[#201E1EFF] shadow-lg w-full h-full flex flex-col items-center">
        <div className="bg-[#1d1a1a] text-xl shadow-lg mt-2 grid grid-cols-4 gap-1 w-120">
          <Button value='(' onClick={handleButtonClick} />
          <Button value=')' onClick={handleButtonClick} />
          <Button value='clr' onClick={handleButtonClr} />
          <Button value='÷' onClick={handleButtonClick} />

          <Button value='7' onClick={handleButtonClick} />
          <Button value='8' onClick={handleButtonClick} />
          <Button value='9' onClick={handleButtonClick} />
          <Button value='×' onClick={handleButtonClick} />

          <Button value='4' onClick={handleButtonClick} />
          <Button value='5' onClick={handleButtonClick} />
          <Button value='6' onClick={handleButtonClick} />
          <Button value='-' onClick={handleButtonClick} />

          <Button value='1' onClick={handleButtonClick} />
          <Button value='2' onClick={handleButtonClick} />
          <Button value='3' onClick={handleButtonClick} />
          <Button value='+' onClick={handleButtonClick} />

          <Button value='del' onClick={handleButtonDel} />
          <Button value='0' onClick={handleButtonClick} />
          <Button value='.' onClick={handleButtonClick} />
          <Button value='=' className='bg-[#F55911FF]' onClick={handleSubmit} />
        </div>
      </div>
    </div>
  );
};


export default Calculator;