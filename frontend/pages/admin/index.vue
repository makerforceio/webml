<template>
  <div>
    <Header title="Sessions">
      <template v-slot:buttons-right>
        <button class="my-2 ml-4" @click="showNewSessionDialog = true">
          <fa-icon size="2x" :icon="['far', 'plus']" />
        </button>
      </template>
    </Header>
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
    <NewSessionDialog :show.sync="showNewSessionDialog" />
  </div>
</template>

<script>
import Header from '~/components/common/Header.vue'
import Cards from '~/components/common/Cards.vue'
import Card from '~/components/common/Card.vue'
import Subcard from '~/components/common/Subcard.vue'
import CenteredText from '~/components/common/CenteredText.vue'
import NewSessionDialog from '~/components/admin/NewSessionDialog.vue'

export default {
  components: {
    Header,
    Cards,
    Card,
    Subcard,
    CenteredText,
    NewSessionDialog
  },
  data() {
    return {
      showNewSessionDialog: false
		models: [],
    };
},
  created: function () {
    // Initialize all the models
    let res = fetch('localhost:10200/models')
  }
}
</script>
