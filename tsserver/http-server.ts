import uwebsocket from "uWebSockets.js";
import config from "./config";

export default class HttpServer {
  private app;
  constructor() {
    this.app = uwebsocket.App();
  }
  start() {
    this.app.listen(config.host, config.port);
  }
}
