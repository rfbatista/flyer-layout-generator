import { create } from "zustand";
import { Template } from "../../entities/template";
import { generateDesignAPI } from "../../api/design/generate";
import { DesignImage } from "../../entities/image";

type Props = {
  isLoading: boolean;
  onTemplateSelected: () => void;
  templates: Template[];
  onSubmit: (e: any) => void;
  imageGerated?: DesignImage;
};

export const useGeneratePage = create<Props>((set) => ({
  isLoading: false,
  templates: [],
  onTemplateSelected: () => { },
  onSubmit: (e: any) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    set({ isLoading: true, imageGerated: undefined });
    generateDesignAPI(formData)
      .then((d) => {
        set({ isLoading: false, imageGerated: DesignImage.create(d) });
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
}));
