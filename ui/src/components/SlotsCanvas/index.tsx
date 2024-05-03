import { useEffect, useRef, useState } from "react";
import { fabric } from "fabric";
import { Canvas } from "fabric/fabric-impl";

type Rectangle = {
  xi: number;
  yi: number;
  width: number;
  height: number;
};

const SlotsCanvas = () => {
  const [canvasState, setCanvasState] = useState({
    xi: 0,
    yi: 0,
    width: 400,
    height: 400,
    isDragOn: false,
    rectangles: [],
  });
  const fabricRef = useRef<Canvas | null>(null);

  const handleWidthInputChange = (e: any) => {
    if (!fabricRef.current) return;
    const value = e.target.value;
    fabricRef.current?.setWidth(value);
    setCanvasState((s) => ({ ...s, width: e.target.value }));
  };

  const handleHeightInputChange = (e: any) => {
    if (!fabricRef.current) return;
    const value = e.target.value;
    fabricRef.current?.setHeight(value);
    setCanvasState((s) => ({ ...s, height: e.target.value }));
  };

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

  const removeSelected = () => {
    if (!fabricRef.current) return;
    const activeObj = fabricRef.current.getActiveObject();
    if (!activeObj) return;
    fabricRef.current?.remove(activeObj);
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

  const handleSubmit = (e: any) => {
    e.preventDefault();
    const objects = fabricRef.current?.getObjects();
    if (!objects) return;
    const positions = [];
    for (const object of objects) {
      positions.push({
        xi: Math.floor(object.left),
        yi: Math.floor(object.top),
        width: Math.floor(object.width),
        height: Math.floor(object.height),
      });
    }
  };

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
