import { useProjectsStore } from "../../../domain/projects/store"
import "./project_table.css"

export function ProjectsTable() {
  const { projects } = useProjectsStore()
  return <table className="project-table">
    <thead>
      <tr>
        <th scope="col" className="project-table__header__item">Name</th>
        <th scope="col" className="project-table__header__item">Client</th>
        <th scope="col" className="project-table__header__item">Advertiser</th>
        <th scope="col" className="project-table__header__item">Created at</th>
      </tr>
    </thead>
    {
      projects.map((p) => {
        return <tr>
          <td className="project-table__body__item">
            {p.name}
          </td >
          <td className="project-table__body__item">
            {p.name}
          </td>
          <td className="project-table__body__item">
            {p.name}
          </td>
          <td className="project-table__body__item">
            {p.name}
          </td>
        </tr>
      })
    }
    <tbody>

    </tbody>

  </table>
}
