import { createSignal, createEffect } from "solid-js";
import { lazy } from "solid-js";
import { Routes, Route, A } from "@solidjs/router";
import Header from "./components/Header";
import Counter from "./components/Counter";
import Nav from "./components/Nav";
// import NotFound from "./pages/NotFound";
// import Home from "./pages/Home";
// import Users from "./pages/Users";

function App() {
  // component state
  //function value, function update
  // const [counter, setCounter] = createSignal(0);

  //to turn this into a computed value you have to wrap it in a fuction
  // const doubleCounter = counter() * 2;
  // const doubleCounter = () => counter() * 2;

  // side effects
  // createEffect(() => {
  //   // update a signal will trigger this effect will run again
  //   console.log("Effect - execute because counter updated", counter());
  // });

  const Home = lazy(() => import("./pages/Home"));
  const Users = lazy(() => import("./pages/Users"));
  const NotFound = lazy(() => import("./pages/NotFound"));
  // component template
  return (
    <main class="container">
      <Nav />
      <Routes>
        <Route path="/" component={Home} />
        <Route path="/users" component={Users} />
        <Route
          path={["login", "sign-up"]}
          element={<h1>This is the login/sign up screen</h1>}
        />
        <Route path="*" component={NotFound}></Route>
      </Routes>
    </main>
  );
}

export default App;
