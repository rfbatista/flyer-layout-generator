import { create } from "zustand";
import { listTemplatesByProjectID } from "./api/list_templates_by_project_id";
import { Template } from "./entities/template";
import { deleteTemplateByID } from "./api/delete_template_by_id";

type Props = {
  isLoading: boolean;
  templates: Template[];
  listTemplatesByProjectID: (id: number) => Promise<void>;
  deleteTemplate: (pid: number, tid: number) => Promise<void>;
};

export const useTemplatesStore = create<Props>((set, get) => ({
  templates: [],
  isLoading: false,
  listTemplatesByProjectID: async (id: number): Promise<void> => {
    try {
      set({ isLoading: true });
      const s = await listTemplatesByProjectID(id);
      set({ isLoading: false, templates: s.data });
    } catch (e) {
      console.error(e);
      set({ isLoading: false });
    }
  },
  deleteTemplate: async (projectID, templateID) => {
    set({ isLoading: true });
    try {
      await deleteTemplateByID(projectID, templateID);
      set({ isLoading: false });
      get()
        .listTemplatesByProjectID(projectID)
        .catch((e) => {
          console.error(e);
        });
      return;
    } catch {
      set({ isLoading: false });
    }
  },
}));
