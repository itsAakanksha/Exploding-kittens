import React, { useState, useEffect } from 'react';

function Leaderboard() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('https://exploding-kittens-7.onrender.com/leaderboard');

        if (!response.ok) {
          throw new Error('Failed to fetch leaderboard data');
        }

        const data = await response.json();
        setUsers(data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchData();
  }, []);

  return (
    <div className="bg-gradient-to-r from-purple-600 to-indigo-800 p-8 rounded-md shadow-md">
      <h1 className="text-3xl font-bold mb-6 text-white">Game Leaderboard</h1>
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-white">
          <thead>
            <tr>
             
              <th className="py-2 px-4 text-left text-white font-semibold">Username</th>
              <th className="py-2 px-4 text-left text-white font-semibold">Wins</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, index) => (
              <tr key={user.username}>
               
                <td className="py-2 px-4 whitespace-nowrap text-white">{user.username}</td>
                <td className="py-2 px-4 whitespace-nowrap text-white">{user.wins}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}

export default Leaderboard;





