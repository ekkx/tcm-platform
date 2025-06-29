import { type RouteConfig, layout, route } from "@react-router/dev/routes";

export default [
  route("/login", "./routes/login.tsx"),

  layout("./routes/layout.tsx", [
    route("/home", "./routes/home.tsx"),
    route("/profile", "./routes/profile.tsx"),
  ]),
] satisfies RouteConfig;
