<template>
  <div>
    <Header title="Session title" subtitle="By Author">
      <template v-slot:buttons-left>
		  <router-link class="my-2 mr-4" :to="{ name: 'admin' }">
          <fa-icon size="2x" :icon="['far', 'arrow-left']" />
        </router-link>
      </template>
      <template v-slot:buttons-right>
        <button class="my-2 ml-4">
          <fa-icon size="2x" :icon="['far', 'times-hexagon']" />
        </button>
        <button class="my-2 ml-4">
          <fa-icon size="2x" :icon="['far', 'pause']" />
        </button>
      </template>
      <template v-slot:content>
        <div class="mx-4 my-8" title="Major graph">
			<div class="font-medium">Share this link</div>
			<div class="font-bold text-xl">{{ shareLink }}</div>
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
          10:45
        </CenteredText>
      </Card>
      <Card subtitle="Epoch">
        <CenteredText class="text-4xl">
          100
        </CenteredText>
      </Card>
      <Card subtitle="Client count">
        <CenteredText class="text-4xl">
          4
        </CenteredText>
      </Card>
      <Card subtitle="Accuracy">
        <CenteredText class="text-4xl">
          78%
        </CenteredText>
      </Card>
      <Card subtitle="Loss">
        <trend :data="major" :height="200" auto-draw smooth> </trend>
      </Card>
      <Card subtitle="Loss">
        <trend :data="major" :height="200" auto-draw smooth> </trend>
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
	computed: {
		
		shareLink() {
			return "https://staging.webml.app/" + this.$route.params.id;
		},
	},
  data() {
    return {
      gradient,
      major: [0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0]
    }
  }
}
</script>
