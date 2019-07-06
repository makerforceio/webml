import * as tf from '@tensorflow/tfjs';

/*
* Params for logging
*  -> Epoch
*  -> Loss
*  -> Accuracy
*  -> Batch num
*  -> Runtime
*/

class DistTensorflow {
  modelId;
  model;
  batchSize;
  batchNo = 0;
  stopped = false;

  statsCallback;

  constructor(modelId, statsCallback) {
    this.modelId = modelId;
    this.statsCallback = statsCallback;

  tf.loadLayersModel(`http://localhost:10200/model?id=${this.modelId}`).then(function (model) {
    this.model = model;

    // Compile the model with default optimizer and loss
    this.model.compile({
      optimizer: tf.train.adam(),
      loss: 'categoricalCrossentropy',
      metrics: ['accuracy'],
    });
  });
  }

  async loadNextBatch() {
    // Load the next batch from the backend
    let res = await http.get('metadata');

    const batchShape = res.data.batch;
    const labelShape = res.data.label;

    // Set batch size
    this.batchSize = batchShape[0];

    // Load the minibatch data
    res = await fetch(`localhost:10200/data/batch?model=${this.modelId}`, {
      method: 'GET',
      redirect: 'follow',
    });

    let batchArray = new UInt8Array(await res.arrayBuffer());

    // Load the minibatch labels
    res = await fetch(`localhost:10200/label/batch?model=${this.modelId}`, {
      method: 'GET',
      redirect: 'follow',
    });

    let labelArray = new UInt8Array(await res.arrayBuffer());

    this.batchNo += 1;
    return {
      "data": tf.tensor(batchArray, batchShape),
      "labels": tf.tensor(labelArray, labelShape)
    };
  }

  async updateWeights() {
    let oldWeights = this.model.getWeights();

    let res = await fetch(`localhost:10300/params/${this.modelId}`, {
      method: 'POST',
      body: JSON.stringify({
        shape: oldWeights.shape,
        data: await oldWeights.flatten().array()
      })
    });

    let resJSON = await res.json();

    let weights = tf.tensor(resJSON.data, resJSON.shape);

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

export default DistTensorflow;
