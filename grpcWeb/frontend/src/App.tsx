import React, { useState } from 'react';
import './App.css';
import { RunnerClient } from './runner_grpc_web_pb';
import { RunnerRequest } from './runner_pb';

const HOST = 'https://localhost:8090';

export const App = () => {
  const onClickRun = () => {
    let client = new RunnerClient(HOST);
    let request = new RunnerRequest();
    request.setName('Golang');
    client.run(request, { "key" : "value" }, (error, reply) => {
      if (reply) {
        console.log(reply.getMessage());
      }
      if (error) {
        console.error(error);
      }
    });
  };
  return (
    <>
      <h1>grpc web</h1>
      <button onClick={onClickRun}>Run</button>
    </>
  );
}