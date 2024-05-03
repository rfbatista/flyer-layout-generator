import { create } from "zustand";
import { ElementTree, PhotoshopFile } from "../entities/photoshop";
import { getPhotoshopListAPI } from "../api/photoshop/listPhotoshopFile";
import { getPhotoshopElementsAPI } from "../api/photoshop/getPhotoshopElements";
import { toast } from "react-toastify";
import { removeComponentAPI } from "../api/photoshop/removeComponent";
import { setBackgroundAPI } from "../api/photoshop/setBackground";
import { createComponentAPI } from "../api/photoshop/createComponent";
import { listComponentsByFileId } from "../api/components/listComponentsByFileId";
import { Component } from "../entities/component";

type Props = {
  isLoading: boolean;
  initPhotoshopStore: () => void;
  wasInit: boolean;
  tree?: ElementTree;
  photoshopList: PhotoshopFile[];
  activePhotoshop?: PhotoshopFile;
  elementsSelected: number[];
  components: Component[];
  selectElement: (id: number) => void;
  listPhotoshop: (limit?: number, skip?: number) => void;
  selectPhotoshop: (psd: PhotoshopFile) => void;
  getPhotoshopElements: (id: number) => Promise<ElementTree>;
  onCreateComponent: () => void;
  onRemoveComponent: () => void;
  onSetBackground: () => void;
  getComponents: (id: number) => void;
};

export const usePhotoshopStore = create<Props>((set, get) => ({
  isLoading: false,
  components: [],
  wasInit: false,
  tree: undefined,
  elementsSelected: [],
  photoshopList: [],
  activePhotoshop: undefined,
  getComponents: (id: number) => {
    set({ isLoading: true });
    listComponentsByFileId(id)
      .then((d) => {
        set({ isLoading: false, components: d });
      })
      .catch(() => {
        set({ isLoading: false });
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
  initPhotoshopStore: () => {
    get().listPhotoshop();
  },
  selectPhotoshop: (psd: PhotoshopFile) => {
    set({ activePhotoshop: psd });
    get().getPhotoshopElements(psd.id);
    get().getComponents(psd.id);
  },
  getPhotoshopElements: async (id: number) => {
    set({ isLoading: true });
    const t = await getPhotoshopElementsAPI(id);
    set({ tree: t, isLoading: false });
    return t;
  },
  listPhotoshop: (_limit, _skip) => {
    set({ isLoading: true });
    getPhotoshopListAPI().then((data) => {
      set({ photoshopList: data, isLoading: false });
    });
  },
  onRemoveComponent: () => {
    if (get().elementsSelected.length === 0) return;
    removeComponentAPI(get().activePhotoshop?.id || 0, get().elementsSelected)
      .then(() => {
        toast.success("Componente atualizado");
        const id = get().activePhotoshop?.id;
        id && get().getPhotoshopElements(id);
        set({ elementsSelected: [] });
      })
      .catch(() => {
        set({ elementsSelected: [] });
        toast.error("Falha ao atualizar componente");
      });
  },
  onSetBackground: () => {
    if (get().elementsSelected.length === 0) return;
    setBackgroundAPI(get().activePhotoshop?.id || 0, get().elementsSelected)
      .then(() => {
        toast.success("Componente atualizado");
        const id = get().activePhotoshop?.id;
        set({ elementsSelected: [] });
        id && get().getPhotoshopElements(id);
      })
      .catch(() => {
        set({ elementsSelected: [] });
        toast.error("Falha ao atualizar componente");
      });
  },
  onCreateComponent: () => {
    if (get().elementsSelected.length === 0) return;
    createComponentAPI(get().activePhotoshop?.id || 0, get().elementsSelected)
      .then(() => {
        toast.success("Componente criado");
        const id = get().activePhotoshop?.id;
        set({ elementsSelected: [] });
        id && get().getPhotoshopElements(id);
      })
      .catch(() => {
        set({ elementsSelected: [] });
        toast.error("Falha ao criar componente");
      });
  },
}));
