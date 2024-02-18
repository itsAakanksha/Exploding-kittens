import React from "react";
import { useDispatch, useSelector } from "react-redux";
import CardFlip from "../Common/CardFlip";
import Card from "../Common/Card";
import { drawCard, startGame, endGame, restartGame } from "../../features/gameSlice";

const Game = () => {
  const dispatch = useDispatch();
  const gamedata = useSelector((state) => state.game.deck);
  const drawnCards = useSelector((state) => state.game.drawnCards);
  const isGameOver = useSelector((state) => state.game.gameover);
  const isGameActive = useSelector((state) => state.game.isGameActive);

  const handleStartGame = () => {
    dispatch(startGame());
  };

  const handleRestartGame = () => {
    dispatch(restartGame());
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-[#990000]">
      <h1 className="text-2xl text-white mb-4">Exploding Kittens Game</h1>

      {isGameActive && !isGameOver ? (
        <div>
          {isGameOver && <p className="text-2xl text-red-500 mt-4">Game Over!</p>}
          <div className="flex items-center justify-center">
            {drawnCards.length > 0 ? (
              drawnCards.map((card, index) => <Card card={card} key={index} />)
            ) : (
              "No drawn cards"
            )}
          </div>

          <div className="flex flex-col items-center">
            <div className="flex items-center justify-center w-40 h-2 text-white text-lg cursor-pointer mb-4"></div>

            <div className="flex gap-2">
              {gamedata.map((card, index) => (
                <CardFlip card={card} key={index} />
              ))}
            </div>

            <button
              className="border-2 border-black rounded-lg px-3 py-2 text-white cursor-pointer hover:bg-[#E32F2E] hover:text-red-200 m-4"
              onClick={handleRestartGame}
            >
              Restart Game
            </button>
            {gamedata.length === 0 &&  <p className="text-2xl text-green-500 mt-4">Congratulations! You won</p>}
          </div>
        </div>
      ) : (
        <div>
          {isGameOver && <p className="text-2xl text-red-500 mt-4">Game Over!</p>}
          <button
            className="border-2 border-black rounded-lg px-3 py-2 text-white cursor-pointer hover:bg-[#E32F2E] hover:text-red-200 m-4"
            onClick={handleStartGame}
          >
            Start Game
          </button>
        </div>
      )}
    </div>
  );
};

export default Game;
