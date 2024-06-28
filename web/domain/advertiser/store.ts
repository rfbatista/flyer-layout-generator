import { create } from "zustand";
import { Advertiser } from "./entities/advertiser";
import { apiClient } from "../../infrastructure/api";

type Store = {
  advertisers: Advertiser[];
  getAdvertisers: (page?: number, limit?: number) => Promise<Advertiser[]>;
};

const useAdvertiserStore = create<Store>((set) => ({
  advertisers: [],
  getAdvertisers: (page = 0, limit = 10): Promise<Advertiser[]> => {
    return apiClient
      .get(`/v1/advertisers?page=${page}&limit=${limit}`)
      .then((r) => {
        const items = r.data.advertisers;
        const advertisers: Advertiser[] = [];
        for (const item of items) {
          advertisers.push(new Advertiser(item));
        }
        set(() => ({ advertisers }));
        return advertisers;
      });
  },
}));

export { useAdvertiserStore };
