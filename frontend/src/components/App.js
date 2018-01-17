import React, { PureComponent } from "react";
import autobind from "class-autobind";
import io from "socket.io-client";
import toastr from "toastr";
import Resource from "./Resource";
import resourcesApi from "../apis/resources";

const styles = {
  width: "100%",
  height: "100%",
  display: "flex",
  flexWrap: "wrap"
};

export default class App extends PureComponent {
  constructor(props) {
    super(props);
    autobind(this);
    this.state = {
      resources: []
    };
  }

  async componentDidMount() {
    const resources = await resourcesApi.getAll();
    this.setState({ resources });

    const socket = io(process.env.REACT_APP_SOCKET_URL, {
      path: "/subscribe"
    });
    this.wsListen(socket);
  }

  setResourceAsBooked(resourceId) {
    const newResources = this.state.resources.map(r => {
      if (r.id === resourceId) {
        return Object.assign({}, r, { available: false });
      }
      return r;
    });
    this.setState({ resources: newResources });
  }

  restoreAllResource() {
    const newResources = this.state.resources.map(r =>
      Object.assign({}, r, { available: true })
    );
    this.setState({ resources: newResources });
  }

  wsListen(socket) {
    socket.on("resource_booked", resourceId => {
      this.setResourceAsBooked(parseInt(resourceId, 10));
    });
    socket.on("all_resources_restored", () => {
      this.restoreAllResource();
    });
  }

  async onSelect(resource) {
    try {
      await resourcesApi.book(resource.id);
      this.setResourceAsBooked(resource.id);
    } catch (error) {
      console.log(error);
      toastr.error("Resource have already been book");
      return;
    }
  }

  async onRestore() {
    await resourcesApi.resetAll();
    this.restoreAllResource();
  }

  render() {
    const { resources } = this.state;
    return (
      <div>
        <p>
          <button onClick={this.onRestore}>Restore all resource</button>
        </p>
        <div style={styles}>
          {resources.map(r => (
            <Resource key={r.id} resource={r} onSelect={this.onSelect} />
          ))}
        </div>
      </div>
    );
  }
}
