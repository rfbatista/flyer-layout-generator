import { create } from "zustand";
import { Client } from "./entities/client";
import { apiClient } from "../../infrastructure/api";

type Store = {
  clients: Client[];
  getClients: (page?: number, limit?: number) => Promise<Client[]>;
};

const useClientsStore = create<Store>((set) => ({
  clients: [],
  getClients: (page = 0, limit = 10): Promise<Client[]> => {
    return apiClient
      .get(`/v1/clients?page=${page}&limit=${limit}`)
      .then((r) => {
        const items = r.data.clients;
        const clients: Client[] = [];
        for (const item of items) {
          clients.push(new Client(item));
        }
        set(() => ({ clients }));
        return clients;
      });
  },
}));

export { useClientsStore };
