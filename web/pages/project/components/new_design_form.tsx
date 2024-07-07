import { useModal } from "../../../components/modal/store";
import { useNewDesignFormStore } from "./new_design_form_store";

export function NewDesignForm() {
  const { close } = useModal();
  const { onSubmit, name, setName, isLoading, setFile, activeProject } =
    useNewDesignFormStore();
  return (
    <form className="stack" onSubmit={onSubmit}>
      <input
        name="project_id"
        className="hide"
        value={activeProject && activeProject.id}
      />
      <fieldset>
        <label htmlFor="name">Design name</label>
        <input
          data-size="md"
          name="filename"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      </fieldset>
      <fieldset>
        <label htmlFor="file">File</label>
        <input data-size="md" name="file" onChange={setFile} type="file" />
      </fieldset>
      <div className="cluster">
        <div>
          <div>
            <button type="submit" data-state={isLoading && "loading"}>
              <div className="ld ld-ring ld-spin"></div>
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
