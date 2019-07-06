import * as tf from '@tensorflow/tfjs';

class DistTensorflow {
  token;
  model;
  http;
  batchsize;

  constructor(token) {
    this.token = token;

    // Initialize axios instance
    this.http = axios.create({
      baseUrl: 'https://some-domain.com/api/'
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
    this.batchsize = batchShape[0];

    // Load the minibatch data
    res = await http.get('batch', {responseType: 'arraybuffer'});
    let batchArray = new UInt8Array(res.data);

    // Load the minibatch labels
    res = await http.get('label', {responseType: 'arraybuffer'});
    let labelArray = new UInt8Array(res.data);

    return {
      "data": tf.tensor(batchArray, {shape: batchShape}),
      "labels": tf.tensor(labelArray, {shape: labelShape})
    };
  }

  updateWeights() async {
    let old_weights = this.model.getWeights();

    await http.post('/weights', {
      shape: old_weights.shape,
      data: old_weights.flatten()
    });

    let res = await http.get('/weights');
    let new_weights = tf.tensor(res.data.data, {shape: res.data.shape});

    this.model.setWeights(new_weights);
  }

  async train() {
    // Train on the minibatch

  }
}
