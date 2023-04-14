import { A } from "@solidjs/router";

function Nav() {
  return (
    <nav>
      <ul>
        <li>
          <strong>
            <A href="/">Brand</A>
          </strong>
        </li>
      </ul>
      <ul>
        <li>
          <A href="/users">Users</A>
        </li>
        <li>
          <A href="/login">Login</A>
        </li>
      </ul>
    </nav>
  );
}

export default Nav;
