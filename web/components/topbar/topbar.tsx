import "./topbar.css";
import "./logo.png";
import { Breadcrump } from "../breadcrump/breadcrump";

type Props = {
  items: { title: string; link: string }[];
};

export function TopBar(props: Props) {
  return (
    <nav className="topbar">
      <div className="topbar__logo">
      <img src={"/dist/vite/assets/logo.png"}  />
      </div>
      <Breadcrump items={props.items} />
      <ul className="topbar__list">
        <li>
          <a href="/" className="topbar__list__item" aria-current="page">
            Projects
          </a>
        </li>
        <li className="topbar__list__profile">
          <img
            src="https://penguinui.s3.amazonaws.com/component-assets/avatar-8.webp"
            alt="User Profile"
            className="topbar__list__profile__pic"
          />
        </li>
      </ul>
    </nav>
  );
}
