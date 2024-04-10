import React from "react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSync } from "@fortawesome/free-solid-svg-icons";

const Claim = ({ taskid, buttonState, handleStart, handleClaim }) => {
  var text1 = "";
  var text2 = "";
  if (taskid === 1) {
    text1 = "Follow XXX on Twitter and repost Twitter.";
    text2 = "+1000USDT";
  } else if (taskid === 2) {
    text1 = "Join XXX on Telegram.";
    text2 = "+1000USDT";
  } else if (taskid === 3) {
    text1 = "Join XXX on Discord.";
    text2 = "+1000USDT";
  }

  const StartButton = () => {
    switch (buttonState) {
      case "claimed":
        return (
          <button disabled style={{ background: "#70c9f6" }}>
            {text2}
          </button>
        );
      case "claiming":
        return <button disabled>Claiming...</button>;
      case "start":
        return <button onClick={() => handleStart(taskid)}>START</button>;
      default:
        return <button disabled>START</button>;
    }
  };

  const RefreshButton = () => {
    if (buttonState === "claiming") {
      return (
        <button className="refresh-button" onClick={() => handleClaim(taskid)}>
          <FontAwesomeIcon icon={faSync} />
        </button>
      );
    } else {
      return (
        <button className="refresh-button" disabled>
          <FontAwesomeIcon icon={faSync} />
        </button>
      );
    }
  };

  return (
    <div className="claim-component">
      <span>{text1}</span>
      <div style={{ flexGrow: 1 }}></div>
      <div
        className="center"
        style={{
          margin: "0 20px",
          color: "#8a2be2",
        }}
      >
        <span>{text2}</span>
      </div>

      <div className="center">
        <StartButton></StartButton>
      </div>

      <div className="center">
        <RefreshButton></RefreshButton>
      </div>
    </div>
  );
};

export default Claim;
