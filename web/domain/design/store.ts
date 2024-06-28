import { create } from "zustand";
import { Design } from "./entities/design";
import { getDesignByIDApi } from "./api/getDesignByIdAPI";
import { listDesignsByProjectID } from "./api/listDesignsByProjectID";
import { processDesignFileAPI } from "./api/process_design_file";
import { uploadDesignAPI } from "./api/upload_design";

type Store = {
  isLoading: boolean;
  designs: Design[];
  activeDesign?: Design;
  createDesign: (f: FormData) => Promise<void>;
  setActiveDesign: (id: number) => Promise<void>;
  listDesigns: (i: number) => Promise<void>;
  processDesignFile: (i: number) => Promise<void>;
  uploadDesign: (f: FormData) => Promise<void>;
};

const useDesignsStore = create<Store>((set) => ({
  designs: [],
  isLoading: false,
  activeDesign: undefined,
  processDesignFile: (id: number) => {
    set({ isLoading: true });
    return processDesignFileAPI(id)
      .then(() => {
        set({ isLoading: false });
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
  createDesign: async (): Promise<void> => {
    set({ isLoading: true });
    return;
  },
  setActiveDesign: (id: number): Promise<void> => {
    set({ isLoading: true });
    return getDesignByIDApi(id)
      .then((p) => {
        set({ isLoading: false, activeDesign: p.data });
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
  listDesigns: async (projectID: number): Promise<void> => {
    set({ isLoading: true });
    return listDesignsByProjectID(projectID)
      .then((res) => {
        set({ isLoading: false, designs: res.data });
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
  uploadDesign: (f: FormData) => {
    set({ isLoading: true });
    return uploadDesignAPI(f)
      .then(() => {
        set({ isLoading: false });
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
}));

export { useDesignsStore };
