import React, { PureComponent } from "react";

const styles = {
  width: 70,
  height: 70,
  margin: 20,
  border: "1px solid black",
  cursor: "pointer",
  display: "flex",
  justifyContent: "center",
  alignItems: "center"
};

const availableStyles = {
  backgroundColor: "#B2EBF2",
  color: "#000"
};

const unavailableStyles = {
  backgroundColor: "#F50057",
  color: "#fff"
};

export default class componentName extends PureComponent {
  render() {
    const { resource, onSelect } = this.props;
    const style = resource.available
      ? Object.assign({}, styles, availableStyles)
      : Object.assign({}, styles, unavailableStyles);
    return (
      <div onClick={() => onSelect(resource)} style={style}>
        {resource.available ? "Select" : "Booked"}
      </div>
    );
  }
}
