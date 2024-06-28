import { capitalizeWords } from "../../shared/captalize"
import "./breadcrump.css"

type Props = {
  items: string[]
}

export function Breadcrump({ items = [] }: Props) {
  return (
    <nav className="breadcrump box">
      <ol className="breadcrump__list">
        {
          items.map((i, idx) => {
            return idx == 0
              ?
              <li className="breadcrump__list__item">
                <a href="#" className="breadcrump__list__item__link">
                  {capitalizeWords(i)}
                </a>
                <span className="breadcrump__list__item__icon" />
              </li>
              :
              <li data-position={idx === (items.length - 1) && "last"} className="breadcrump__list__item" aria-current="page">{capitalizeWords(i)}</li>

          })
        }
      </ol>
    </nav>
  )
}

