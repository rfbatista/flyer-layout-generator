import { apiClient } from "../../../infrastructure/api";

type Response = {};

export function addNewPropertyToAssetAPI(assetID: number, text: string) {
  return apiClient.post(`/v1/assets/${assetID}`, {
    text,
  });
}
