  import { createSlice } from "@reduxjs/toolkit";
  import cat from '../assets/cat.jpg'
  import defuse from '../assets/defuse.jpg'
  import explode from '../assets/exploding.jpg'
  import shuffle from '../assets/shuffle.jpg'
   

  const shuffleArray = (array) => {
    
    for (let i = array.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [array[i], array[j]] = [array[j], array[i]];
    }
    return array;
  };
  
  const initialState = {
    deck: [
      { type: "CAT", img:cat, id:0,isFlipped:false, },
      { type: "CAT",img:cat, id: 1,isFlipped:false, },
      { type: "EXPLODING_KITTEN",img:explode, id:2,isFlipped:false,},
      { type: "DEFUSE" ,img:defuse, id:3,isFlipped:false,},
      { type: "SHUFFLE",img:shuffle , id:4,isFlipped:false,},
    ],
    drawnCards: [],
    isGameActive: false,
    gameover: false,
    
  };
  const gameSlice = createSlice({
    name: "game",
    initialState,
    reducers: {
      drawCard: (state, action) => {
        const drawnCard = action.payload;
        state.deck = state.deck.filter((cardInDeck) => cardInDeck.id !== drawnCard.id);
        switch (drawnCard.type) {
          case "CAT":
            state.drawnCards.push(drawnCard);
            break;
          case "EXPLODING_KITTEN":
            if (state.drawnCards.some((card) => card.type === "DEFUSE")) {
              state.drawnCards = state.drawnCards.filter(
                (card) => card.type != "DEFUSE"
              );
            } else {
              state.gameover = true;
            }
            break;
          case "DEFUSE":
            state.drawnCards.push(drawnCard);
            break;
          case "SHUFFLE":
            state.isGameActive = true;
        state.gameover = false;
        state.drawnCards = [];
       state.deck = shuffleArray([...initialState.deck]);
            break;
        
        }
      },
      flipCard: (state, action) => {
        const cardId  = action.payload; // Assume payload contains the card's ID
        const cardIndex = state.deck.findIndex((card) => card.id === cardId);
  
        if (cardIndex !== -1) {
          state.deck[cardIndex].isFlipped = !state.deck[cardIndex].isFlipped;
        }
      },
      startGame: (state, action) => {
        state.isGameActive = true;
        state.gameover = false;
        state.drawnCards = [];
        state.deck = shuffleArray([...initialState.deck]);
      },
      endGame:(state,action)=>{
        state.isGameActive = false;
        state.gameover = true;
      },
      restartGame: (state, action) => {
        state.isGameActive = true;
        state.gameover = false;
        state.drawnCards = [];
       state.deck = shuffleArray([...initialState.deck]);
      },
   
      
    },
  });

  export const { drawCard,flipCard,startGame,endGame,restartGame} = gameSlice.actions;
  export default gameSlice.reducer;