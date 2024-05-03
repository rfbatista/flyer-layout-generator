import { api } from "../../infra/api";
import { getRandomColor } from "../../shared/color";

export async function createComponentAPI(
  photoshopId: number,
  ids: number[],
): Promise<any> {
  await api.post(`/api/v1/photoshop/${photoshopId}/component`, {
    type: "logotipo_marca",
    elements_id: ids,
    color: `#${getRandomColor()}`,
  });
}
