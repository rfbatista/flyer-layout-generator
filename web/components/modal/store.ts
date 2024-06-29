import { ReactNode } from "react";
import { create } from "zustand";

type Store = {
  isOpen: boolean;
  open: () => void;
  close: () => void;
  toggle: () => void;
  ch?: ReactNode;
  setCh: (ch: ReactNode) => void;
  title: string;
  setTitle: (t: string) => void;
};

const useModal = create<Store>((set, get) => ({
  isOpen: false,
  ch: undefined,
  title: "",
  setTitle: (t: string) => {
    set({ title: t });
  },
  setCh: (c: ReactNode) => {
    set({ ch: c });
  },
  toggle: () => {
    if (!get().isOpen) {
      get().open();
    } else {
      get().close();
    }
  },
  open: () => {
    set(() => ({ isOpen: true }));
  },
  close: () => {
    set(() => ({ isOpen: false }));
  },
}));

export { useModal };
