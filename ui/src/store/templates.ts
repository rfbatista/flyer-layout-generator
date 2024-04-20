import { create } from "zustand";
import { Template } from "../entities/template";
import { api } from "../infra/api";
import { PhotoshopFile } from "../entities/photoshop";

type TemplateStore = {
  templates: Template[];
  isLoading: boolean;
  getTemplates: (limit?: number, skip?: number) => Promise<Template[]>;
  generateDesign: (
    template: Template,
    photoshop: PhotoshopFile,
  ) => Promise<{id: number}>;
  createTemplate: (data: any) => Promise<void>;
};

export const useTemplates = create<TemplateStore>((set) => ({
  templates: [],
  isLoading: false,
  getTemplates: async (limit = 10, skip = 0) => {
    set({ isLoading: true });
    return api
      .get(`/api/v1/template?limit=${limit}&skip=${skip}`)
      .then((res: any) => {
        const data = Template.fromApiList(res);
        set({ templates: data, isLoading: false });
        return data;
      });
  },
  createTemplate: async (data: any) => {
    set({ isLoading: true });
    return api.post(`/api/v1/template`, data).then(() => {
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
        set({ isLoading: false });
        return res.data
      });
  },
}));
