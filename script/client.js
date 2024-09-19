const net = require('net');

// 目标主机和端口
const host = '127.0.0.1';
const port = 8888;
const client = new net.Socket();

client.connect(port, host, () => {
    // 当连接建立时发送数据
    client.write(`Request`);
});

client.on('data', (data) => {
    // 处理服务器的响应
    console.log(`Received: ${data}`);

    // 关闭连接
    // client.destroy();
});

client.on('close', () => {
    // 当连接关闭时，记录完成的请求
});

client.on('error', (err) => {
    console.error(`Error: ${err.message}`);
});