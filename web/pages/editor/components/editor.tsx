import { Canvas } from "fabric";
import React, { useEffect } from "react";
import LoadingScreen from "../../../components/loading_screen/loading_screen";
import { useLayoutStore } from "../../../domain/layout/layout_store";
import "./editor.css";
import { useEditorStore } from "./editor_store";

export function Editor() {
  const { activeLayout } = useLayoutStore();
  const { isLoading, editor, init: initEditor } = useEditorStore();
  const canvasRef = React.useRef<HTMLCanvasElement>(null);
  const containerRef = React.useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (!activeLayout) return;
    const init = async () => {
      if (canvasRef.current != null && containerRef.current && !editor) {
        const editor = new Canvas(canvasRef.current, {
          backgroundColor: "#e2e8f0",
          selection: false,
        });
        initEditor(
          editor,
          activeLayout,
          containerRef.current.offsetWidth,
          containerRef.current.offsetHeight,
        );
      }
    };
    init();
  }, [activeLayout]);

  return (
    <>
      <LoadingScreen isLoading={isLoading}>
        <div className="w-full editor__canvas-container" ref={containerRef}>
          <canvas id="canvas" ref={canvasRef} />
        </div>
      </LoadingScreen>
    </>
  );
}
