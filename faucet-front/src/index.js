import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <React.StrictMode>
    <AppContainer></AppContainer>
  </React.StrictMode>
);

function AppContainer() {
  //先检查页面参数
  const urlParams = new URLSearchParams(window.location.search);
  const source = urlParams.get("source");
  const state = urlParams.get("state");
  const code = urlParams.get("code");

  var taskid = 0;
  if (state !== "") {
    if (source === "twitter") {
      taskid = 1;
    } else if (source === "telegram") {
      taskid = 2;
    } else if (source === "discord") {
      taskid = 3;
    }
  }

  return (
    <div class="appContainer">
      <App
        initialAddress={state}
        initialTaskid={taskid}
        twitterCode={code}
      ></App>
    </div>
  );
}
