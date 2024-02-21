import React from 'react';

export default function Card({ card }) {
  return (
    <div className=" mx-1 card-container w-12 h-12 sm:w-16 sm:h-16 md:w-20 md:h-20 lg:w-24 lg:h-24 xl:w-32 xl:h-32 rounded-md overflow-hidden shadow-md transform transition-transform duration-300 hover:scale-105">
      <div
        className="card-inner w-full h-full perspective-400"
        style={{ perspective: "400px" }} // Set perspective for 3D effect
      >
        <div
          className="card-face w-full h-full rounded-xl bg-cover bg-center transition-transform duration-300 transform"
          style={{
            backgroundImage: `url(${card.img})`,
          }}
        />
      </div>
    </div>
  );
}
