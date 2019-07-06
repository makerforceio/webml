const express = require('express');
const tf = require('@tensorflow/tfjs-node');

const paramsRouter = require('./routes/params');
const adminRouter = require('./routes/params');

// Creating the app
const app = express();

// Register the routes
app.use('/params', paramsRouter);

