import AddressInput from "./components/address";
import Claim from "./components/claim";
import DualLinks from "./components/dualLinks";
import "./App.css";
import { useState, useEffect } from "react";

function App({ initialAddress, initialTaskid, twitterCode }) {
  // "" start claimed claiming
  const [btnStates, setBtnStates] = useState(["", "", ""]);
  const [inputAddress, setInputAddress] = useState("");

  useEffect(() => {
    mounted();
  }, []);

  //生命周期
  const mounted = async () => {
    const customEvent = { target: { value: initialAddress } };
    await onAddressChange(customEvent);

    if (initialTaskid === 1 && twitterCode.length > 0) {
      //上传授权结果
      const address = initialAddress;
      const extra = twitterCode;
      const taskid = initialTaskid;
      if (!checkAddress(address)) {
        return;
      }

      const requestData = { address, taskid, extra };
      try {
        const response = await fetch("/api/v1/twitterUploadCode", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(requestData),
        });

        const data = await response.text();
        console.log(data);

        handleStart(taskid);
      } catch (error) {
        console.error("Error:", error);
      }
    }
  };

  const onAddressChange = async (event) => {
    const value = event.target.value;
    setInputAddress(value);

    if (!checkAddress(value)) {
      setBtnStates(["", "", ""]);
      return;
    }

    try {
      const requestData = { address: value };
      const response = await fetch("/api/v1/check_reward", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestData),
      });

      const data = await response.text();
      console.log(data);

      const info = JSON.parse(data);
      const newArray = [...btnStates];
      newArray[0] = info.hash1.length === 0 ? "start" : "claimed";
      newArray[1] = info.hash2.length === 0 ? "start" : "claimed";
      newArray[2] = info.hash3.length === 0 ? "start" : "claimed";
      setBtnStates(newArray);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const handleStart = async (taskid) => {
    if (!checkAddress(inputAddress)) {
      return;
    }

    const requestData = { address: inputAddress, taskid };
    try {
      const response = await fetch("/api/v1/shareUrl", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestData),
      });

      const data = await response.text();
      console.log(data);
      const info = JSON.parse(data);

      const code = info.code;
      const url = info.extra;
      if (code === 100 && url.length > 1) {
        if (url.includes("oauth2/")) {
          window.location.href = url;
        } else {
          window.open(url, "_blank");
        }

        const newArray = [...btnStates];
        newArray[taskid - 1] = "claiming";
        setBtnStates(newArray);
      } else {
        alert("Failed to get the share link!");
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const handleClaim = async (taskid) => {
    const address = inputAddress;
    const extra = twitterCode;
    if (!checkAddress(address)) {
      return;
    }

    const requestData = { address, taskid, extra };
    try {
      const response = await fetch("/api/v1/claim_reward", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestData),
      });

      const data = await response.text();
      console.log(data);
      const info = JSON.parse(data);

      if (info.code !== 100 && info.msg.length > 1) {
        alert(info.msg);
      }

      const newArray = [...btnStates];
      newArray[0] = info.hash1.length === 0 ? "start" : "claimed";
      newArray[1] = info.hash2.length === 0 ? "start" : "claimed";
      newArray[2] = info.hash3.length === 0 ? "start" : "claimed";
      setBtnStates(newArray);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <div class="app">
      <h2>XXX Testnet Faucet</h2>

      <AddressInput address={inputAddress} onAddressChange={onAddressChange} />

      <h3>Get XXX Testnet 3000USDT</h3>

      {btnStates.map((item, index) => (
        <Claim
          taskid={index + 1}
          buttonState={item}
          handleStart={handleStart}
          handleClaim={handleClaim}
        ></Claim>
      ))}

      <DualLinks link1Text="XXX Faucet" link2Text="XXX Testnet" />
    </div>
  );

  function checkAddress(address) {
    const regex = /^(0x)?[0-9a-fA-F]{40}$/;
    return regex.test(address);
  }
}

export default App;
