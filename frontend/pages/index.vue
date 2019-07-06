<template>
  <div>
    <Header title="Session title" subtitle="By Author">
      <template v-slot:buttons-left>
        <button class="my-3 mr-4">
          <fa-icon size="2x" :icon="['far', 'arrow-left']" />
        </button>
      </template>
      <template v-slot:buttons-right>
        <button class="my-3 ml-4">
          <fa-icon size="2x" :icon="['far', 'pause']" />
        </button>
      </template>
      <template v-slot:content>
        <div class="mx-4 my-8" title="Major graph">
          <h2 class="font-medium">
            Loss
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
      <Card title="Elapsed">
        <CenteredText class="text-4xl">
          10m 45s
        </CenteredText>
      </Card>
      <Card title="Elapsed">
        <CenteredText class="text-4xl">
          10:45
        </CenteredText>
      </Card>
      <Card title="Elapsed">
        <CenteredText class="text-4xl">
          10000.1
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
      gradient,
      major: [0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0]
    }
  }
}
</script>
