import React, { useState, useEffect } from "react";
import mqtt from "mqtt";

const DeviceForm: React.FC = () => {

  const [mqttClient, mqqtSetClient] = useState<mqtt.MqttClient>();

  useEffect(() => {
    const mqttURL = import.meta.env.VITE_MQTT_URL;
    const clientPass = import.meta.env.VITE_MQTT_KEY;
    const clientName = import.meta.env.VITE_MQTT_USERNAME;

    const client = mqtt.connect(mqttURL, {
      username: clientName,
      password: clientPass,
    });

    client.on("connect", () => {
      console.log("Connected to MQTT broker");
      mqqtSetClient(client);
    });

    // Cleanup function to disconnect MQTT client when the component unmounts
    return () => {
      client.end(); // Disconnect MQTT client
    };
  }, []); // Define state variables to hold form data

  const deviceTypes = ["Temperature Sensor", "Refrigerator", "House Lights"];

  const [formData, setFormData] = useState({
    name: "",
    selectedType: deviceTypes[0]
  });

  // Define a function to handle form input changes
  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleTypeChange = (event: React.ChangeEvent<HTMLSelectElement>) => {

    const { name, value } = event.target;
    setFormData({
      ...formData,
      [name]: value,
    })
  }

  // Define a function to handle form submission
  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();

    if(mqttClient) {
      mqttClient.publish('home', JSON.stringify(formData));
    } 
  };

  return (
    <div className="max-w-md mx-auto mt-8 p-6 bg-gray-50 rounded-lg shadow-lg">
      <h1 className="text-2xl font-semibold mb-4">Create an IoT Device</h1>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700">
            Name
          </label>
          <input
            type="text"
            name="name"
            value={formData.name}
            onChange={handleInputChange}
            className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
          />
        </div>
        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700">
            Device Type
          </label>
          <select
            name="selectedType"
            onChange={handleTypeChange}
            className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:outline-none focus:ring focus:border-blue-300"
          >
            {Array.from(deviceTypes, (data) => (
              <option key={data} value={data}>
                {data}
              </option>
            ))}
          </select>
        </div>
        <div>
          <button
            type="submit"
            className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:bg-blue-600"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default DeviceForm;
