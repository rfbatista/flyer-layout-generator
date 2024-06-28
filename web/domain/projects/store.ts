import { create } from "zustand";
import { Project } from "./entities/project";
import { apiClient } from "../../infrastructure/api";
import { listProjectsAPI } from "./api/listprojects";
import { getProjectByIdAPI } from "./api/get_project_by_id";

type Store = {
  isLoading: boolean;
  projects: Project[];
  activeProject?: Project;
  createProject: (f: FormData) => Promise<void>;
  setActiveProject: (id: number) => Promise<void>;
  listProjects: (p?: number, l?: number) => Promise<void>;
};

const useProjectsStore = create<Store>((set, get) => ({
  projects: [],
  isLoading: false,
  activeProject: undefined,
  createProject: async (data): Promise<void> => {
    set({ isLoading: true });
    return apiClient
      .post("/v1/project", data, {
        headers: { "content-type": "multipart/form-data" },
      })
      .then(() => {
        set({ isLoading: false });
        get().listProjects();
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
  setActiveProject: (id: number): Promise<void> => {
    set({ isLoading: true });
    return getProjectByIdAPI(id)
      .then((p) => {
        set({ activeProject: p.data });
        set({ isLoading: false });
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
  listProjects: async (page = 0, limit = 10): Promise<void> => {
    set({ isLoading: true });
    return listProjectsAPI(page, limit)
      .then((res) => {
        set({ isLoading: false, projects: res.data });
      })
      .catch(() => {
        set({ isLoading: false });
      });
  },
}));

export { useProjectsStore };
