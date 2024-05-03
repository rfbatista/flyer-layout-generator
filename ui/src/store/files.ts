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
import { getPhotoshopListAPI } from "../api/photoshop/listPhotoshopFile";

type PhotoshopStore = {
  init: Function;
  isInitiated: boolean;
  files: PhotoshopFile[];
  activeFile?: PhotoshopFile;
  activeTree?: ElementTree;
  images: PhotoshopImage[];
  activeElements: PhotoshopElement[];
  mainBoardSize: { width: number; height: number };
  setMainBoardSize: (d: { width: number; height: number }) => void;
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
  isInitiated: false,
  images: [],
  designs: [],
  activeFile: undefined,
  activeTree: undefined,
  isLoading: false,
  mainBoardSize: { height: 0, width: 0 },
  activeElements: [],
  elementsSelected: [],
  clearDesigns: () => {
    set({ designs: [] });
  },
  init: async () => {
    if (get().isInitiated) return;
    set({ isInitiated: true });
    set({ isLoading: true });
    getPhotoshopListAPI().then((data) => {
      set({ files: data, isLoading: false });
    });
  },
  setMainBoardSize: (d: { height: number; width: number }) => {
    set({ mainBoardSize: d });
  },
  selectPhotoshop: async (psd: PhotoshopFile) => {
    set({ activeFile: psd });
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
