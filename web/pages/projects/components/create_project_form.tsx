import { useEffect } from "react";
import { useProjectFormStore } from "./create_project_form_store";

export function CreateProjectForm() {
  const {
    name,
    setName,
    setClientID,
    client_id,
    setAdvertiserID,
    advertiser_id,
    clients,
    getClients,
    advertisers,
    getAdvertisers,
    close,
    isLoading,
    onSubmit,
  } = useProjectFormStore();

  useEffect(() => {
    getAdvertisers();
    getClients();
  }, []);

  return (
    <form className="stack" onSubmit={onSubmit}>
      <fieldset>
        <label htmlFor="name">Project name</label>
        <input
          data-size="md"
          name="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      </fieldset>
      <fieldset data-type="select">
        <label htmlFor="os">Client</label>
        <span className="arrow" />
        <select
          id="os"
          name="client_id"
          value={client_id}
          onChange={(e) => setClientID(e.target.value)}
        >
          {clients.map((c) => {
            return <option value={c.id}>{c.name}</option>;
          })}
        </select>
      </fieldset>
      <fieldset data-type="select">
        <label htmlFor="os">Advertiser</label>
        <span className="arrow" />
        <select
          id="os"
          name="advertiser_id"
          value={advertiser_id}
          onChange={(e) => setAdvertiserID(e.target.value)}
        >
          {advertisers.map((c) => {
            return <option value={c.id}>{c.name}</option>;
          })}
        </select>
      </fieldset>
      <div className="cluster">
        <div>
          <div>
            <button type="submit" data-state={isLoading && "loading"}>
              <span data-type="spinner" />
              Create
            </button>
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
