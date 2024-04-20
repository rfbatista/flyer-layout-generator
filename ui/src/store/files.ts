import { PhotoshopImage } from "../entities/image";
import {
  ElementTree,
  PhotoshopElement,
  PhotoshopFile,
} from "../entities/photoshop";
import { api } from "../infra/api";
import { create } from "zustand";
import { getRandomColor } from "../shared/color";
import { getRandomId } from "../shared/uuid";
import { Design } from "../entities/designs";

type PhotoshopStore = {
  execute: Function;
  files: PhotoshopFile[];
  activeFile?: PhotoshopFile;
  activeTree?: ElementTree;
  images: PhotoshopImage[];
  activeElements: PhotoshopElement[];
  elementsSelected: number[];
  isLoading: boolean;
  designs: Design[];
  selectPhotoshop: (psd: PhotoshopFile) => Promise<void>;
  selectElement: (id: number) => void;
  getElements: (psdId: number) => void;
  createComponent: () => Promise<void>;
  removeComponent: () => Promise<void>;
  setBackground: () => Promise<void>;
  getDesigns: (psdId: number) => Promise<void>;
  getDesignsByRequest: (id: number) => Promise<void>;
  clearDesigns: () => void;
};

export const usePhotoshopFiles = create<PhotoshopStore>((set, get) => ({
  files: [],
  images: [],
  designs: [],
  activeFile: undefined,
  activeTree: undefined,
  isLoading: false,
  activeElements: [],
  elementsSelected: [],
  clearDesigns: () => { set({ designs: [] }) },
  execute: async () => {
    set({ isLoading: true });
    api.get("/api/v1/photoshop").then((res) => {
      set({ isLoading: false });
      const data = PhotoshopFile.from_api_list(res);
      set({ files: data });
    });
  },
  selectPhotoshop: async (psd: PhotoshopFile) => {
    set({ activeFile: psd });
    set({ isLoading: true });
    api.get(`/api/v1/photoshop/${psd.id}/images`).then((res) => {
      set({ isLoading: false });
      const data = PhotoshopImage.from_api_list(res);
      set({ images: data });
    });
    get().getElements(psd.id);
  },
  getDesigns: (psdId: number) => {
    set({ isLoading: true });
    return api.get(`/api/v1/photoshop/${psdId}/design`).then((res) => {
      const data = Design.from_api_list(res);
      set({ designs: data, isLoading: false });
      return;
    });
  },
  getDesignsByRequest: (psdId: number) => {
    set({ isLoading: true });
    return api.get(`/api/v1/request/${psdId}/design`).then((res) => {
      const data = Design.from_api_list(res);
      set({ designs: data, isLoading: false });
      return;
    });
  },
  getElements: (psdId: number) => {
    set({ isLoading: true });
    api.get(`/api/v1/photoshop/${psdId}/elements`).then((res) => {
      set({ isLoading: false });
      const data = PhotoshopElement.from_api_list(res);
      const tree = ElementTree.layout(data);
      set({ activeElements: data, activeTree: tree });
    });
  },
  selectElement: (id: number) => {
    if (get().elementsSelected.includes(id)) {
      set((state) => ({
        elementsSelected: state.elementsSelected.filter((i) => i !== id),
      }));
    } else {
      set((state) => ({ elementsSelected: [...state.elementsSelected, id] }));
    }
  },
  createComponent: async () => {
    if (get().elementsSelected.length === 0) return;
    return api
      .post(`/api/v1/component`, {
        elements_id: get().elementsSelected,
        color: getRandomColor(),
        component_id: getRandomId(),
      })
      .then(() => {
        set({ elementsSelected: [] });
        get().getElements(get().activeFile?.id);
      });
  },
  removeComponent: async () => {
    if (get().elementsSelected.length === 0) return;
    return api
      .post(`/api/v1/component/remove`, {
        elements_id: get().elementsSelected,
      })
      .then(() => {
        set({ elementsSelected: [] });
        get().getElements(get().activeFile?.id);
      });
  },
  setBackground: async () => {
    if (get().elementsSelected.length === 0) return;
    return api
      .post(`/api/v1/background`, {
        elements_id: get().elementsSelected,
      })
      .then(() => {
        set({ elementsSelected: [] });
        get().getElements(get().activeFile?.id);
      });
  },
}));
