import { useEffect } from "react";
import { Breadcrump } from "../../components/breadcrump/breadcrump";
import { Modal } from "../../components/modal/modal";
import { TopBar } from "../../components/topbar/topbar";
import { useProjectsStore } from "../../domain/projects/store";
import "./App.css";
import { useDesignsStore } from "../../domain/design/store";
import { DesginCard } from "../../components/designcard/design_card";
import { NewDesignForm } from "./components/new_design_form";
import { useModal } from "../../components/modal/store";
import { SideBar } from "./components/side-bar";

export default function App() {
  const { toggle, setCh, setTitle } = useModal();
  const { activeProject, setActiveProject } = useProjectsStore();
  const { listDesigns, designs } = useDesignsStore();
  useEffect(() => {
    let params = new URLSearchParams(document.location.search);
    let id = params.get("id");
    setActiveProject(Number(id));
    listDesigns(Number(id));
  }, []);
  const openModal = () => {
    setCh(<NewDesignForm />);
    setTitle("New design");
    toggle();
  };
  return (
    <>
      <Modal title="New design"></Modal>
      <TopBar
        items={[
          { title: "projects", link: "/" },
          {
            title: activeProject ? activeProject.name : "",
            link: "",
          },
        ]}
      />
      <div className="right-sidebar">
        <div>
          <div className="box">
            <div className="container">
              <div>
                <div className="stack">
                  <div className="cluster">
                    <div className="projects-page__upload-design-btn">
                      <button onClick={openModal}>New Design</button>
                    </div>
                  </div>
                  <div className="stack">
                    <div>
                      <h2>Masters</h2>
                    </div>
                    <div className="cluster">
                      <div>
                        {designs.map((d) => {
                          return <DesginCard design={d} />;
                        })}
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="box">
            <SideBar />
          </div>
        </div>
      </div>
    </>
  );
}
