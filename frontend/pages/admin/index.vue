<template>
  <div>
    <Header title="Sessions"></Header>
    <Cards>
      <Card title="Hey 2" arrow>
        <div class="flex">
          <Subcard subtitle="Elapsed">
            <CenteredText class="text-4xl">
              10:25.4
            </CenteredText>
          </Subcard>
          <Subcard subtitle="Loss">
            <CenteredText class="text-4xl">
              10000.1
            </CenteredText>
          </Subcard>
        </div>
      </Card>
      <Card title="Lol" arrow>
        <div class="flex">
          <Subcard subtitle="Elapsed">
            <CenteredText class="text-4xl">
              10:25.4
            </CenteredText>
          </Subcard>
          <Subcard subtitle="Loss">
            <CenteredText class="text-4xl">
              10000.1
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
  }
}
</script>
