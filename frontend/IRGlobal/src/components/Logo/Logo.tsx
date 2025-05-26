import React from "react";

interface LogoProps {
  size?: "sm" | "md" | "lg" | "xlg";
  className?: string;
}

export const Logo: React.FC<LogoProps> = ({ size = "md", className = "" }) => {
  const sizeClasses = {
    sm: { icon: "w-6 h-6", text: "text-lg", container: "gap-2" },
    md: { icon: "w-8 h-8", text: "text-xl", container: "gap-2.5" },
    lg: { icon: "w-10 h-10", text: "text-2xl", container: "gap-3" },
    xlg: { icon: "w-20 h-20", text: "text-4xl", container: "gap-4" },
  };

  const currentSize = sizeClasses[size];

  return (
    <div className={`flex items-center ${currentSize.container} ${className}`}>
      {/* Ícone moderno representando gráficos/finanças */}
      <div className="relative">
        <svg
          className={`${currentSize.icon} text-[#4160ff]`}
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
        >
          {/* Círculo externo representando globalidade */}
          <circle
            cx="12"
            cy="12"
            r="10"
            strokeWidth="1.5"
            className="opacity-30"
          />
          {/* Gráfico de barras crescente */}
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            d="M7 14l2-2 3 3 4-4"
          />
          {/* Pontos nos vértices do gráfico */}
          <circle cx="7" cy="14" r="1.5" fill="currentColor" />
          <circle cx="9" cy="12" r="1.5" fill="currentColor" />
          <circle cx="12" cy="15" r="1.5" fill="currentColor" />
          <circle cx="16" cy="11" r="1.5" fill="currentColor" />
        </svg>
        {/* Pequeno indicador de crescimento */}
        <div className="absolute -top-1 -right-1 w-2 h-2 bg-[#4160ff] rounded-full animate-pulse"></div>
      </div>

      {/* Texto da logo */}
      <div className="flex flex-col leading-none">
        <span
          className={`font-bold text-slate-800 ${currentSize.text} tracking-tight`}
        >
          IR<span className="text-[#4160ff]">Global</span>
        </span>
        <span className="text-xs text-slate-500 font-medium tracking-wider uppercase">
          Tax Solutions
        </span>
      </div>
    </div>
  );
};

export default Logo;
