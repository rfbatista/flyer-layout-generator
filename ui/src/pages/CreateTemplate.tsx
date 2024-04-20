import { useEffect, useRef, useState } from "react";
import { Button, Label, TextInput } from "flowbite-react";
import { fabric } from "fabric";
import { Scan, MousePointer2, Trash2 } from "lucide-react";
import { useTemplates } from "../store/templates";
import { toast } from "react-toastify";

type Rectangle = {
  xi: number;
  yi: number;
  width: number;
  height: number;
};

export function CreateTemplatePage() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const { create, isLoading } = useTemplates((s) => ({
    create: s.createTemplate,
    isLoading: s.isLoading,
  }));
  const [canvasState, setCanvasState] = useState({
    xi: 0,
    yi: 0,
    width: 400,
    height: 400,
    isDragOn: false,
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
    const form = e.target;
    const formData = new FormData(form);
    const formJson = Object.fromEntries(formData.entries());
    formJson["positions"] = positions;
    create(formJson)
      .then(() => toast.success("Template criado"))
      .catch(() => toast.error("Falha ao criar template"));
  };

  return (
    <div className="w-screen">
      <div className="max-w-6xl w-full mx-auto">
        <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Nome" />
            </div>
            <TextInput
              name="name"
              id="name"
              type="text"
              placeholder=""
              required
            />
          </div>
          <div>
            <Label htmlFor="name" value="Dimensões" />
            <div className="grid grid-cols-2 gap-x-3">
              <div>
                <div className="mb-2 block">
                  <Label htmlFor="name" value="Largura" />
                </div>
                <TextInput
                  name="width"
                  id="width"
                  type="number"
                  placeholder=""
                  required
                  onChange={handleWidthInputChange}
                  value={canvasState.width}
                />
              </div>
              <div>
                <div className="mb-2 block">
                  <Label htmlFor="name" value="Altura" />
                </div>
                <TextInput
                  id="height"
                  name="height"
                  type="number"
                  placeholder=""
                  required
                  onChange={handleHeightInputChange}
                  value={canvasState.height}
                />
              </div>
            </div>
          </div>
          <Button className="w-full" type="submit" isProcessing={isLoading}>
            Cadastrar
          </Button>
        </form>
        <hr className="my-5" />
        <div className="flex justify-center items-center mt-5">
          <div className="my-2 flex justify-center items-center">
            <div className="flex">
              <button
                type="button"
                className={`nav-left`}
                onClick={() => addRectangle()}
              >
                <Scan size={"20px"} className="mr-2" />
              </button>
              <button type="button" className="nav-mid">
                <MousePointer2 size={"20px"} className="mr-2" />
              </button>
              <button
                type="button"
                className="nav-right"
                onClick={() => removeSelected()}
              >
                <Trash2 size={"20px"} className="mr-2" />
              </button>
            </div>
          </div>
        </div>

        <div className="flex justify-center items-center mt-5">
          <canvas
            ref={canvasRef}
            className="border border-gray-300 w-fit"
            style={{ border: "1px solid black" }}
          />
        </div>
      </div>
    </div>
  );
}
