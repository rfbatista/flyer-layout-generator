import { useState } from "react";
import { create } from "zustand";
import { useModal } from "../../../components/modal/store";
import { useAdvertiserStore } from "../../../domain/advertiser/store";
import { useClientsStore } from "../../../domain/clients/store";
import { useProjectsStore } from "../../../domain/projects/store";

type Props = {
  name: string;
  client_id: string;
  advertiser_id: string;
  isLoading: false;
  setName: (s: string) => void;
  setClientID: (s: string) => void;
  setAdvertiserID: (s: string) => void;
};

const useStore = create<Props>((set) => ({
  name: "",
  client_id: "",
  advertiser_id: "",
  isLoading: false,
  setName: (name: string) => {
    set({ name });
  },
  setClientID: (clientID: string) => {
    set({ client_id: clientID });
  },
  setAdvertiserID: (advertiserID: string) => {
    set({ advertiser_id: advertiserID });
  },
}));

const useProjectFormStore = () => {
  const { clients, getClients } = useClientsStore();
  const { advertisers, getAdvertisers } = useAdvertiserStore();
  const { close } = useModal();
  const [isLoading, setIsLoading] = useState(false);
  const { createProject } = useProjectsStore();
  const {
    name,
    setName,
    setClientID,
    client_id,
    setAdvertiserID,
    advertiser_id,
  } = useStore();
  const onSubmit = (e: any) => {
    e.preventDefault();
    const data = new FormData(e.target);
    setIsLoading(true);
    createProject(data)
      .then(() => {
        setIsLoading(false);
      })
      .catch(() => setIsLoading(false));
  };
  return {
    clients,
    getClients,
    advertisers,
    getAdvertisers,
    close,
    name,
    setName,
    setClientID,
    client_id,
    setAdvertiserID,
    advertiser_id,
    onSubmit,
    isLoading,
  };
};

export { useProjectFormStore };
