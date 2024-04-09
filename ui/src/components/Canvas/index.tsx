import { useEffect, useRef } from "react";
import { usePhotoshopFiles } from "../../store/files";

export function Canvas() {
  const { images, photoshopfile } = usePhotoshopFiles((data) => ({
    images: data.images,
    photoshopfile: data.activeFile,
  }));
  const canvasRef = useRef(null);
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    if (!images.length) return;
    const image = new Image();
    console.log(images);
    const context = canvas.getContext("2d");
    image.src = images[0].src;
    image.onload = () => {
      context.drawImage(image, 0, 0, canvas.width, canvas.height);
    };
  }, [images]);
  console.log(photoshopfile?.toScale(500, 500));
  return (
    <>
      <canvas
        ref={canvasRef}
        width={photoshopfile?.toScale(500, 500)?.width}
        height={photoshopfile?.toScale(500, 500)?.height}
      ></canvas>
    </>
  );
}
