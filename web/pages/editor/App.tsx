import { useEffect } from "react";
import "./App.css";
import { Editor } from "./components/editor";
import { useDesignsStore } from "../../domain/design/store";
import { EditorPanel } from "./components/editor_panel";

export default function App() {
  const { setActiveDesign } = useDesignsStore();
  useEffect(() => {
    let params = new URLSearchParams(document.location.search);
    let id = params.get("design");
    setActiveDesign(Number(id));
  }, []);
  return (
    <>
      <main className="right-sidebar">
        <div>
          <div className="">
            <Editor />
          </div>
          <div className="box">
            <EditorPanel />
          </div>
        </div>
      </main>
    </>
  );
}
