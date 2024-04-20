import { useEffect, useRef, useState } from "react";
import { fabric } from "fabric";
import { useTemplates } from "../store/templates";
import { Button, Label, Select } from "flowbite-react";
import { usePhotoshopFiles } from "../store/files";

export function GeneratePartPage() {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const fabricRef = useRef<Canvas | null>(null);
  const [made, setMade] = useState(0);
  const [curInterval, setCurInterval] = useState(null);
  const { files, isLoading, images, clear, byRequest } = usePhotoshopFiles(
    (d) => ({
      files: d.files,
      isLoading: d.isLoading,
      images: d.designs,
      clear: d.clearDesigns,
      byRequest: d.getDesignsByRequest,
    }),
  );
  const { templates, get, generate } = useTemplates((data) => ({
    templates: data.templates,
    get: data.getTemplates,
    generate: data.generateDesign,
  }));
  const [data, setData] = useState<number>(templates[0]?.id);

  const onOptionChangeHandler = (event: any) => {
    setData(event.target.value);
    const template = templates.find((e) => e.id === Number(event.target.value));
    console.log(template);
    if (!template) return;

    fabricRef.current?.setHeight(template.height);
    fabricRef.current?.setWidth(template.width);
    fabricRef.current.clear();
    console.log(template.positions);
    for (const pos of template.positions) {
      const rect = new fabric.Rect({
        top: pos.yi,
        left: pos.xi,
        width: pos.width,
        height: pos.height,
        fill: "red",
      });
      fabricRef.current.add(rect);
    }
  };

  useEffect(() => {
    clear();
    get().then((d) => {
      setData(d[0]?.id);
    });
  }, []);

  useEffect(() => {
    if (made === 0) return;
    const intervalId = setInterval(() => byRequest(made), 1000);
    setCurInterval(intervalId)
    return () => {
      clearInterval(intervalId);
    };
  }, [made]);

  const onSubmit = (e: any) => {
    e.preventDefault();
    if(curInterval) clearInterval(curInterval)
    const form = e.target;
    const formData = new FormData(form);
    const formJson = Object.fromEntries(formData.entries());
    const templateChosed = templates.find(
      (t) => t.id === Number(formJson.template),
    );
    const fileChosed = files.find((f) => f.id === Number(formJson.file));
    if (!templateChosed || !fileChosed) return;
    generate(templateChosed, fileChosed).then((d) => {
      setMade(d.id);
    });
  };

  useEffect(() => {
    const template = templates.find((e) => e.id === data);
    if (!template) return;
    const initFabric = () => {
      fabricRef.current = new fabric.Canvas(canvasRef.current);
      fabricRef.current?.setHeight(template.height);
      fabricRef.current?.setWidth(template.width);
      fabricRef.current.skipTargetFind = true;
    };
    initFabric();
    for (const pos of template.positions) {
      const rect = new fabric.Rect({
        top: pos.yi,
        left: pos.xi,
        width: pos.width,
        height: pos.height,
        fill: "red",
      });
      fabricRef.current.add(rect);
    }
  }, [data]);

  return (
    <div className="w-screen">
      <div className="max-w-6xl w-full mx-auto">
        <form className="flex flex-col gap-4" onSubmit={onSubmit}>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Arquivo" />
            </div>
            <Select name="file">
              {files.map((t) => {
                return (
                  <option key={t.id} value={t.id}>
                    {t.filename}
                  </option>
                );
              })}
            </Select>
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Template" />
            </div>
            <Select name="template" onChange={onOptionChangeHandler}>
              {templates.map((t) => {
                return (
                  <option key={t.id} value={t.id}>
                    {t.name}
                  </option>
                );
              })}
            </Select>
          </div>
          <Button type="submit" className="w-full" isProcessing={isLoading}>
            Gerar
          </Button>
        </form>
        <div className="flex justify-center items-center mt-5">
          {data && (
            <canvas ref={canvasRef} className="border border-gray-300" />
          )}
        </div>
      </div>
      <div className="max-w-6xl w-full grid grid-cols-2 md:grid-cols-3 gap-4 mx-auto mb-12">
        {images.map((t) => {
          return (
            <div
              className="bg-clip-padding bg-no-repeat w-[250px] h-[250px]"
              style={{ backgroundImage: `url("${t.src}")` }}
              onClick={() => window.open(t.src, "_blank")}
            ></div>
          );
        })}
      </div>
    </div>
  );
}
