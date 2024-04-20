import { useEffect, useRef } from "react";
import { Template } from "../../entities/template";

type Props = {
  template: Template;
};

export const TemplateLayout: React.FC<Props> = (props) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext("2d");
    if (!ctx) return;
    // Clear the canvas
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    console.log(props.template.widthRatio(250));
    for (const pos of props.template.positions) {
      console.log(pos);
      pos.resize(
        props.template.widthRatio(250),
        props.template.heightRatio(250),
      );
      console.log(pos);
      ctx.strokeRect(pos.xi, pos.yi, pos.width, pos.height);
    }
  }, []);
  return (
    <div
      className={`border border-gray-300 w-[${props.template.toScale(250, 250).width}px] h-[${props.template.toScale(250, 250).height}px]`}
    >
      <canvas
        ref={canvasRef}
        width={props.template.toScale(250, 250).width}
        height={props.template.toScale(250, 250).height}
      ></canvas>
    </div>
  );
};
