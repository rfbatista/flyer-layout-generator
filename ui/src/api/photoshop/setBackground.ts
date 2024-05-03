import { api } from "../../infra/api";

export async function setBackgroundAPI(
  photoshopId: number,
  ids: number[],
): Promise<any> {
  await api.post(`/api/v1/photoshop/${photoshopId}/background`, {
    elements: ids,
  });
}
