import { FabricJSCanvas, useFabricJSEditor } from "fabricjs-react";
import "./editor.css"

export function Editor() {
  const { editor, onReady } = useFabricJSEditor();
  const onAddCircle = () => {
    editor?.addCircle();
  };
  const onAddRectangle = () => {
    editor?.addRectangle();
  };
  return (
    <>
      <button onClick={onAddCircle}>Add circle</button>
      <button onClick={onAddRectangle}>Add Rectangle</button>
      <div className="w-full editor__canvas-container">
        <FabricJSCanvas onReady={onReady} />
      </div>
    </>
  );
}
