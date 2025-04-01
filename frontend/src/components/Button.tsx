// components/Button.tsx
import React from 'react';

interface ButtonProps {
  value: string;
  onClick: (value: string) => void;
  className?: string;
}

const Button: React.FC<ButtonProps> = ({ value, onClick, className = '' }) => {
  const baseClasses = 'bg-[#2B2A2AFF] text-secondary-foreground py-3 rounded';

  return (
    <button
      onClick={() => onClick(value)}
      className={`${className} ${baseClasses} `}
    >
      {value}
    </button>
  );
};

export default Button;