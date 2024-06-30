import { Canvas } from "fabric";
import React, { useEffect } from "react";
import { useDesignsStore } from "../../../domain/design/store";
import { LayoutManager } from "../../../domain/layout/layout_manager";
import "./editor.css";
import { useEditorStore } from "./editor_store";

export function Editor() {
  const { activeDesign } = useDesignsStore();
  const { setEditor, editor } = useEditorStore();
  const canvasRef = React.useRef<HTMLCanvasElement>(null);
  const containerRef = React.useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (canvasRef.current != null && containerRef.current && !editor) {
      const editor = new Canvas(canvasRef.current);
      setEditor(editor);
      editor.setWidth(containerRef.current.offsetWidth);
      editor.setHeight(containerRef.current.offsetHeight);
    }
    if (activeDesign && editor) {
      const l = new LayoutManager(editor);
      activeDesign.layout && l.drawLayout(activeDesign.layout);
    }
  }, [activeDesign]);
  return (
    <>
      <div className="w-full editor__canvas-container" ref={containerRef}>
        <canvas id="canvas" ref={canvasRef} />
      </div>
    </>
  );
}
