import { useEffect } from "react";
import "./App.css";
import { Editor } from "./components/editor";
import { useDesignsStore } from "../../domain/design/store";

export default function App() {
  const { setActiveDesign } = useDesignsStore();
  useEffect(() => {
    let params = new URLSearchParams(document.location.search);
    let id = params.get("design");
    console.log(id)
    setActiveDesign(Number(id));
  }, []);
  return (
    <>
      <main className="right-sidebar">
        <div>
          <div className="box">
            <Editor />
          </div>
          <div className="box"></div>
        </div>
      </main>
    </>
  );
}
