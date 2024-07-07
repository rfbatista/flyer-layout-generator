import { apiClient } from "../../../infrastructure/api";
import { DesignAsset, DesignAssetProps } from "../entities/design_asset";

export interface Response {
  data: DesignAssetProps[];
}

export async function getProjectDesignAssets(projectId: number): Promise<DesignAsset[]> {
  const r = await apiClient
    .get<Response>(`/v1/project/${projectId}/assets`);
  return r.data.data.map((d) => DesignAsset.create(d));
}
