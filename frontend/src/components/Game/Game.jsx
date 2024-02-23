import React, { useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import CardFlip from "../Common/CardFlip";
import Card from "../Common/Card";
import { drawCard, startGame, restartGame } from "../../features/gameSlice";

const Game = () => {
  const dispatch = useDispatch();
  const isUserValid = useSelector((state) => state.game.isvalid);
  const gamedata = useSelector((state) => state.game.deck);
  const drawnCards = useSelector((state) => state.game.drawnCards);
  const isGameOver = useSelector((state) => state.game.gameover);
  const isGameActive = useSelector((state) => state.game.isGameActive);
  const username = useSelector((state) => state.game.username);

  const [isUserWins, setIsUserWins] = useState(false);

  const userWins = async () => {
    try {
      const response = await fetch(
        `https://exploding-kittens-7.onrender.com/users/${username}/wins`,
        {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to update user wins");
      }

      setIsUserWins(true);
    } catch (error) {
      console.error(error);
    }
  };

  const handleStartGame = async () => {
    await dispatch(startGame());
    console.log(isGameActive);
  };

  const handleRestartGame = () => {
    dispatch(restartGame());
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-[#591718] text-[#FDDEA8]">
      <h1 className="text-4xl font-extrabold mb-4">Exploding Kittens</h1>

      {isUserValid ? (
        <>
          {isGameActive && !isGameOver ? (
            <div>
              {isGameOver && (
                <p className="text-2xl text-red-500 mt-4">Game Over!</p>
              )}
              <div className="flex items-center justify-center flex-wrap">
                {drawnCards.length > 0 ? (
                  drawnCards.map((card, index) => (
                    <Card card={card} key={index} />
                  ))
                ) : (
                  "No drawn cards"
                )}
              </div>

              <div className="flex flex-col items-center mt-4">
                <div className="flex flex-wrap gap-2">
                  {gamedata.map((card, index) => (
                    <CardFlip card={card} key={index} />
                  ))}
                </div>

                <button
                  className="border-2 border-white rounded-lg px-3 py-2 cursor-pointer hover:bg-[#E32F2E] hover:text-red-200 mt-4"
                  onClick={handleRestartGame}
                >
                  Restart Game
                </button>

                {gamedata.length === 0 && isUserWins && (
                  <p className="text-2xl text-green-500 mt-4">
                    Congratulations! You won
                  </p>
                )}
              </div>
            </div>
          ) : (
            <div>
              {isGameOver && (
                <p className="text-2xl text-red-500 mt-4">Game Over!</p>
              )}
              <button
                className="border-2 border-white rounded-lg px-3 py-2 cursor-pointer hover:bg-[#E32F2E] hover:text-red-200 mt-4"
                onClick={handleStartGame}
              >
                Start Game
              </button>
            </div>
          )}
        </>
      ) : (
        <p className="text-center text-2xl text-white mt-4">
          Please create a valid user profile to play.
        </p>
      )}
    </div>
  );
};

export default Game;
