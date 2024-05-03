import { Button, Label, Select } from "flowbite-react";
import { useEffect } from "react";
import { useTemplates } from "../../store/templates";
import { useGeneratePage } from "./state";
import { usePhotoshopStore } from "../../store/photoshop";

export function GeneratePartPage() {
  const { onTemplateSelected, isLoading, onSubmit } = useGeneratePage();
  const { photoshopList: files, initPhotoshopStore } = usePhotoshopStore();
  const { templates, initTemplatesStore, image } = useTemplates((data) => ({
    templates: data.templates,
    get: data.getTemplates,
    generate: data.generateDesign,
    initTemplatesStore: data.initTemplatesStore,
    image: data.designGerated,
  }));

  useEffect(() => {
    initTemplatesStore();
    initPhotoshopStore();
  }, []);

  return (
    <div className="w-screen">
      <div className="max-w-6xl w-full mx-auto">
        <form className="flex flex-col gap-4" onSubmit={onSubmit}>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Arquivo" />
            </div>
            <Select name="photoshop_id">
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
            <Select name="template_id" onChange={onTemplateSelected}>
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
      </div>
      <img src={image} />
    </div>
  );
}
