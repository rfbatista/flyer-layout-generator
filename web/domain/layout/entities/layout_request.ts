import { LayoutRequestJob, LayoutRequestJobProps } from "./layout_request_job";

export type LayoutRequestProps = {
  id: number;
  design_id: number;
  total: number;
  done: number;
  created_at: string;
};

export class LayoutRequest {
  private p: LayoutRequestProps;
  private _jobs: LayoutRequestJob[];
  constructor(p: LayoutRequestProps, j: LayoutRequestJob[] = []) {
    this.p = p;
    this._jobs = j;
  }

  static create(
    p: LayoutRequestProps,
    jobsProps: LayoutRequestJobProps[] = [],
  ) {
    const jobs: LayoutRequestJob[] = [];
    for (const j of jobsProps) {
      jobs.push(LayoutRequestJob.create(j));
    }
    return new LayoutRequest(p, jobs);
  }

  get id() {
    return this.p.id;
  }

  get done() {
    return this.p.done;
  }

  get total() {
    return this.p.total;
  }

  get jobs() {
    return this._jobs;
  }
}
