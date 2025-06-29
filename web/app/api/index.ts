import { createConnectTransport } from "@connectrpc/connect-web";

export const transport = createConnectTransport({
  baseUrl: "http://localhost:50051",
  // useBinaryFormat: true,
});
