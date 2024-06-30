import { useEffect } from "react";
import { useEditorStore } from "./editor_store";

const SCALE_STEP = 0.8;

export function EditorPanel() {
  const { editor } = useEditorStore();
  const zoomIn = () => {
    if (editor) {
      const zoom = editor.getZoom();
      editor.setZoom(zoom / SCALE_STEP);
    }
  };
  const zoomOut = () => {
    if (editor) {
      const zoom = editor.getZoom();
      editor.setZoom(zoom * SCALE_STEP);
    }
  };

  useEffect(() => { }, [editor]);

  return (
    <div className="stack">
      <div>
        <div className="box">
          <div className="cluster">
            <div>
              <button onClick={zoomIn}>Zoom in</button>
              <button onClick={zoomOut}>Zoom out</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
