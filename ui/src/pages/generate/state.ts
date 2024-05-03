import { create } from "zustand";
import { Template } from "../../entities/template";
import { generateDesignAPI } from "../../api/design/generate";

type Props = {
  isLoading: boolean;
  onTemplateSelected: () => void;
  templates: Template[];
  onSubmit: (e: any) => void;
};

export const useGeneratePage = create<Props>(() => ({
  isLoading: false,
  templates: [],
  onTemplateSelected: () => {},
  onSubmit: (e: any) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    generateDesignAPI(formData)
  },
}));
