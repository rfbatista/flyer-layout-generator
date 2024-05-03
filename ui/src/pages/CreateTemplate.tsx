import { Button, Label, Select, TextInput } from "flowbite-react";
import { Scan, MousePointer2, Trash2 } from "lucide-react";
import { useTemplates } from "../store/templates";
import { toast } from "react-toastify";

export function CreateTemplatePage() {
  const { create, isLoading } = useTemplates((s) => ({
    create: s.createTemplate,
    isLoading: s.isLoading,
  }));

  const handleSubmit = (e: any) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    const formJson = Object.fromEntries(formData.entries());
    formJson.type = "distortion";
    formJson.width = Number(formJson.width);
    formJson.height = Number(formJson.height);
    formJson.x = Number(formJson.x);
    formJson.y = Number(formJson.y);
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
            <div className="mb-2 block">
              <Label htmlFor="name" value="Tipo" />
            </div>
            <Select id="types" required>
              <option>Distorção</option>
            </Select>
          </div>
          <div>
            <Label htmlFor="name" value="Grid" />
            <div className="grid grid-cols-2 gap-x-3">
              <div>
                <div className="mb-2 block">
                  <Label htmlFor="name" value="X" />
                </div>
                <TextInput
                  name="x"
                  id="x"
                  type="number"
                  placeholder=""
                  required
                />
              </div>
              <div>
                <div className="mb-2 block">
                  <Label htmlFor="name" value="Y" />
                </div>
                <TextInput
                  id="y"
                  name="y"
                  type="number"
                  placeholder=""
                  required
                />
              </div>
            </div>
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
        </div>
      </div>
    </div>
  );
}
