import React from "react";

const AddressInput = ({ address, onAddressChange }) => {
  return (
    <div>
      <p>XXX Testnet Address</p>
      <input
        className="input"
        type="text"
        placeholder="Recipient address"
        value={address}
        onChange={onAddressChange}
      />
    </div>
  );
};

export default AddressInput;
