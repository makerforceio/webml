import * as tf from '@tensorflow/tfjs';
import axios from 'axios';

/*
* Params for logging
*  -> Epoch
*  -> Loss
*  -> Accuracy
*  -> Batch num
*  -> Runtime
*/
class DistTensorflow {
  token;
  model;
  http;
  batchSize;
  batchNo = 0;
  stopped = false;

  statsCallback;

  constructor(token, statsCallback) {
    this.token = token;
    this.statsCallback = statsCallback;

    // Initialize axios instance
    this.http = axios.create({
		baseUrl: 'https://some-domain.com/api/',
      timeout: 10000,
    });

    this.model = await tf.loadLayersModel(‘path/to/model.json’);

    // Compile the model with default optimizer and loss
    this.model.compile({
      optimizer: tf.train.adam(),
      loss: 'categoricalCrossentropy',
      metrics: ['accuracy'],
    });
  }

  loadNextBatch() async {
    // Load the next batch from the backend
    let res = await http.get('metadata');

    const batchShape = res.data.batch;
    const labelShape = res.data.label;

    // Set batch size
    this.batchSize = batchShape[0];

    // Load the minibatch data
    res = await http.get('batch', {responseType: 'arraybuffer'});
    let batchArray = new UInt8Array(res.data);

    // Load the minibatch labels
    res = await http.get('label', {responseType: 'arraybuffer'});
    let labelArray = new UInt8Array(res.data);

    this.batchNo += 1;
    return {
      "data": tf.tensor(batchArray, {shape: batchShape}),
      "labels": tf.tensor(labelArray, {shape: labelShape})
    };
  }

  updateWeights() async {
    let oldWeights = this.model.getWeights();

    let res = await http.post('/weights', {
      shape: oldWeights.shape,
      data: await oldWeights.flatten().array()
    });

    let weights = tf.tensor(res.data.data, {shape: res.data.shape});

    this.model.setWeights(weights);
  }

  async train() {
    // Train on the minibatch
    while(!stopped) {
      let minibatch = await loadNextBatch()
      let metrics = await this.model.trainOnBatch(minibatch.data, minibatch.label);
      await updateWeights();

      // Callbacks for statistics
      statsCallback(metrics, this.batchNo);
    }
  }

  stop() {
    this.stopped = true;
  }
}
