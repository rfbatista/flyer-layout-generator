import "./tabs.css";
type Props = {
  items: {
    name: string;
    onClick: () => void;
    active: boolean;
  }[];
};

export function Tabs(props: Props) {
  return (
    <>
      <div className="tabs">
        {props.items.map((i) => {
          return (
            <div className="tabs__container" data-state={i.active && "active"}>
              <button data-type="ghost" className="tabs__container__item" onClick={i.onClick}>
                {i.name}
              </button>
            </div>
          );
        })}
      </div>
    </>
  );
}
