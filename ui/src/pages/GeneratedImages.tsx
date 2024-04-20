import { useEffect } from "react";
import { TemplateLayout } from "../components/Template";
import { usePhotoshopFiles } from "../store/files";

export const GeneratedImagesPage = () => {
  const { images, get } = usePhotoshopFiles((data) => ({
    images: data.designs,
    get: data.getDesigns,
  }));

  useEffect(() => {
    get(1);
  }, []);

  return (
    <div className="w-screen">
      <div>
        <form></form>
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
};
