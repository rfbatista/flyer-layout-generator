import "../../../components/input/input.css";
import "../../../components/label/label.css";
import "../../../components/select/select.css";
import "../../../components/button/button.css";

import { useAdvertiserStore } from "../../../domain/advertiser/store";
import { useClientsStore } from "../../../domain/clients/store";
import { useModal } from "../../../components/modal/store";

export function CreateProjectForm() {
  const { clients } = useClientsStore();
  const { advertisers } = useAdvertiserStore();
  const { close } = useModal();
  return (
    <form className="stack">
      <fieldset>
        <label htmlFor="name">Project name</label>
        <input data-size="md" name="name" value={""} />
      </fieldset>
      <fieldset data-type="select">
        <label htmlFor="os">Client</label>
        <span className="arrow" />
        <select id="os" name="os">
          {clients.map((c) => {
            return <option value={c.id}>{c.name}</option>;
          })}
        </select>
      </fieldset>
      <fieldset data-type="select">
        <label htmlFor="os">Advertiser</label>
        <span className="arrow" />
        <select id="os" name="os">
          {advertisers.map((c) => {
            return <option value={c.id}>{c.name}</option>;
          })}
        </select>
      </fieldset>
      <div className="cluster">
        <div>
          <div>
            <button type="button">Create</button>
          </div>
          <div>
            <button data-type="outline" type="button" onClick={close}>
              Close
            </button>
          </div>
        </div>
      </div>
    </form>
  );
}
