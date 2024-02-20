import React from "react";
import Game from "./components/Game/Game.jsx";
import CreateUserForm from "./components/user/CreateUser.jsx";
import { useState, useEffect } from "react";
function App() {
  const [username, setUsername] = useState("");
  const [isCreated, setIsCreated] = useState(false);
  const [isUserValid, setIsUserValid] = useState(false);

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const createUser = async () => {
    try {
      const response = await fetch("http://localhost:8080/createuser", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username }),
      });

      if (!response.ok) {
        throw new Error("Failed to create user");
      }

      setIsCreated(true);
    } catch (error) {
      console.error(error);
    }
  };

  const checkUser = async () => {
    try {
      const response = await fetch("http://localhost:8080/users/" + username);

      if (response.status === 404) {
        
        setIsUserValid(false);
      } else if (!response.ok) {
        throw new Error("Failed to check user");
      } else {
        const data = await response.json();
        console.log(data.exists);
        setIsUserValid(true);
      }
    } catch (error) {
      console.error(error);
  
    }
  };

  useEffect(() => {
    if (username) {
      checkUser();
    }
  }, [isCreated]);

  return (
    <>
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-red-600 to-[#990000]">
      <div className="flex-col justify-center items-center mx-auto">
        {!isCreated && (
          <>
            <h1 className="text-4xl font-extrabold mb-6 text-red-600">Create Your Avatar!</h1>
            <input
              className="bg-gray-100 p-3 rounded-md w-full mb-4"
              type="text"
              placeholder="Enter your username"
              value={username}
              onChange={handleUsernameChange}
            />
            <div className="flex justify-center">
            <button
              className="w-full bg-yellow-400 text-red-600 px-6 py-3 rounded-md hover:bg-yellow-500 transition duration-300 "
              onClick={createUser}
            >
              Create
            </button>
            </div>
          </>
        )}
        {isCreated && (
          <>
            <h1 className="text-4xl font-extrabold mb-6 text-yellow-400 mx-auto">Welcome, {username}!</h1>
            {isUserValid ? (
             <Game/>
            ) : (
              <p className="text-red-600">User not found. Create an account to join the fun.</p>
            )}
          </>
        )}
      </div>
    </div>
    </>
  );
}

export default App;
