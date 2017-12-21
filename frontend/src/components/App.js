import React, { PureComponent } from "react";
import toastr from "toastr";
import Resource from "./Resource";
import getResouces from "../apis/get_resource";

const styles = {
  width: "100%",
  height: "100%",
  display: "flex",
  flexWrap: "wrap"
};

export default class componentName extends PureComponent {
  constructor(props) {
    super(props);
    this.state = {
      resources: []
    };
  }

  async componentDidMount() {
    const resources = await getResouces();
    this.setState({ resources });
  }

  onSelect(resource) {
    if (!resource.available) {
      toastr.error("Cannot select a booked resource");
      return;
    }
    console.log(resource);
  }

  render() {
    const { resources } = this.state;
    return (
      <div style={styles}>
        {resources.map(r => (
          <Resource key={r.id} resource={r} onSelect={this.onSelect} />
        ))}
      </div>
    );
  }
}
