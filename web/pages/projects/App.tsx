import { Breadcrump } from "../../components/breadcrump/breadcrump";
import { TopBar } from "../../components/topbar/topbar";
import { ProjectsTable } from "./components/projects_table";
import { Modal } from "../../components/modal/modal";
import { useModal } from "../../components/modal/store";
import { CreateProjectForm } from "./components/create_project_form";
import "./App.css";
import "../../components/input/input.css";
import "../../components/label/label.css";
import "../../components/select/select.css";
import "../../components/button/button.css";

export default function App() {
  const { toggle } = useModal();
  return (
    <>
      <Modal title="Create a Project">
        <CreateProjectForm />
      </Modal>
      <TopBar />
      <div className="container">
        <div>
          <div className="stack">
            <Breadcrump items={[{ title: "projects", link: "" }]} />
            <div className="cluster">
              <div className="projects-page__create-project-btn">
                <button className="btn" onClick={toggle}>
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
