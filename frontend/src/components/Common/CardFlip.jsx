import React, { useState, useEffect } from "react";
import { motion } from "framer-motion";
import back from '../../assets/backkitten.jpg';
import { useDispatch, useSelector } from "react-redux";
import { drawCard,flipCard } from "../../features/gameSlice";

export default function CardFlip({ card,index }) {
  // const flip = useSelector((state) => state.game.deck)
  const flip = card.isFlipped;
  console.log(flip);
  const dispatch = useDispatch();
 console.log(card.id);
  // Event handler for smooth click handling and transition control
  const handleFlip = () => {
    dispatch(flipCard(card.id))
    console.log("i");
  };
  useEffect(()=>{
    setTimeout(()=>{
   if(flip)
      {
        dispatch(drawCard(card))
      }
    },1000)
  },[flip])



  return (
    <div
      className="flip-card w-[200px] h-[200px] rounded-md"
      onClick={handleFlip}
    >
      <div
        className=" w-[100%] h-[100%] perspective-400"
        style={{ perspective: "400px" }} // Set perspective for 3D effect
      >
        <div
          className={`flip-card-front w-[100%] h-[100%] rounded-xl transition`}
          style={{
            backgroundImage: `url(${back})`,
            backgroundSize: "contain",
            backgroundPosition: "center",
            backgroundRepeat: "no-repeat",
            transform: flip ? "rotateY(180deg)" : "rotateY(0)", // Use transform for flip
            transition: "transform 0.3s ease-out", // Customize transition
          }}
        />
        <div
          className={`flip-card-back w-[100%] h-[100%] rounded-xl transition`}
          style={{
            backgroundImage: `url(${card.img})`,
            backgroundSize: "contain",
            backgroundPosition: "center",
            backgroundRepeat: "no-repeat",
            transform: flip ? "rotateY(0)" : "rotateY(180deg)", // Use transform for flip
            transition: "transform 0.3s ease-out", // Customize transition
          }}
        />
      </div>
    </div>
  );
}
