<template>
  <div>
    <Header title="Active Sessions"></Header>
    <Cards>
      <Card v-for="model in models" :title="model.title" arrow>
        <div class="flex">
          <Subcard subtitle="Elapsed">
            <CenteredText class="text-4xl">
            {{ model.elapsed }}
            </CenteredText>
          </Subcard>
          <Subcard subtitle="Loss">
            <CenteredText class="text-4xl">
            {{ model.loss }}
            </CenteredText>
          </Subcard>
        </div>
      </Card>
    </Cards>
  </div>
</template>

<script>
import Header from '~/components/common/Header.vue'
import Cards from '~/components/common/Cards.vue'
import Card from '~/components/common/Card.vue'
import Subcard from '~/components/common/Subcard.vue'
import CenteredText from '~/components/common/CenteredText.vue'

export default {
  components: {
    Header,
    Cards,
    Card,
    Subcard,
    CenteredText
  },
  data: () => ({
    models: [
      {
        title: "Hello",
        elapsed: "10:43.4",
        loss: 43,
      },
      {
        title: "Hello",
        elapsed: "10:43.4",
        loss: 43,
      },
    ],
  }),
  created: function () {
    // Initialize all the models and format them follow the data format above
    let res = fetch('localhost:10200/models')
  }
}
</script>
