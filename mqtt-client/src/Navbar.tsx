import React from "react";

const sections = [
  {
    name: 'Home'
  },
  {
    name: 'Create device'
  },
  {
    name: 'About'
  }
]

const Navbar: React.FC = () => {
  return (
    <nav className="bg-gray-800 p-2">
      <div className="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex-shrink-0 flex items-center">
            <span className="text-white font-semibold text-lg">IoT Simulator</span>
          </div>
          <div className="hidden md:block">
            <div className="ml-10 flex items-baseline space-x-4">
              {Array.from(sections, (section) => (
                <a
                  href="#"
                  className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
                  key={section.name}
                >
                  {section.name}
                </a>               
              )
              )}
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;