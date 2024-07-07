import "../../components/button/button.css";
import "../../components/input/input.css";
import "../../components/label/label.css";
import { Modal } from "../../components/modal/modal";
import { useModal } from "../../components/modal/store";
import "../../components/select/select.css";
import { TopBar } from "../../components/topbar/topbar";
import "./App.css";
import { CreateProjectForm } from "./components/create_project_form";
import { ProjectsTable } from "./components/projects_table";

export default function App() {
  const { toggle } = useModal();
  return (
    <>
      <Modal title="Create a Project">
        <CreateProjectForm />
      </Modal>
      <TopBar items={[{ title: "projects", link: "" }]} />
      <div className="container">
        <div className="" style={{marginTop: "var(--s1)"}}>
          <div className="stack">
            <div className="cluster">
              <div className="projects-page__create-project-btn p-1">
                <button onClick={toggle}>
                  <div className="ld ld-ring ld-spin"></div>
                  Criar projeto
                </button>
              </div>
            </div>
            <ProjectsTable />
          </div>
        </div>
      </div>
    </>
  );
}
