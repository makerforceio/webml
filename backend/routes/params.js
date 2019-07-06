const express = require('express');
const bodyParser = require('body-parser')
const tf = require('@tensorflow/tfjs-node');

const router = express.Router();
const jsonParser = bodyParser.json()

const ALPHA = 0.95;

let paramsMap = new Map();

router.post('/update/:token', jsonParser, function(req, res) {
  let shape = req.body.shape;
  let rawWeights = req.body.data;
  const token = req.params.token;

  let weights = tf.tensor(rawWeights, {shape: shape});

  if(paramsMap.has(token)) {
    let oldWeights = paramsMap.get(token);
    let newWeights = tf.movingAverage(oldWeights, weights, ALPHA);

    res.send({
      shape: newWeights.shape,
      data: newWeights.flatten().array()
    });
  } else {
    paramsMap.set(token, weights);

    res.send({
      shape: weights.shape,
      data: weights.flatten().array()
    });
  }
});

module.exports = router;
