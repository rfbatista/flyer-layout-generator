import { api } from "../../infra/api";

export async function removeComponentAPI(
  photoshopId: number,
  ids: number[],
): Promise<any> {
  await api.post(`/api/v1/photoshop/${photoshopId}/components/remove`, {
    elements: ids,
  });
}
