import React from "react";
import ReactDOM from "react-dom";
import ErrorBoundary from "./components/ErrorBoundary";
import App from "./components/App";

import toastr from "toastr";
import "toastr/build/toastr.min.css";
toastr.options.timeOut = 500;

ReactDOM.render(
  <ErrorBoundary>
    <App />
  </ErrorBoundary>,
  document.getElementById("root")
);
