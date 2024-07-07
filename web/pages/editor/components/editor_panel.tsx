import { useEffect, useState } from "react";
import "../../../components/table/table.css";
import { useDesignsStore } from "../../../domain/design/store";
import { defaultPriorities } from "../../../domain/layout/entities/priorities";
import { useEditorStore } from "./editor_store";

const SCALE_STEP = 0.8;

export function EditorPanel() {
  const { editor, layers, addActiveItem, save } = useEditorStore();
  const { activeDesign } = useDesignsStore();
  const [compType, setCompType] = useState("");
  const [isLoading, setLoading] = useState(false)

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
    formData.append("type", compType);
    const o = editor.getActiveObjects();
    for (const e of o) {
      const id = e.get("id");
      formData.append("elements[]", id);
    }
    try {
      setLoading(true)
      const response = await fetch(
        `/api/v1/editor/design/${activeDesign.id}/layout/${activeDesign.layoutID}/component`,
        {
          method: "POST",
          body: formData,
        },
      );
      if (response.ok) {
        const result = await response.json();
        console.log("Success:", result);
      setLoading(false)
        window.location.reload();
      } else {
        console.error("Error:", response.statusText);
      }
      setLoading(false)
    } catch (error) {
      setLoading(false)
      console.error("Error:", error);
    }
  };

  useEffect(() => {
    const data = layers.map((l) => ({ x: l.x, y: l.y, id: l.name }));
  }, [layers]);

  useEffect(() => { }, [editor]);

  return (
    <div className="stack">
      <div>
        <div className="box">
          <div className="cluster">
            <div>
              <button onClick={zoomIn}>Zoom in</button>
              <button onClick={zoomOut}>Zoom out</button>
              <button onClick={save}>Save</button>
            </div>
          </div>
        </div>
        <div>
          <div className="cluster">
            <div>
              <fieldset data-type="select">
                <label htmlFor="type">Component types</label>
                <span className="arrow" />
                <select
                  id="type"
                  name="tye"
                  onChange={(e) => setCompType(e.target.value)}
                >
                  <option selected></option>
                  {defaultPriorities.map((p) => {
                    return <option value={p.text}>{p.text}</option>;
                  })}
                </select>
                <button onClick={createComponent} data-state={isLoading && "loading"}>
                  <div className="ld ld-ring ld-spin"></div>
                  Create
                </button>
              </fieldset>
            </div>
          </div>
        </div>
        <div className="box">
          <div className="table-container">
            <table>
              <thead>
                <tr>
                  <th scope="col" className=""></th>
                  <th scope="col" className=""></th>
                  <th scope="col" className=""></th>
                </tr>
              </thead>
              <tbody>
                {layers.map((l) => {
                  return (
                    <tr>
                      <td className="">
                        <span className=" text-ellipsis">{l.name}</span>
                      </td>
                      <td className="">
                        {l.type && <span data-type="badge">{l.type}</span>}
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
