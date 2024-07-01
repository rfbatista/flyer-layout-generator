import { apiClient } from "../../../infrastructure/api";
import { LayoutRequestJobProps } from "../../layout/entities/layout_request_job";
import { LayoutRequest } from "../entities/layout_request";

type Response = {
  data: {
    id: number;
    design_id: number;
    created_at: string;
    total: number;
    done: number;
    jobs: Array<LayoutRequestJobProps>;
  };
};

export function getLastLayoutRequestAPI(designID: number) {
  return apiClient
    .get<Response>(`/v1/project/design/${designID}/last_request`)
    .then((r) => {
      return LayoutRequest.create(r.data.data, r.data.data.jobs);
    });
}
