import { Component } from "../../entities/component";
import { api } from "../../infra/api";

export interface ListComponentsByFileIdAPIResult {
  status: string;
  data: Data[];
}

export interface Data {
  id: number;
  photoshop_id: number;
  width: number;
  height: number;
  color: any;
  xi: number;
  xii: number;
  yi: number;
  yii: number;
  created_at: string;
}

export async function listComponentsByFileId(id: number): Promise<Component[]> {
  const res = await api.get<ListComponentsByFileIdAPIResult>(
    `/api/v1/file/${id}/components`,
  );
  const raw = res.data.data;
  const entities = [];
  for (const c of raw) {
    entities.push(new Component(c));
  }
  return entities;
}
