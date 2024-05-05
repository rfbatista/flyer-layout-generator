import { useEffect, useRef, useState } from "react";
import { fabric } from "fabric";
import { Canvas } from "fabric/fabric-impl";

const SlotsCanvas = () => {
  const [canvasState] = useState({
    xi: 0,
    yi: 0,
    width: 400,
    height: 400, isDragOn: false,
    rectangles: [],
  });
  const fabricRef = useRef<Canvas | null>(null);

  const addRectangle = () => {
    if (!fabricRef.current) return;
    const rect = new fabric.Rect({
      top: 50,
      left: 50,
      width: 50,
      height: 50,
      fill: "red",
    });
    fabricRef.current.add(rect);
  };

  useEffect(() => {
    const initFabric = () => {
      fabricRef.current = new fabric.Canvas(canvasRef.current);
      fabricRef.current?.setDimensions({
        width: canvasState.width,
        height: canvasState.height,
      });
    };

    const disposeFabric = () => {
      if (!fabricRef.current) return;
      fabricRef.current.dispose();
    };

    initFabric();
    addRectangle();

    return () => {
      disposeFabric();
    };
  }, []);
  const canvasRef = useRef<HTMLCanvasElement>(null);

  return (
    <>
      <canvas
        ref={canvasRef}
        className="border border-gray-300 w-fit"
        style={{ border: "1px solid black" }}
      />
    </>
  );
};

export default SlotsCanvas;
