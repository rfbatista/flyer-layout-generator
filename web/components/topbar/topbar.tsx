import "./topbar.css";
import "./logo.png";

export function TopBar() {
  return (
    <nav className="topbar">
      <img src={"/dist/vite/assets/logo.png"} className="topbar__logo"/>
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
