import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="bg-black p-4">
     <div className="container mx-auto flex justify-between items-center">
 {    
        <Link to="/game" className="text-white text-xl font-bold">
          Game
        </Link>
}
        <div className="flex space-x-4">
          <Link to="/leaderboard" className="text-white">
            Leaderboard
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
