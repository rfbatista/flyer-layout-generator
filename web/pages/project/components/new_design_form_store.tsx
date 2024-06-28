import { useState } from "react";
import { create } from "zustand";
import { useModal } from "../../../components/modal/store";
import { useAdvertiserStore } from "../../../domain/advertiser/store";
import { useClientsStore } from "../../../domain/clients/store";
import { useDesignsStore } from "../../../domain/design/store";
import { useProjectsStore } from "../../../domain/projects/store";
// import { useProjectsStore } from "../../../domain/projects/store";

type Props = {
  name: string;
  client_id: string;
  advertiser_id: string;
  isLoading: false;
  file?: any;
  setName: (s: string) => void;
  setClientID: (s: string) => void;
  setAdvertiserID: (s: string) => void;
  setFile: (s: any) => void;
};

const useStore = create<Props>((set) => ({
  name: "",
  client_id: "",
  advertiser_id: "",
  isLoading: false,
  file: undefined,
  setName: (name: string) => {
    set({ name });
  },
  setClientID: (clientID: string) => {
    set({ client_id: clientID });
  },
  setAdvertiserID: (advertiserID: string) => {
    set({ advertiser_id: advertiserID });
  },
  setFile: (event: any) => {
    const f = event.target.files[0];
    set({ file: f });
  },
}));

const useNewDesignFormStore = () => {
  const { clients, getClients } = useClientsStore();
  const { advertisers, getAdvertisers } = useAdvertiserStore();
  const { close } = useModal();
  const [isLoading, setIsLoading] = useState(false);
  const { uploadDesign, listDesigns } = useDesignsStore();
  // const { createProject } = useProjectsStore();

  const { activeProject } = useProjectsStore();
  const {
    name,
    setName,
    setClientID,
    client_id,
    setAdvertiserID,
    advertiser_id,
    setFile,
  } = useStore();
  const onSubmit = (e: any) => {
    e.preventDefault();
    const data = new FormData(e.target);
    setIsLoading(true);
    uploadDesign(data)
      .then(() => {
        setIsLoading(false);
        activeProject && listDesigns(activeProject.id);
      })
      .catch(() => {
        setIsLoading(false);
      });
  };
  return {
    clients,
    getClients,
    advertisers,
    getAdvertisers,
    activeProject,
    close,
    name,
    setName,
    setClientID,
    client_id,
    setAdvertiserID,
    advertiser_id,
    onSubmit,
    isLoading,
    setFile,
  };
};

export { useNewDesignFormStore };
