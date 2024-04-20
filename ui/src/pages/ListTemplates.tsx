import { useEffect } from "react";
import { useTemplates } from "../store/templates";
import { TemplateLayout } from "../components/Template";
import { Plus } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { appConfig } from "../config";
export function ListTemplatePage() {
  const { templates, get } = useTemplates((data) => ({
    templates: data.templates,
    get: data.getTemplates,
  }));

  const navigate = useNavigate();

  useEffect(() => {
    get();
  }, []);

  return (
    <div className="w-screen">
      <div className="max-w-6xl w-full grid grid-cols-2 md:grid-cols-3 gap-4 mx-auto">
        <div
          onClick={() => navigate(appConfig.paths.createTemplates)}
          className="flex items-center justify-center max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700 w-[250px] h-[250px]"
        >
          <Plus size={48} />
        </div>
        {templates.map((t) => {
          return <TemplateLayout key={t.id} template={t} />;
        })}
      </div>
    </div>
  );
}
