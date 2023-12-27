// server.js
const express = require('express');
const cors = require('cors'); 
const { exec } = require('child_process');

const app = express();
const port = 3000;

app.use(cors());

app.use(express.json());

app.post('/execute', (req, res) => {
    const command = req.body.command;

    exec(command, (error, stdout, stderr) => {
        if (error) {
            res.json({ output: stderr });
        } else {
            res.json({ output: stdout });
        }
    });
});

app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
