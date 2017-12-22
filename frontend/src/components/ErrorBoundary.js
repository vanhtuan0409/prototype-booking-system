import { PureComponent } from "react";

export default class ErrorBoundary extends PureComponent {
  componentDidCatch(error, info) {
    console.log(error, info);
  }

  render() {
    return this.props.children;
  }
}
