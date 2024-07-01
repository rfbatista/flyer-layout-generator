import { create } from "zustand";
import { persist } from "zustand/middleware";
import { defaultPriorities } from "../../../domain/layout/entities/priorities";

export interface Item {
  id: number;
  text: string;
}

type Props = {
  priorities: Item[];
  setPriorities: (i: Item[]) => void;
};

export const useGenerationStore= create<Props>()(
  persist(
    (set) => ({
      priorities: defaultPriorities,
      setPriorities: (t: Item[]) => {
        set({ priorities: t });
      },
    }),
    {
      name: "priorities-storage-v1", // name of the item in the storage (must be unique)
    },
  ),
);
