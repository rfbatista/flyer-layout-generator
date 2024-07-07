import { create } from "zustand";
import { Layout } from "./entities/layout";
import { getLayoutByIdAPI } from "./api/get_layout_by_id";

type Props = {
  activeLayout?: Layout;
  setActiveLayout: (id: number) => Promise<void>;
};

export const useLayoutStore = create<Props>((set) => ({
  activeLayout: undefined,
  setActiveLayout: async (id: number) => {
    try {
      const l = await getLayoutByIdAPI(id);
      set({ activeLayout: l });
      return;
    } catch (e) {
      console.error(e);
    }
  },
}));
