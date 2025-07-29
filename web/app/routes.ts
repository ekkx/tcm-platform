import { type RouteConfig, layout, route } from "@react-router/dev/routes";

export default [
  route("/login", "./routes/login.tsx"),

  layout("./routes/internal/layout.tsx", [
    route("/home", "./routes/internal/home.tsx"),
    route("/profile", "./routes/internal/profile.tsx"),
  ]),
] satisfies RouteConfig;
