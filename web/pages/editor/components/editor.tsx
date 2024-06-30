import { Canvas, Point, Rect } from "fabric";
import React, { useEffect, useLayoutEffect, useState } from "react";
import LoadingScreen from "../../../components/loading_screen/loading_screen";
import { useDesignsStore } from "../../../domain/design/store";
import { LayoutManager } from "../../../domain/layout/layout_manager";
import "./editor.css";
import { useEditorStore } from "./editor_store";

export function Editor() {
  const { activeDesign } = useDesignsStore();
  const {
    setEditor,
    editor,
    onMouseDown,
    onMouseMove,
    onMouseUp,
    setLayoutManager,
    addLayer,
    onUnselect,
  } = useEditorStore();
  const canvasRef = React.useRef<HTMLCanvasElement>(null);
  const containerRef = React.useRef<HTMLDivElement>(null);
  const [isLoading, setLoading] = useState(true);
  useEffect(() => {
    if (canvasRef.current != null && containerRef.current && !editor) {
      const editor = new Canvas(canvasRef.current, {
        backgroundColor: "#e2e8f0",
        selection: false,
      });
      setEditor(editor);
      editor.setWidth(containerRef.current.offsetWidth);
      editor.setHeight(containerRef.current.offsetHeight);
      editor.on("mouse:down", onMouseDown);
      editor.on("mouse:move", onMouseMove);
      editor.on("mouse:up", onMouseUp);
      editor.on("before:selection:cleared", onUnselect);
    }
    if (activeDesign && editor && activeDesign.layout) {
      const viewport = new Rect({
        left: 0,
        top: 0,
        width: activeDesign.layout.width,
        height: activeDesign.layout.height,
        selectable: false,
        fill: "white",
        lockMovementX: true,
        lockMovementY: true,
      });
      editor.add(viewport);
      editor.centerObject(viewport);
      editor.zoomToPoint(viewport.getCenterPoint(), 0.5);
      const l = new LayoutManager(editor, addLayer);
      l.setOrigin(new Point(viewport.left, viewport.top));
      if (!activeDesign.layout) return;
      setLoading(true);
      l.drawLayout(activeDesign.layout)
        .then(() => {
          setLoading(false);
          setLayoutManager(l);
        })
        .catch((e) => {
          setLoading(false);
          console.error(e);
        });
    }
  }, [activeDesign]);
  useLayoutEffect(() => {
    function updateSize() {
      if (editor && containerRef.current) {
        editor.setWidth(containerRef.current.offsetWidth);
        editor.setHeight(containerRef.current.offsetHeight);
      }
    }
    window.addEventListener("resize", updateSize);
    updateSize();
    return () => window.removeEventListener("resize", updateSize);
  }, []);
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
