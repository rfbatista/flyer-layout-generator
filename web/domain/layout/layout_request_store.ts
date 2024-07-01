import { create } from "zustand";
import { LayoutRequest } from "./entities/layout_request";
import { LayoutRequestJob } from "./entities/layout_request_job";
import { getLastLayoutRequestAPI } from "../layout/api/get_last_layout_request";

type Props = {
  request?: LayoutRequest;
  jobs: LayoutRequestJob[];
  getJobs: (designId: number) => Promise<void>;
};

export const useLayoutRequestStore = create<Props>((set) => ({
  request: undefined,
  jobs: [],
  getJobs: (designId: number): Promise<void> => {
    return getLastLayoutRequestAPI(designId)
      .then((s) => {
        set({ request: s, jobs: s.jobs });
      })
      .catch((e) => {
        console.error(e);
      });
  },
}));
