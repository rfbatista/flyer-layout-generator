import { create } from "zustand";

type Store = {
  isOpen: boolean;
  open: () => void;
  close: () => void;
  toggle: () => void;
};

const useModal = create<Store>((set, get) => ({
  isOpen: false,
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
