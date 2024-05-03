import { ElementTree, PhotoshopElement } from "../../entities/photoshop";
import { api } from "../../infra/api";

export async function getPhotoshopElementsAPI(id: number): Promise<ElementTree> {
  const res = await api.get(`/api/v1/photoshop/${id}/elements?limit=100`);
  const data = PhotoshopElement.from_api_list(res);
  const tree = ElementTree.layout(data);
  return tree;
}
