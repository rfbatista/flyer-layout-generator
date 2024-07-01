import "./layout_card.css";
import { Layout } from "../../domain/layout/entities/layout";

type Props = {
  layout: Layout;
};

export function LayoutCard(props: Props) {
  const d = props.layout;
  return (
    <article className="layout-card">
      <div className="stack">
        <div className="layout-card__image-container">
          <img
            src={
              d.imageURL
                ? d.imageURL
                : "https://placehold.co/240x176?text=Need+Process"
            }
            alt="view of a coastal Mediterranean village on a hillside, with small boats in the water."
          />
        </div>
        <div className="layout-card__body">
          <div className="cluster">
            <div>
              <a href={d.imageURL} download data-type="button">
                download
              </a>
              <a href={`/editor?design=${d.designID}&layout=${d.id}`} data-type="button">
                edit
              </a>
            </div>
          </div>
        </div>
      </div>
    </article>
  );
}
