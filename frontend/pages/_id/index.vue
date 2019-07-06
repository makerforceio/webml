<template>
  <div>
    <Header :title="title" subtitle="by MakerForce">
      <template v-slot:buttons-left>
        <button class="my-2 mr-4">
          <fa-icon size="2x" :icon="['far', 'arrow-left']" />
        </button>
      </template>
      <template v-slot:buttons-right>
        <button v-if="running" class="my-2 ml-4" @click="toggleRunning">
          <fa-icon size="2x" :icon="['far', 'pause']" />
        </button>
        <button v-else class="my-2 ml-4" @click="toggleRunning">
          <fa-icon size="2x" :icon="['far', 'play']" />
        </button>
      </template>
      <template v-slot:content>
        <div class="mx-3 my-6" title="Loss graph">
          <h2 class="font-medium">
          </h2>
          <trend
            :data="major"
            :gradient="gradient"
            :height="200"
            auto-draw
            smooth
          >
          </trend>
        </div>
      </template>
    </Header>
    <Cards>
      <Card subtitle="Elapsed time">
        <CenteredText class="text-4xl">
          10m 45s
        </CenteredText>
      </Card>
      <Card subtitle="Loss">
        <CenteredText class="text-4xl">
          0.567
        </CenteredText>
      </Card>
      <Card subtitle="Batch No.">
        <CenteredText class="text-4xl">
          20
        </CenteredText>
      </Card>
      <Card subtitle="Accuracy">
        <CenteredText class="text-4xl">
          67%
        </CenteredText>
      </Card>
    </Cards>
  </div>
</template>

<script>
import Header from '~/components/common/Header.vue'
import Cards from '~/components/common/Cards.vue'
import Card from '~/components/common/Card.vue'
import CenteredText from '~/components/common/CenteredText.vue'
import trend from 'vuetrend'
import DistTensorflow from '~/lib/tensorflow.js';

const gradient = ['#ffffff', '#ff974d']

export default {
  layout: 'client',
  components: {
    Header,
    Cards,
    Card,
    CenteredText,
    trend
  },
  data() {
    return {
      running: false,
      gradient,
      title: "MNIST",
      major: [0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0],
      tf: new DistTensorflow(this.$route.params.id, function (metrics, batchNo) {
        console.log(metrics);
        console.log(batchNo);
      }),
    }
  },
  created: function () {
  },
  methods: {
    toggleRunning: function () {
      (this.running) ? console.log("stop") : console.log("start");
      this.running = !this.running;
    }
  },
}
</script>
