import { create } from "zustand";
import { Template } from "../entities/template";
import { api } from "../infra/api";
import { PhotoshopFile } from "../entities/photoshop";
import { createTemplateAPI } from "../api/template/createTemplate";
import { listTemplateAPI } from "../api/template/listTemplate";

type TemplateStore = {
  templates: Template[];
  isLoading: boolean;
  wasInit: boolean;
  designGerated: string;
  initTemplatesStore: () => void;
  getTemplates: (limit?: number, skip?: number) => Promise<Template[]>;
  generateDesign: (
    template: Template,
    photoshop: PhotoshopFile,
  ) => Promise<{ id: number }>;
  createTemplate: (data: any) => Promise<void>;
};

export const useTemplates = create<TemplateStore>((set, get) => ({
  templates: [],
  isLoading: false,
  wasInit: false,
  designGerated: "",
  initTemplatesStore: () => {
    if (get().wasInit) return;
    get().getTemplates();
  },
  getTemplates: async (limit = 10, skip = 0) => {
    set({ isLoading: true });
    return listTemplateAPI(limit, skip).then((t) => {
      set({ templates: t, isLoading: false });
      return t;
    });
  },
  createTemplate: async (data: any) => {
    set({ isLoading: true });
    return createTemplateAPI(data).then(() => {
      set({ isLoading: false });
      return;
    });
  },
  generateDesign: (template: Template, psd: PhotoshopFile) => {
    set({ isLoading: true });
    return api
      .post(`/api/v1/design`, {
        photoshop_id: psd.id,
        templates: [
          {
            id: template.id,
            background_position: {
              xi: 0,
              yi: 0,
            },
          },
        ],
      })
      .then((res: any) => {
        console.log(res.data);
        set({ isLoading: false, designGerated: res.data.image_url });
        return res.data;
      });
  },
}));
