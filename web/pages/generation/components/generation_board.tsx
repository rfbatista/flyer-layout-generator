import { useState } from "react";
import { Tabs } from "../../../components/tabs/tabs";
import "./generation_board.css";
import { TryBoard } from "./try_board";
import { BatchBoard } from "./batch_board";

export function GenerationBoard() {
  const [activeTab, setActiveTab] = useState("try");
  return (
    <>
      <div className="box">
        <Tabs
          items={[
            {
              name: "Batch",
              onClick: () => {
                setActiveTab("batch");
              },
              active: activeTab === "batch",
            },
            {
              name: "Try",
              onClick: () => {
                setActiveTab("try");
              },
              active: activeTab === "try",
            },
          ]}
        />
      </div>
      <div className="box">
        {activeTab === "try" ? <TryBoard /> : <BatchBoard />}
      </div>
    </>
  );
}
