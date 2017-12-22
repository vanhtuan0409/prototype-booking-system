import React, { PureComponent } from "react";
import autobind from "class-autobind";
import toastr from "toastr";
import Resource from "./Resource";
import resourcesApi from "../apis/resources";

const styles = {
  width: "100%",
  height: "100%",
  display: "flex",
  flexWrap: "wrap"
};

export default class componentName extends PureComponent {
  constructor(props) {
    super(props);
    autobind(this);
    this.state = {
      resources: []
    };
  }

  async componentDidMount() {
    try {
      const resources = await resourcesApi.getAll();
      this.setState({ resources });
    } catch (error) {
      toastr(error);
    }
  }

  async onSelect(resource) {
    if (!resource.available) {
      toastr.error("Cannot select a booked resource");
      return;
    }
    try {
      await resourcesApi.book(resource.id);
      const newResources = this.state.resources.map(r => {
        if (r.id === resource.id) {
          return Object.assign({}, r, { available: false });
        }
        return r;
      });
      this.setState({ resources: newResources });
    } catch (error) {
      toastr(error);
    }
  }

  async onRestore() {
    try {
      await resourcesApi.resetAll();
      const newResources = this.state.resources.map(r =>
        Object.assign({}, r, { available: true })
      );
      this.setState({ resources: newResources });
    } catch (error) {
      toastr(error);
    }
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
