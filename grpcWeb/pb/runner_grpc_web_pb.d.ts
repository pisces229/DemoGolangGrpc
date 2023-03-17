import * as grpcWeb from 'grpc-web';

import * as runner_pb from './runner_pb';


export class RunnerClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  run(
    request: runner_pb.RunnerRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: runner_pb.RunnerResponse) => void
  ): grpcWeb.ClientReadableStream<runner_pb.RunnerResponse>;

}

export class RunnerPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  run(
    request: runner_pb.RunnerRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<runner_pb.RunnerResponse>;

}

