const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');

// Load the Protocol Buffers definition
const packageDefinition = protoLoader.loadSync('../protofile.proto', {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const grpcPackage = grpc.loadPackageDefinition(packageDefinition);

const client = new grpcPackage.main.PingPongService('localhost:50051', grpc.credentials.createInsecure());
let executionsCount = 0;
const clientId = Math.floor(Math.random() * 100)
const numberOfExecutions = 20;
const timeOut = 2000;
const stream = client.HandleMessage({
    body: `Hello, gRPC from nodejs client ${clientId}, number ${executionsCount}!`,
    important: true,
})

const sendMessage = () => {
    const message = {
        body: `Hello, gRPC from nodejs client ${clientId}, number ${executionsCount}!`,
        important: true,
    };
    stream.write(message)
}

// Receive messages from the server
stream.on('data', (response) => {
    console.log('Received from server:', response.body);
});

// Handle errors and end event
stream.on('error', (error) => {
    console.error('Error:', error);
});

stream.on('end', () => {
    console.log('Server stream ended.');
});


const interval = setInterval(sendMessage, timeOut);
const stopInterval = setInterval(() => {
    executionsCount++;
    if (executionsCount >= numberOfExecutions) {
        clearInterval(interval);
        clearInterval(stopInterval);
        stream.end();
    }
}, timeOut);
