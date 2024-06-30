import { useEffect } from "react";
import { useEditorStore } from "./editor_store";
import "../../../components/table/table.css";
import { useDesignsStore } from "../../../domain/design/store";

const SCALE_STEP = 0.8;

export function EditorPanel() {
  const { editor, layers, addActiveItem } = useEditorStore();
  const { activeDesign } = useDesignsStore();
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

  const createComponent = async () => {
    if (!editor) return;
    if (!activeDesign || !activeDesign.layout) return;
    if (!editor.getActiveObject()) {
      return;
    }
    const formData = new FormData();
    const o = editor.getActiveObjects;
    for (const e of o()) {
      const id = e.get("element_id");
      formData.append("elements[]", id);
    }
    try {
      const response = await fetch(
        `/editor/design/${activeDesign.id}/layout/${activeDesign.layout.id}/component`,
        {
          method: "POST",
          body: formData,
        },
      );
      if (response.ok) {
        const result = await response.json();
        console.log("Success:", result);
      } else {
        console.error("Error:", response.statusText);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  useEffect(() => {}, [editor]);

  return (
    <div className="stack">
      <div>
        <div className="box">
          <div className="cluster">
            <div>
              <button onClick={zoomIn}>Zoom in</button>
              <button onClick={zoomOut}>Zoom out</button>
              <button>Save</button>
              <button onClick={createComponent}>Create component</button>
            </div>
          </div>
        </div>
        <div className="box">
          <div className="table-container">
            <table>
              <tbody>
                {layers.map((l) => {
                  return (
                    <tr>
                      <td>
                        <span className="max-w-10 text-ellipsis">{l.name}</span>
                      </td>
                      <td>
                        <button onClick={() => addActiveItem(l)}>Select</button>
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}
