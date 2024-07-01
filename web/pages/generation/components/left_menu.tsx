import update from "immutability-helper";
import { useCallback, useEffect, useState } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import { Card } from "./card";
import { Item, useGenerationStore } from "./generation_store";

export function LeftMenu() {
  const { priorities, setPriorities } = useGenerationStore();
  const [cards, setCards] = useState(priorities);

  const moveCard = useCallback((dragIndex: number, hoverIndex: number) => {
    setCards((prevCards: Item[]) =>
      update(prevCards, {
        $splice: [
          [dragIndex, 1],
          [hoverIndex, 0, prevCards[dragIndex] as Item],
        ],
      }),
    );
  }, []);

  useEffect(() => {
    setPriorities(cards);
  }, [cards]);

  const renderCard = useCallback(
    (card: { id: number; text: string }, index: number) => {
      return (
        <Card
          key={card.id}
          index={index}
          id={card.id}
          text={card.text}
          moveCard={moveCard}
        />
      );
    },
    [],
  );

  return (
    <DndProvider backend={HTML5Backend}>
      <div className="w-full">
        {cards.map((card, i) => renderCard(card, i))}
      </div>
    </DndProvider>
  );
}
