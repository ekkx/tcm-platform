import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("login", "routes/login.tsx"),
  route("login-grpc", "routes/login-grpc.tsx"),
] satisfies RouteConfig;
