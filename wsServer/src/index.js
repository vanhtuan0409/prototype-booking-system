import express from "express";
import { createServer } from "http";
import morgan from "morgan";
import bodyParser from "body-parser";
import cors from "cors";
import ws from "socket.io";

const app = express();
const server = createServer(app);
const io = ws(server, {
  path: "/subscribe",
  serveClient: false,
  pingInterval: 10000,
  pingTimeout: 5000
});

// Config express
app.use(morgan("dev"));
app.use(cors());
app.options("*", cors());
app.use(bodyParser.json());

app.post("/resources/:id/book", function(req, res) {
  const resourceId = req.params.id;
  io.sockets.emit("resource_booked", resourceId);
  res.json({
    success: true
  });
});
app.post("/resources/restore", function(req, res) {
  io.sockets.emit("restore_all_resources", true);
  res.json({
    success: true
  });
});

// Send ws message
setInterval(function() {
  io.sockets.emit("hello", "world");
}, 2000);

const port = process.env.HTTP_PORT || 4000;
server.listen(port, function() {
  console.log(`Socket server is running on :${port}`);
});
