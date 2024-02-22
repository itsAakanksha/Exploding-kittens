import React, { useState, useEffect } from "react";
import Game from "./components/Game/Game.jsx";
import { useDispatch } from "react-redux";
import { authenticate } from "./features/gameSlice";

function App() {
  const [username, setUsername] = useState("");
  const [isCreated, setIsCreated] = useState(false);
  const [isUserValid, setIsUserValid] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const dispatch = useDispatch();

  const handleUsernameChange = (event) => {
    setUsername(event.target.value);
  };

  const createUser = async () => {
    setIsLoading(true);
    setError(null);

    try {
      // const response = await fetch("http://localhost:8080/createuser", {
      const response = await fetch("https://exploding-kittens-7.onrender.com/createuser", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username }),
      });

      if (!response.ok) {
        throw new Error("Failed to create user");
      }

      setIsCreated(true);
    } catch (error) {
      setError(error);
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  const checkUser = async () => {
    try {
      const response = await fetch(`https://exploding-kittens-7.onrender.com/users/${username}`);

      if (response.ok) {
        const userData = await response.json();
        setIsUserValid(true);
        console.log("User data:", userData);
      } else {
        console.error("Error fetching user data:", response.statusText);
        setIsUserValid(false);
      }
    } catch (error) {
      console.error("Error fetching user data:", error.message);
      setIsUserValid(false);
    }
  };

  useEffect(() => {
    if (isCreated) {
      checkUser();
    }
  }, [isCreated]);

  useEffect(() => {
    if (isUserValid) {
      console.log("in app", isUserValid);
      dispatch(authenticate(isUserValid));
    }
  }, [isUserValid]);

  return (
    <>
      <div className="min-h-screen flex items-center justify-center bg-[#591718]">
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
                  className="w-full bg-yellow-400 text-red-600 px-6 py-3 rounded-md hover:bg-yellow-500 transition duration-300"
                  onClick={createUser}
                  disabled={isLoading} // Disable button while loading
                >
                  {isLoading ? "Creating..." : "Create"}
                </button>
              </div>
              {error && <p className="text-red-600">{error.message}</p>}
            </>
          )}
          {
            isCreated && isUserValid ? (
            <p>
                <Game />
             </p>
              ) : (
                <p className="text-red-600 mx-auto w-full">Create an account to join the fun.</p>
              )
              })
        
        </div>
      </div>
    </>
  );
              }

export default App;
