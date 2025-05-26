import React from "react";
import Logo from "../Logo/Logo";

interface LayoutProps {
  children: React.ReactNode;
  showLogo?: boolean;
  logoSize?: "sm" | "md" | "lg";
}

export const Layout: React.FC<LayoutProps> = ({
  children,
  showLogo = true,
  logoSize = "lg",
}) => {
  return (
    <div className="min-h-screen bg-gray-50">
      {showLogo && (
        <header className="bg-white shadow-sm border-b border-gray-200">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex justify-center py-4">
              <Logo size={logoSize} />
            </div>
          </div>
        </header>
      )}
      <main className="flex-1">{children}</main>
    </div>
  );
};

export default Layout