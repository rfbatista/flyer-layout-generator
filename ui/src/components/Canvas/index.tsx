import { useEffect, useRef, useState } from "react";
import { usePhotoshopFiles } from "../../store/files";

export function Canvas() {
  const widthRef = useRef<any>(null);
  const { images, photoshopfile } = usePhotoshopFiles((data) => ({
    images: data.images,
    photoshopfile: data.activeFile,
  }));

  const [width, setWidth] = useState(500);
  const canvasRef = useRef(null);
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    if (!images.length) return;
    const image = new Image();
    const context = canvas.getContext("2d");
    image.src = images[0].src;
    image.onload = () => {
      context.drawImage(image, 0, 0, canvas.width, canvas.height);
    };
  }, [images]);

  useEffect(() => {
    const newWidth = widthRef.current ? widthRef.current.offsetWidth : 500;
    setWidth(newWidth);
  }, [widthRef.current]);

  console.log("ok");

  return (
    <>
      <div className="flex justify-center items-center mt-5 px-7" ref={widthRef}>
        <canvas
          ref={canvasRef}
          width={photoshopfile?.toScale(width, 500)?.width}
          height={photoshopfile?.toScale(width, 500)?.height}
        ></canvas>
      </div>
    </>
  );
}
