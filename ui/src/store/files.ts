import { PhotoshopImage } from "../entities/image";
import { PhotoshopFile } from "../entities/photoshop";
import { api } from "../infra/api";
import { create } from "zustand";

type PhotoshopStore = {
  execute: Function;
  files: PhotoshopFile[];
  activeFile?: PhotoshopFile;
  images: PhotoshopImage[];
  getImageFromPsd: (psd: PhotoshopFile) => Promise<void>;
};

export const usePhotoshopFiles = create<PhotoshopStore>((set) => ({
  files: [],
  images: [],
  activeFile: undefined,
  execute: async () => {
    api.get("/api/v1/photoshop").then((res) => {
      const data = PhotoshopFile.from_api_list(res);
      set({ files: data });
    });
  },
  getImageFromPsd: async (psd: PhotoshopFile) => {
    set({ activeFile: psd });
    api.get(`/api/v1/photoshop/${psd.id}/images`).then((res) => {
      const data = PhotoshopImage.from_api_list(res);
      console.log(data);
      set({ images: data });
    });
  },
}));
