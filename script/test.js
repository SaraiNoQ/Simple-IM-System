const net = require('net');

// 目标主机和端口
const host = '127.0.0.1';
const port = 8888;

// 发送请求的次数
const numRequests = 10;

// 追踪已完成的请求数量
let completedRequests = 0;

for (let i = 0; i < numRequests; i++) {
    const client = new net.Socket();

    client.connect(port, host, () => {
        // 当连接建立时发送数据
        client.write(`Request ${i + 1}`);
    });

    client.on('data', (data) => {
        // 处理服务器的响应
        console.log(`Received+${i + 1}: ${data}`);

        // 关闭连接
        // client.destroy();
    });

    client.on('close', () => {
        // 当连接关闭时，记录完成的请求
        completedRequests++;
        if (completedRequests === numRequests) {
            console.log('All requests completed.');
        }
    });

    client.on('error', (err) => {
        console.error(`Error: ${err.message}`);
    });
}