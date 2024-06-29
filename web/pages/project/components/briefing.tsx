import { useEffect, useState } from "react";
import "../../../components/textarea/textarea.css";
import { updateProjectByIdAPI } from "../../../domain/projects/api/update_project_by_id";
import { useProjectsStore } from "../../../domain/projects/store";

export function Briefing() {
  const { activeProject, setActiveProject } = useProjectsStore();
  const [briefing, setBriefing] = useState(
    activeProject ? activeProject.briefing : "",
  );
  const [useAi, setUseAi] = useState<boolean>(
    activeProject ? activeProject.useAI : false,
  );
  const [isLoading, setIsLoading] = useState(false);
  useEffect(() => {
    if (activeProject) {
      setBriefing(activeProject.briefing);
      setUseAi(activeProject.useAI);
    }
  }, [activeProject]);
  const onSubmit = (e: any) => {
    e.preventDefault();
    const data = new FormData();
    data.set("briefing", briefing);
    data.set("use_ai", String(useAi));
    setIsLoading(true);
    activeProject &&
      updateProjectByIdAPI(activeProject?.id, data)
        .then(() => {
          setIsLoading(false);
          activeProject && setActiveProject(activeProject.id);
        })
        .catch(() => {
          setIsLoading(false);
        });
  };
  return (
    <form onSubmit={onSubmit}>
      <div className="stack">
        <div>
          <div className="box">
            <fieldset>
              <label htmlFor="briefing">Briefing</label>
              <textarea
                id="briefing"
                value={briefing}
                name="briefing"
                onChange={(e) => setBriefing(e.target.value)}
              ></textarea>
            </fieldset>
          </div>
          <div className="box">
            <label
              htmlFor="checkboxDefault"
              className="w-full inline-flex min-w-[14rem] cursor-pointer items-center justify-between rounded-xl gap-3 border border-slate-300 bg-slate-100 px-4 py-2 text-sm font-medium text-slate-700 dark:border-slate-700 dark:bg-slate-800 dark:text-slate-300 [&:has(input:checked)]:text-black dark:[&:has(input:checked)]:text-white [&:has(input:disabled)]:opacity-75 [&:has(input:disabled)]:cursor-not-allowed"
            >
              <span>Enable AI Texts</span>
              <div className="relative flex items-center">
                <input
                  id="checkboxDefault"
                  type="checkbox"
                  name="use_ai"
                  checked={useAi}
                  onChange={() => {
                    console.log(!useAi);
                    setUseAi(!useAi);
                  }}
                />
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  aria-hidden="true"
                  stroke="currentColor"
                  fill="none"
                  stroke-width="4"
                  className="pointer-events-none invisible absolute left-1/2 top-1/2 size-3 -translate-x-1/2 -translate-y-1/2 text-slate-100 peer-checked:visible dark:text-slate-100"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    d="M4.5 12.75l6 6 9-13.5"
                  />
                </svg>
              </div>
            </label>
          </div>
          <div className="box">
            <fieldset>
              <button type="submit" data-state={isLoading && "loading"}>
                Salvar
              </button>
            </fieldset>
          </div>
        </div>
      </div>
    </form>
  );
}
